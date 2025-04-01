package roomusers

import (
	"context"
	"github.com/Gurshan94/chatapp/internal/user"
	"github.com/Gurshan94/chatapp/internal/room"

)

type Room_Users struct {
	ID int64 `json:"id"`
	Room_ID int64 `json:"room_id"`
	Admin_ID int64 `json:"admin_id"`
}

type Repository interface {
	AddUserToRoom(ctx context.Context, roomID, userID int64) error 
	DeleteUserFromRoom(ctx context.Context, roomID, userID int64) error
	GetUsersInRoom(ctx context.Context,roomID int64) ([]*user.User,error)
	GetRoomsJoinedByUser(ctx context.Context, userID int64) ([]*room.Room, error)

}

type Service interface {
	AddUserToRoom(ctx context.Context, roomID, userID int64) error
	DeleteUserFromRoom(ctx context.Context, roomID, userID int64) error
	GetUsersInRoom(ctx context.Context, roomID int64) ([]*user.User, error)
	GetRoomsJoinedByUser(ctx context.Context, userID int64) ([]*room.Room, error)

}
