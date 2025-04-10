package middleware

import (
    "net/http"
    "github.com/Gurshan94/chatapp/util"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString, err := c.Cookie("jwt")
        if err != nil || tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized - missing token"})
            c.Abort()
            return
        }

        claims, err := util.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized - invalid token"})
            c.Abort()
            return
        }

        // Store user info in context
        c.Set("user", claims)
        c.Next()
    }
}