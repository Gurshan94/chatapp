package room_users

import (
	"context"
	"time"
)

type service struct{
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2)*time.Second,
	}
}


func (s *service) AddUserToRoom(c context.Context,req *AddUserToRoomReq) (*AddUserToRoomRes, error) {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	r:= &RoomUser{
		RoomID:req.RoomID,
		UserID: req.UserID,
	}

	roomuser, err:= s.Repository.AddUserToRoom(ctx, r)
    if err!=nil {
		return nil, err
	}

	res := &AddUserToRoomRes{
		ID: roomuser.ID,
		RoomID: roomuser.RoomID,
		UserID: roomuser.UserID,
	}
	return res, nil
}
func (s *service) DeleteUserFromRoom(c context.Context, req *DeleteUserFromRoomReq) error {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()
	
	r:= &RoomUser{
		RoomID:req.RoomID,
		UserID: req.UserID,
	}

	err:= s.Repository.DeleteUserFromRoom(ctx, r)
	if err!=nil {
		return err
	}
	
	return nil

}

func (s *service) GetRoomsJoinedByUser(c context.Context, userID int64) ([]*GetRoomsJoinedByUserRes, error) {

	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	rooms, currentusers, err:= s.Repository.GetRoomsJoinedByUser(ctx, userID)
    if err!=nil {
		return nil, err
	}

	res:=make([]*GetRoomsJoinedByUserRes,0,len(rooms))

	for index,value:= range rooms {

		room :=GetRoomsJoinedByUserRes{
			ID:value.ID,
			Roomname: value.Roomname,
			Maxusers: value.Maxusers,
			AdminID: value.AdminID,
			CurrentUsers: currentusers[index],
		}
		
		res=append(res,&room)
	}

	return res, nil

}
