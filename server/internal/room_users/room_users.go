package room_users

import (
	"context"
	"github.com/Gurshan94/chatapp/internal/room"

)

type RoomUser struct {
	ID int64 `json:"id"`
	RoomID int64 `json:"roomid"`
	UserID int64 `json:"userid"`
}

type GetUsersInRoomRes struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetRoomsJoinedByUserRes struct {
	ID           int64  `json:"id"`
	Roomname     string `json:"roomname"`
	Maxusers     int    `json:"maxusers"`
	AdminID      int64  `json:"adminid"`
	CurrentUsers int    `json:"currentusers"`
}

type AddUserToRoomReq struct {
	RoomID int64 `json:"roomid"`
	UserID int64 `json:"userid"`
}

type AddUserToRoomRes struct {
	ID int64 `json:"id"`
	RoomID int64 `json:"roomid"`
	UserID int64 `json:"userid"`
}

type DeleteUserFromRoomReq struct {
	RoomID int64 `json:"roomid"`
	UserID int64 `json:"userid"`
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
