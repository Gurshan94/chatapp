package router

import (
	"github.com/Gurshan94/chatapp/internal/message"
	"github.com/Gurshan94/chatapp/internal/room"
	"github.com/Gurshan94/chatapp/internal/room_users"
	"github.com/Gurshan94/chatapp/internal/user"
	"github.com/Gurshan94/chatapp/internal/ws"
	"github.com/Gurshan94/chatapp/middleware"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler, roomHandler *room.Handler,roomUserHandler *room_users.Handler, messageHandler *message.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/signup",userHandler.CreateUser)
	r.POST("/login",userHandler.Login)
	r.GET("/logout",userHandler.Logout)
	r.GET("/getuserbyid/:userid",userHandler.GetUserByID)

    auth:=r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/me",userHandler.Me)

	auth.POST("/createroom",roomHandler.CreateRoom)
	auth.GET("/getrooms",roomHandler.GetRooms)
	auth.GET("/getroombyid/:roomId",roomHandler.GetRoomByID)
	auth.POST("/deleteroom/:roomId",roomHandler.DeleteRoom)

	auth.POST("/addusertoroom",roomUserHandler.AddUserToRoom)
	auth.POST("/deleteuserfromroom",roomUserHandler.DeleteUserFromRoom)
	auth.GET("/getroomsjoinedbyuser/:userId",roomUserHandler.GetRoomsJoinedByUser)

	auth.POST("/addmessage",messageHandler.AddMessage)
	auth.POST("/deletemessage/:messageId",messageHandler.DeleteMessage)
	auth.GET("/fetchmessage/:roomId",messageHandler.FetchMessage)
	r.GET("/ws/:userid", wsHandler.Servews)

}

func Start(addr string) error {
	return r.Run(addr)
}