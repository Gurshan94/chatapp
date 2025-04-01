package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func Newhandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Maxusers int64 `json:"maxusers"`
	Admin int64 `json:"admin"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:		req.ID,
		Name:	req.Name,
		Clients:make(map[string]*Client),
		Maxusers: req.Maxusers,
		Admin: req.Admin,
	}

	c.JSON(http.StatusOK,req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:	1024,
	WriteBufferSize:1024,
	CheckOrigin: func(r *http.Request)bool {
		return true
	},
}

func (h *Handler) JoinRoom (c *gin.Context) {
	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	room,exists := h.hub.Rooms[roomID]
    if !exists {
		c.JSON(http.StatusNotFound,gin.H{"error":"room not found"})
		return
	}

	if len(room.Clients)==int(room.Maxusers) {
		c.JSON(http.StatusBadRequest,gin.H{"error":"The room is Full"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	cl := &Client{
		ID : clientID,
		RoomID : roomID,
		Username : username,
		Conn : conn,
		Message : make(chan *Message,10),
	}

     m := &Message{
		Content:"a new user has joined the room",
		RoomID: roomID,
		Username: username,
	 }

	 // register a new client through register channel
	 h.hub.Register <- cl

	 // broadcast that message
	 h.hub.Broadcast <- m
     
	 go cl.WriteMessage()
	 cl.ReadMessage(h.hub)
}

type RoomRes struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Maxusers int64 `json:"maxusers"`
	Admin int64 `json:"admin"`
}

func (h *Handler) GetRooms (c* gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:r.ID,
			Name:r.Name,
			Maxusers: r.Maxusers,
			Admin: r.Admin,
		})
	}
	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	ID string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients (c* gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients,ClientRes{
			ID: c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}

