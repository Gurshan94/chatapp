package router

import (
	"github.com/Gurshan94/chatapp/internal/message"
	"github.com/Gurshan94/chatapp/internal/room"
	"github.com/Gurshan94/chatapp/internal/room_users"
	"github.com/Gurshan94/chatapp/internal/user"
	"github.com/Gurshan94/chatapp/internal/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler, roomHandler *room.Handler,roomUserHandler *room_users.Handler, messageHandler *message.Handler) {
	r = gin.Default()

	r.POST("/signup",userHandler.CreateUser)
	r.POST("/login",userHandler.Login)
	r.GET("/logout",userHandler.Logout)
	r.GET("/getuserbyid/:userid",userHandler.GetUserByID)

	r.POST("/createroom",roomHandler.CreateRoom)
	r.GET("/getrooms",roomHandler.GetRooms)
	r.GET("/getroombyid/:roomId",roomHandler.GetRoomByID)
	r.POST("/deleteroom/:roomId",roomHandler.DeleteRoom)

	r.POST("/addusertoroom",roomUserHandler.AddUserToRoom)
	r.POST("/deleteuserfromroom",roomUserHandler.DeleteUserFromRoom)
	r.GET("/getroomsjoinedbyuser/:userId",roomUserHandler.GetRoomsJoinedByUser)

	r.POST("/addmessage",messageHandler.AddMessage)
	r.POST("/deletemessage/:messageId",messageHandler.DeleteMessage)
	r.GET("/fetchmessage",messageHandler.FetchMessage)

	r.GET("/ws/:userid", wsHandler.Servews)

}

func Start(addr string) error {
	return r.Run(addr)
}