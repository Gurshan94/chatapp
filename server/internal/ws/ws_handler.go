package ws

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
	"log"
)

type Handler struct {
	hub *Hub
}

func Newhandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:	1024,
	WriteBufferSize:1024,
	CheckOrigin: func(r *http.Request)bool {
		return true
	},
}


func(h *Handler) Servews(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
    
	userID, err := strconv.ParseInt(c.Param("userid"), 10, 64)
    if err != nil {
        log.Println("Invalid user ID")
        conn.Close()
        return
    }

    // Create new client
    client := &Client{
        ID:   userID,
        Conn: conn,
        MessageChannel: make(chan *IncomingMessage),
    }

    // Register client (Hub will fetch rooms)
    h.hub.Register <- client

    // Start client's message handling
    go client.ReadMessage(h.hub)
    go client.WriteMessage()

}