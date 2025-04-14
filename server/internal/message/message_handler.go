package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler (s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) AddMessage(c *gin.Context) {
	var req AddMessagereq
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	res, err:=h.Service.AddMessage(c.Request.Context(),&req)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) FetchMessage(c *gin.Context) {
	roomIDStr := c.Param("roomId")
    roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}	
	limitStr:=c.Query("limit")
	offsetStr:=c.Query("offset")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
        return
    }

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
        return
    }
    
	req:=FetchMessageReq{
		RoomID: roomID,
		Limit: limit,
		Offset: offset,
	}

	res, err:=h.Service.FetchMessage(c.Request.Context(),&req)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}


func (h *Handler) DeleteMessage(c *gin.Context) {
	messageIDStr := c.Param("messageId")
    messageID, err := strconv.ParseInt(messageIDStr, 10, 64)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	ok:=h.Service.DeleteMessage(c.Request.Context(),messageID)
	if ok!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":ok.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"messsage":"Message Deleted sucessfully"})
}