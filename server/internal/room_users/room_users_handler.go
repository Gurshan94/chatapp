package room_users

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	Service
}

func NewHandler (s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) AddUserToRoom(c *gin.Context) {
	var req AddUserToRoomReq
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	res, err:=h.Service.AddUserToRoom(c.Request.Context(),&req)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}
func (h *Handler) DeleteUserFromRoom(c *gin.Context) {
	var req DeleteUserFromRoomReq
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

    ok:=h.Service.DeleteUserFromRoom(c.Request.Context(),&req)
	if ok!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":ok.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messsage":"Deleted user from room"})

}

func (h *Handler) GetRoomsJoinedByUser(c *gin.Context) {
	userIDStr := c.Param("userId")
    userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	res, err:=h.Service.GetRoomsJoinedByUser(c.Request.Context(),userID)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}