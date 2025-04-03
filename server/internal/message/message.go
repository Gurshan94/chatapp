package message

import (
	"context"
	"time"
)

type Message struct {
	ID        int64 `json:"id"`
	RoomID    int64 `json:"roomid"`
	SenderID  int64 `json:"senderid"`
	Content   string `json:"content"`
	Deleted   bool `json:"deleted"`
	CreatedAt time.Time
}

type AddMessagereq struct {
	RoomID    int64 `json:"roomid"`
	SenderID  int64 `json:"senderid"`
	Content   string `json:"content"`
}

type AddMessageRes struct {
	ID        int64 `json:"id"`
	RoomID    int64 `json:"roomd"`
	SenderID  int64 `json:"senderid"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"createdat"`
}

type FetchMessage struct {
	ID        int64 `json:"id"`
	RoomID    int64 `json:"roomid"`
	SenderID  int64 `json:"senderid"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Deleted   bool `json:"deleted"`
	CreatedAt time.Time
}

type FetchMessageRes struct {
	Messages []*FetchMessage `json:"messages"`
	Cursor *time.Time `json:"cursor"`
}

type FetchMessageReq struct {
	RoomID int64 `json:"roomid"`
	Limit int64 `json:"limit"`
	Cursor *time.Time `json:"cursor"`
}

type Repository interface {
	AddMessage (ctx context.Context, message *Message) (*Message, error) 
	FetchMessage (ctx context.Context, roomID, limit int64, cursor *time.Time) ([]*FetchMessage, *time.Time, error)
	DeleteMessage (ctx context.Context, messageID int64) error 
}

type Service interface {
	AddMessage (c context.Context, req *AddMessagereq) (*AddMessageRes, error) 
	FetchMessage (c context.Context, req *FetchMessageReq) (*FetchMessageRes, error)
	DeleteMessage (c context.Context, messageID int64) error 
}