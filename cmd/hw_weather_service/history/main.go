package main

import (
	"net/http"
	"os"

	historyService "github.com/rodkevich/go-course/homework/hw_weather_service/history"
	"github.com/rodkevich/go-course/homework/hw_weather_service/weather/types"

	"github.com/gin-gonic/gin"
)

const serviceName = "history_service"

var (
	port    = os.Getenv("HISTORYSERVICEPORT")
	history *historyService.Client
)

func init() {
	history = historyService.NewEsClient(serviceName)
}

func setupRouter() (engine *gin.Engine) {
	engine = gin.Default()

	authorized := engine.Group("/", gin.BasicAuth(gin.Accounts{
		"foo": "bar", // user:foo password:bar
	}))

	authorized.POST("/logs/create", func(c *gin.Context) {
		var req types.LogPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rtn, err := history.Save(req.Title)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error post logs": "Check your request"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
	})

	authorized.GET("/logs/:querySearch", func(c *gin.Context) {
		querySearch := c.Params.ByName("querySearch")
		rtn, err := history.SearchForEntries(querySearch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": "Check your request"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
	})
	return
}

func main() {
	port = "9090"
	r := setupRouter()
	r.Run(":" + port)
}
