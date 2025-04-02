package room_users

import (
	"context"
	"github.com/Gurshan94/chatapp/internal/room"

)

type RoomUser struct {
	ID int64 `json:"id"`
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}

type GetUsersInRoomRes struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type GetRoomsJoinedByUserRes struct {
	ID           int64  `json:"id"`
	Roomname     string `json:"room_name"`
	Maxusers     int    `json:"max_users"`
	AdminID      int64  `json:"admin_id"`
	CurrentUsers int    `json:"current_users"`
}

type AddUserToRoomReq struct {
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}

type AddUserToRoomRes struct {
	ID int64 `json:"id"`
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}

type DeleteUserFromRoomReq struct {
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}

type Repository interface {
	AddUserToRoom(ctx context.Context,roomuser *RoomUser) (*RoomUser,error) 
	DeleteUserFromRoom(ctx context.Context, roomuser *RoomUser) error
	GetRoomsJoinedByUser(ctx context.Context, userID int64) ([]*room.Room, []int, error)
}

type Service interface {
	AddUserToRoom(c context.Context,req *AddUserToRoomReq) (*AddUserToRoomRes,error)
	DeleteUserFromRoom(c context.Context, req *DeleteUserFromRoomReq) error
	GetRoomsJoinedByUser(c context.Context, userID int64) ([]*GetRoomsJoinedByUserRes, error)
}
