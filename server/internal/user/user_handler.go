package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Gurshan94/chatapp/util"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var user LoginUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", u.AccessToken, 3600, "/", "localhost", false, false)

	res := &LoginUserRes{
		AccessToken: u.AccessToken,
		Username: u.Username,
		ID:       u.ID,
		Email: u.Email,
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout( c*gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK,gin.H{"message":"logout successfull"})
}

func(h *Handler) GetUserByID(c *gin.Context) {
	userIDStr:=c.Param("userid")
	userID, err:= strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.Service.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) Me(c *gin.Context) {
    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        return
    }

    claims, ok := user.(*util.MyJWTClaims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":       claims.ID,
        "username": claims.Username,
        "email":    claims.Email,
    })
}
