package main

import (
	"log"
	"github.com/Gurshan94/chatapp/db"
	"github.com/Gurshan94/chatapp/internal/user"
	"github.com/Gurshan94/chatapp/internal/ws"
	"github.com/Gurshan94/chatapp/internal/room"
	"github.com/Gurshan94/chatapp/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err!=nil {
		log.Fatal("could not initialize database",err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	roomRep := room.NewRepository(dbConn.GetDB())
	roomSvc := room.NewService(roomRep)
	roomHandler :=room.NewHandler(roomSvc)

    hub := ws.NewHub()
	wsHandler := ws.Newhandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler, roomHandler)
	router.Start("0.0.0.0:8080")
}