package room

import (
	"context"
	"database/sql"
	"github.com/Gurshan94/chatapp/internal/user"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt,error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateRoom(ctx context.Context, room *Room) (*Room, error) {
	var lastinsertid int
	query := "INSERT INTO rooms(room_name, max_users, admin_id) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, room.Roomname, room.Maxusers, room.AdminID).Scan(&lastinsertid)
    if err!=nil {
		return nil, err
	}
	room.ID = int64(lastinsertid)

	return room,nil
}

func (r *repository) GetRoomByID(ctx context.Context, roomID int64) (*Room, []*user.User, int, error) {

	var room Room
	var CurrentUsers int

	query := `SELECT r.id, r.room_name, r.max_users, r.admin_id, 
       		  COALESCE(COUNT(ru.user_id), 0) AS current_users
			  FROM rooms r
			  LEFT JOIN room_users ru ON r.id = ru.room_id
			  WHERE r.id = $1
			  GROUP BY r.id, r.room_name, r.max_users, r.admin_id`

	err := r.db.QueryRowContext(ctx, query, roomID).Scan(&room.ID,&room.Roomname,&room.Maxusers,&room.AdminID,&CurrentUsers)
	if err!=nil {
		return nil, nil ,0 ,err
	}
    
	userquery := `SELECT u.id, u.username, u.email 
	              FROM users u INNER JOIN room_users ru ON u.id=ru.user_id 
				  WHERE ru.room_id=$1`
	
	rows, err :=r.db.QueryContext(ctx, userquery, roomID)
	if err!=nil {
		return nil, nil ,0 ,err
	}
	defer rows.Close()
    
	var Users []*user.User

	for rows.Next() {
		var user user.User
		if err:= rows.Scan(&user.ID,&user.Username,&user.Email); err!=nil {
			return nil, nil ,0 ,err
		}

        Users=append(Users,&user)
	}

 	return &room, Users, CurrentUsers, nil
}

func (r *repository) DeleteRoom(ctx context.Context, roomID int64) error {
	query := "DELETE FROM rooms WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, roomID)
	if err!=nil {
		return err
	}

	return nil
}

func (r *repository) GetRooms(ctx context.Context,Limit, Offset int, UserId int64) ([]*Room, []int, error) {
    
    var rooms []*Room
	var currentusers []int

	query := `SELECT r.id, r.room_name, r.max_users, r.admin_id, COUNT(ru.user_id) AS current_users 
	          FROM rooms r
			  LEFT JOIN room_users ru ON r.id=ru.room_id
			  WHERE r.id NOT IN (
              SELECT room_id FROM room_users WHERE user_id = $3
              )
			  GROUP BY r.id, r.room_name, r.max_users, r.admin_id
			  ORDER BY r.id
			  LIMIT $1 OFFSET $2`
			  
	rows, err := r.db.QueryContext(ctx, query, Limit, Offset, UserId)
	if err!=nil {
		return nil, nil, err
	}
    defer rows.Close()

	for rows.Next() {
		var room Room
		var CurrentUser int

		if err:= rows.Scan(&room.ID,&room.Roomname,&room.Maxusers,&room.AdminID,&CurrentUser); err!=nil {
			return nil, nil, err
		}

		rooms=append(rooms,&room)
		currentusers=append(currentusers,CurrentUser)

	}

	return rooms,currentusers,nil
}


