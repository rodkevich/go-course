package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var DB = make(map[string]interface{})

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "i'm oline don't worry")
	})
	r.POST("/save-token", func(c *gin.Context) {
		// token, _ := c.Params.Get("token")
		type Token struct {
			Name    string `json:"name" binding:"required"`
			Address string `json:"address" binding:"required"`
		}
		var token Token
		err := c.Bind(&token)
		if err != nil {
			c.JSON(400, gin.H{"status": "validation_error", "error": err})
			return
		}

		DB[token.Name] = token.Address
		c.JSON(http.StatusOK, gin.H{"status": "ok", "token": string(token.Address)})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":5051")
}
