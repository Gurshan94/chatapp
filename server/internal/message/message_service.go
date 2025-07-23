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
func (s *service) FetchMessage(c context.Context, req *FetchMessageReq) ([]*FetchMessageRes, error) {
	ctx, cancel:= context.WithTimeout(c, s.timeout)
	defer cancel()

	messages, err:= s.Repository.FetchMessage(ctx, req.RoomID, req.Limit, req.Cursor)
	if err!=nil {
		return nil, err
	}
	
	res:=make([]*FetchMessageRes,0,len(messages))

	for _, value := range messages {
		r:=FetchMessageRes{
			ID:           value.ID,
			RoomID:     value.RoomID,
			SenderID:     value.SenderID,
			Username:      value.Username,
			Content: value.Content,
			Deleted: value.Deleted,
			CreatedAt: value.CreatedAt,
		}

		res=append(res,&r) 
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
