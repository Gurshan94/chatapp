package router

import (
	"github.com/Gurshan94/chatapp/internal/user"
	"github.com/Gurshan94/chatapp/internal/ws"
	"github.com/Gurshan94/chatapp/internal/room"
	"github.com/Gurshan94/chatapp/internal/room_users"
	
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler, roomHandler *room.Handler,roomUserHandler *room_users.Handler) {
	r = gin.Default()

	r.POST("/signup",userHandler.CreateUser)
	r.POST("/login",userHandler.Login)
	r.GET("/logout",userHandler.Logout)

	r.POST("/createroom",roomHandler.CreateRoom)
	r.GET("/getrooms",roomHandler.GetRooms)
	r.GET("/getroombyid/:roomId",roomHandler.GetRoomByID)
	r.POST("/deleteroom/:roomId",roomHandler.DeleteRoom)

	r.POST("/addusertoroom",roomUserHandler.AddUserToRoom)
	r.GET("/deleteuserfromroom",roomUserHandler.DeleteUserFromRoom)
	r.GET("/getroomsjoinedbyuser/:userId",roomUserHandler.GetRoomsJoinedByUser)

	r.POST("/ws/createroom", wsHandler.CreateRoom)
	r.GET("/ws/joinroom/:roomId",wsHandler.JoinRoom)
	r.GET("/ws/getrooms",wsHandler.GetRooms)
	r.GET("/ws/getclients/:roomId",wsHandler.GetClients)

}

func Start(addr string) error {
	return r.Run(addr)
}