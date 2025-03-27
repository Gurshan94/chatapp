package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct{
	ID string `json:"id"`
	RoomID string `json:"roomid"`
	Username string `json:"username"`
	Conn *websocket.Conn
	Message chan *Message
}

type Message struct {
	Content string `json:"content"`
	RoomID string `json:"roomid"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage () {
	defer func(){
		c.Conn.Close()
	}()

	for {
		message, ok := <- c.Message
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage (h *Hub) {
	defer func() {
		h.Unregister <- c
        c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err!=nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Printf("error %v",err)
			}
			break
		}

		msg := &Message{
			Content: string(m),
			RoomID: c.RoomID,
			Username :c.Username,
		}

		h.Broadcast <- msg
	}
}