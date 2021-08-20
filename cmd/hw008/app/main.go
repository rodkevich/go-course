package main

import (
	"net/http"

	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	port = os.Getenv("SAVETOKENSERVICEPORT")
	DB   = make(map[string]Entry)
)

// Entry to save to fake-db
type Entry struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func setupRouter() (engine *gin.Engine) {
	engine = gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "i'm oline don't worry")
	})

	engine.POST("/save-token", func(c *gin.Context) {
		var entry Entry
		err := c.Bind(&entry)
		if err != nil {
			c.JSON(400, gin.H{"status": "validation_error", "error": err.Error()})
			return
		}
		DB[entry.Token] = entry
		c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": &entry})
	})
	return
}

func main() {
	r := setupRouter()
	r.Run(":" + port)
}
