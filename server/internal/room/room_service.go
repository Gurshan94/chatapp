package room

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}
 
func (s *service) CreateRoom(c context.Context, req *CreateRoomReq) (*CreateRoomRes, error) {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	r := &Room{
		Roomname: req.Roomname,
		Maxusers: req.Maxusers,
		AdminID:  req.AdminID,
	}
    
	room, err:= s.Repository.CreateRoom(ctx, r)
	if err!=nil {
		return nil, err
	}

	res := &CreateRoomRes{
		ID: room.ID,
		Roomname:  room.Roomname,
		Maxusers: room.Maxusers,
		AdminID: room.AdminID,
	}

	return res, nil
}
func (s *service) GetRoomByID(c context.Context, roomID int64) (*GetRoomByIDRes, error) {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	room, users, CurrentUsers, err := s.Repository.GetRoomByID(ctx, roomID)
	if err!=nil {
		return nil, err
	}

	res := &GetRoomByIDRes{
		ID: room.ID,
		Roomname: room.Roomname,
		Maxusers: room.Maxusers,
		AdminID: room.AdminID,
		CurrentUsers: CurrentUsers,
		Users: users,
	}

	return res, nil
	

}
func (s *service) GetRooms(c context.Context, req *GetRoomsReq) ([]*GetRoomsRes, error){
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	rooms, currentusers, err:= s.Repository.GetRooms(ctx,req.Limit, req.Offset)
	if err!=nil {
		return nil, err
	}
    
	res:=make([]*GetRoomsRes,0,len(rooms))

	for index, value := range rooms {
		r:=GetRoomsRes{
			ID:           value.ID,
			Roomname:     value.Roomname,
			Maxusers:     value.Maxusers,
			AdminID:      value.AdminID,
			CurrentUsers: currentusers[index],
		}

		res=append(res,&r) 
	}
	
	return res, nil

}
func (s *service) DeleteRoom(c context.Context, roomID int64) error{
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	err:= s.Repository.DeleteRoom(ctx, roomID)
	if err!=nil {
		return err
	}
	
	return nil
}