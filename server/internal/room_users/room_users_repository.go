package roomusers

import (
	"context"
	"database/sql"
	"github.com/Gurshan94/chatapp/internal/user"
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

func (r *repository) AddUserToRoom(ctx context.Context, roomID, userID int64) error {
	return nil
}
func (r *repository) DeleteUserFromRoom(ctx context.Context, roomID, userID int64) error{
	return nil
}
func (r *repository) GetUsersInRoom(ctx context.Context,roomID int64) ([]*user.User,error) {
	var users []*user.User
	return users,nil
	
}
func (r *repository) GetRoomsJoinedByUser(ctx context.Context, userID int64) ([]*room.Room, error) {
	var rooms []*room.Room

	query:= `SELECT r.id, r.room_name, r.max_users, r.admin_id, COUNT(ru2.user_id) AS current_users
		     FROM rooms r 
			 JOIN room_users ru ON r.id = ru.room_id  
			 LEFT JOIN room_users ru2 ON r.id = ru2.room_id  
			 WHERE ru.user_id = $1  
			 GROUP BY r.id, r.room_name, r.max_users, r.admin_id;`

	rows, err:= r.db.QueryContext(ctx, query, userID)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()


	for rows.Next() {
		room:= &room.Room{}
		if err:= rows.Scan(&room.ID,&room.Roomname,&room.Maxusers,&room.AdminID); err!=nil{
			return nil, err
		}

		rooms= append(rooms, room)
	}

	return rooms, nil
	
}