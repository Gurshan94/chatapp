package ws

import (
	"context"
	"github.com/Gurshan94/chatapp/internal/message"
	"github.com/Gurshan94/chatapp/internal/user"
	"github.com/Gurshan94/chatapp/internal/room_users"
	"log"
)

type Hub struct {
	Clients    map[int64]*Client
	Rooms      map[int64]map[int64]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *IncomingMessage
	MessageRepo message.Repository
	RoomUsersRepo room_users.Repository
	Userrepo user.Repository
}

func NewHub(userrepo user.Repository, messagerepo message.Repository,roomuserrepo room_users.Repository) *Hub {
	return &Hub{
		Clients:make(map[int64]*Client),
		Rooms: make(map[int64]map[int64]*Client),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan *IncomingMessage, 10),
		MessageRepo: messagerepo,
		RoomUsersRepo: roomuserrepo,
		Userrepo: userrepo,
	}
}

func (h *Hub) Run () {

	for {
		select {
			case cl := <-h.Register:
				h.HandleRegister(cl)

			case cl := <-h.Unregister:
				h.HandleUnregister(cl)
		
			case m := <-h.Broadcast:
				h.HandleBroadcast(m)
		}
	}
}

func(h *Hub) HandleRegister(cl *Client) {
	h.Clients[cl.ID]=cl

	rooms, _, err:=h.RoomUsersRepo.GetRoomsJoinedByUser(context.Background(),cl.ID)
	if err != nil {
		log.Println("Error fetching user rooms:", err)
		return
	}

	for _, room := range rooms {
		if h.Rooms[room.ID]==nil {
			h.Rooms[room.ID]=make(map[int64]*Client)
		}
		h.Rooms[room.ID][cl.ID]=cl
	}
}

func(h *Hub) HandleUnregister(cl *Client) {
	delete(h.Clients,cl.ID)

	rooms, _, err := h.RoomUsersRepo.GetRoomsJoinedByUser(context.Background(), cl.ID)
    if err != nil {
        log.Printf("Error fetching rooms for user %d: %v", cl.ID, err)
        return
    }

    // Remove user only from the rooms they have joined
    for _, room := range rooms {
        if _, exists := h.Rooms[room.ID]; exists {
            delete(h.Rooms[room.ID], cl.ID)

            // If the room becomes empty, delete it
            if len(h.Rooms[room.ID]) == 0 {
                delete(h.Rooms, room.ID)
            }
        }
    }

}

func(h *Hub) HandleBroadcast(m *IncomingMessage) {

	msg:=&message.Message{
		RoomID: m.RoomID,
		SenderID: m.SenderID,
		Content: m.Content,
	}

	_, err:=h.MessageRepo.AddMessage(context.Background(),msg)
	if err!=nil {
		log.Printf("Error saving message to db %v", err)
        return
	}

	user, err:=h.Userrepo.GetUserByID(context.Background(),m.SenderID)
	if err!=nil{
		log.Printf("Error fetching user from db %v", err)
        return
	}

	outgoingmessage:=&OutgoingMessage{
		RoomID: m.RoomID,
		SenderID: m.SenderID,
		Content: m.Content,
		Username: user.Username,
	}

	if clients, ok := h.Rooms[m.RoomID]; ok { // Get clients in the room
        for _, client := range clients {
            err := client.Conn.WriteJSON(outgoingmessage) // Send message
            if err != nil {
                log.Printf("Error broadcasting message to user %d: %v", client.ID, err)
                client.Conn.Close()
                h.HandleUnregister(client)
            }
        }
    }
	
}


