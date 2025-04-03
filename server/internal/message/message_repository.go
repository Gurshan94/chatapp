package message

import (
	"context"
	"time"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) AddMessage(ctx context.Context, message *Message) (*Message, error) {
	var lastinsertid int64
	var createdat time.Time
	query := `INSERT INTO messages (room_id, sender_id, content) 
		      VALUES ($1, $2, $3) RETURNING id, created_at`

	err := r.db.QueryRowContext(ctx, query, message.RoomID, message.SenderID, message.Content).Scan(&lastinsertid, &createdat)
	if err != nil {
		return nil, err
	}
	message.ID= lastinsertid
	message.CreatedAt= createdat

	return message, nil
}

func (r *repository) FetchMessage(ctx context.Context, roomID, limit int64, cursor *time.Time) ([]*FetchMessage, *time.Time, error) {
	query := `
		SELECT m.id, m.room_id, m.sender_id, u.username,  m.content, m.deleted, m.created_at 
		FROM messages m 
		JOIN users u ON m.sender_id = u.id  
		WHERE m.room_id = $1 
		AND ($2::TIMESTAMP IS NULL OR m.created_at < $2) 
		AND m.deleted = FALSE
		ORDER BY m.created_at DESC 
		LIMIT $3;`
	rows, err := r.db.QueryContext(ctx, query, roomID, cursor, limit)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var messages []*FetchMessage
	var lastCreatedAt *time.Time

	for rows.Next() {
		var msg FetchMessage
		err := rows.Scan(&msg.ID, &msg.RoomID, &msg.SenderID, &msg.Username,&msg.Content, &msg.Deleted, &msg.CreatedAt)
		if err != nil {
			return nil, nil, err
		}
		messages = append(messages, &msg)
		lastCreatedAt = &msg.CreatedAt
	}
	return messages, lastCreatedAt, nil
}

func (r *repository) DeleteMessage(ctx context.Context, messageID int64) error {
	query := `UPDATE messages SET deleted = TRUE WHERE id = $1;`
	_, err := r.db.ExecContext(ctx, query, messageID)
	if err!=nil {
		return err
	}
	return nil
}
