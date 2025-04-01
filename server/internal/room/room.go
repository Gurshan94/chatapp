package room

import (
	"context"
	"github.com/Gurshan94/chatapp/internal/user"
)

type Room struct {
	ID       int64 `json:"id"`
	Roomname string `json:"room_name"`
	Maxusers int `json:"max_users"`
	AdminID  int64 `json:"admin_id"`
}

type CreateRoomReq struct {
	Roomname string `json:"room_name"`
	Maxusers int `json:"max_users"`
	AdminID  int64 `json:"admin_id"`
}

type CreateRoomRes struct {
	ID       int64 `json:"id"`
	Roomname string `json:"room_name"`
	Maxusers int `json:"max_users"`
	AdminID  int64 `json:"admin_id"`
}

type GetRoomsReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetRoomsRes struct {
	ID           int64  `json:"id"`
	Roomname     string `json:"room_name"`
	Maxusers     int  `json:"max_users"`
	AdminID      int64  `json:"admin_id"`
	CurrentUsers int  `json:"current_users"`
}

type GetRoomByIDRes struct {
	ID           int64  `json:"id"`
	Roomname     string `json:"room_name"`
	Maxusers     int  `json:"max_users"`
	AdminID      int64  `json:"admin_id"`
	CurrentUsers int  `json:"current_users"`
	Users        []*user.User `json:"users"`
}

type Repository interface {
	CreateRoom(ctx context.Context, room *Room) (*Room, error)
	GetRoomByID(ctx context.Context, roomID int64) (*Room, []*user.User, int, error)
	GetRooms(ctx context.Context, Limit, Offset int) ([]*Room,[]int, error)
	DeleteRoom(ctx context.Context, roomID int64) error
}

type Service interface {
	CreateRoom(c context.Context, req *CreateRoomReq) (*CreateRoomRes, error)
	GetRoomByID(c context.Context, roomID int64) (*GetRoomByIDRes, error) 
	GetRooms(c context.Context, req *GetRoomsReq) ([]*GetRoomsRes, error)
	DeleteRoom(c context.Context, roomID int64) error
}