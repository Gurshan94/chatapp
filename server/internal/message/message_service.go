package message

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

func (s *service) AddMessage(c context.Context, req *AddMessagereq) (*AddMessageRes, error) {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	m :=&Message{
		RoomID: req.RoomID,
		SenderID: req.SenderID,
		Content: req.Content,
	}

	message, err:=s.Repository.AddMessage(ctx, m)
	if err!=nil {
		return nil, err
	}

	res:=&AddMessageRes{
		ID: message.ID,
		RoomID: message.RoomID,
		SenderID: message.SenderID,
		Content: message.Content,
		CreatedAt: message.CreatedAt,
	}
	return res, nil

}
func (s *service) FetchMessage(c context.Context, req *FetchMessageReq) (*FetchMessageRes, error) {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	messages, lastcreateat, err:= s.Repository.FetchMessage(ctx, req.RoomID, req.Limit, req.Cursor)
	if err!=nil {
		return nil, err
	}

    res:=&FetchMessageRes{
		Messages: messages,
		Cursor: lastcreateat,
	}
	return res, nil
}
func (s *service) DeleteMessage(c context.Context, messageID int64) error {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	err:= s.Repository.DeleteMessage(ctx, messageID)
	if err!=nil {
		return err
	}
	
	return nil
}
