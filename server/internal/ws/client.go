package ws

import (
	"log"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type Client struct{
	ID int64 `json:"id"`
	Conn *websocket.Conn
	MessageChannel chan *IncomingMessage
}

type IncomingMessage struct {
	Content string `json:"content"`
	RoomID int64 `json:"roomid"`
	SenderID int64 `json:"senderid"`
}

type OutgoingMessage struct {
	ID int64 `json:"id"`
	Content string `json:"content"`
	RoomID int64 `json:"roomid"`
	SenderID int64 `json:"senderid"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage () {
	defer func(){
		c.Conn.Close()
	}()

	for {
		message, ok := <- c.MessageChannel
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage (h *Hub) {
	defer func() {
		h.HandleUnregister(c)
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

		var incomingMsg IncomingMessage
		err = json.Unmarshal(m, &incomingMsg)
		if err != nil {
			log.Println("unmarshal error:", err)
			return
		}

		h.Broadcast <- &incomingMsg
	}
}