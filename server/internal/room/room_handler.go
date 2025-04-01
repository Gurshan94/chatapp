package room

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

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	res, err:=h.Service.CreateRoom(c.Request.Context(),&req)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetRooms(c *gin.Context) {
	var req GetRoomsReq
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	res, err:=h.Service.GetRooms(c.Request.Context(),&req)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetRoomByID(c *gin.Context) {
	
	roomIDStr := c.Param("roomId")
    roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	res, err:=h.Service.GetRoomByID(c.Request.Context(),roomID)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h *Handler) DeleteRoom(c *gin.Context) {
	roomIDStr := c.Param("roomId")
    roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	ok:=h.Service.DeleteRoom(c.Request.Context(),roomID)
	if ok!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":ok.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"messsage":"Room Delete sucessfully"})
}