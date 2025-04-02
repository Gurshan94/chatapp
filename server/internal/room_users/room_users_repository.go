package room_users

import (
	"context"
	"database/sql"
	"github.com/Gurshan94/chatapp/internal/room"
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

func (r *repository) AddUserToRoom(ctx context.Context,roomuser *RoomUser) (*RoomUser, error) {
    var lastinsertid int
	query:= `INSERT INTO room_users (room_id,user_id) VALUES ($1, $2) returning id`
	err := r.db.QueryRowContext(ctx, query, roomuser.RoomID, roomuser.UserID).Scan(&lastinsertid)
	if err!=nil {
		return nil, err
	}
	roomuser.ID=int64(lastinsertid)

	return roomuser, nil
}

func (r *repository) DeleteUserFromRoom(ctx context.Context, roomuser *RoomUser) error{
	query:= `DELETE FROM room_users WHERE room_id=$1 AND user_id=$2`
	_, err:= r.db.ExecContext(ctx, query, roomuser.RoomID, roomuser.UserID)
	if err!=nil {
		return err
	}

	return nil
}

func (r *repository) GetRoomsJoinedByUser(ctx context.Context, userID int64) ([]*room.Room, []int, error) {
	var rooms []*room.Room
	var users []int

	query:= `SELECT r.id, r.room_name, r.max_users, r.admin_id, COUNT(ru2.user_id) AS current_users
		     FROM rooms r 
			 JOIN room_users ru ON r.id = ru.room_id  
			 LEFT JOIN room_users ru2 ON r.id = ru2.room_id  
			 WHERE ru.user_id = $1  
			 GROUP BY r.id, r.room_name, r.max_users, r.admin_id;`

	rows, err:= r.db.QueryContext(ctx, query, userID)
	if err!=nil {
		return nil, nil, err
	}
	defer rows.Close()


	for rows.Next() {
		room:= &room.Room{}
		var CurrentUser int

		if err:= rows.Scan(&room.ID,&room.Roomname,&room.Maxusers,&room.AdminID,&CurrentUser); err!=nil{
			return nil, nil, err
		}

		rooms= append(rooms, room)
		users= append(users,CurrentUser)
	}

	return rooms, users, nil
	
}