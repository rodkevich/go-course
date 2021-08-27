package history

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rodkevich/go-course/homework/hw_weather_service/history/client"
	"github.com/rodkevich/go-course/homework/hw_weather_service/history/types"
)

const serviceName = "history_service_001"

var esClient *client.Client

func init() {
	esClient = client.NewEsClient(serviceName)
}

// SetupService ...
func SetupService() (engine *gin.Engine) {
	engine = gin.Default()

	// waiting for "Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl"
	authorized := engine.Group("/", gin.BasicAuth(gin.Accounts{
		"gopher": "historyService",
	}))

	authorized.POST("/logs/create", func(c *gin.Context) {
		var req types.LogPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "request doesn't match LogPostRequest - prototype "})
			return
		}
		rtn, err := esClient.SaveWithIndex(serviceName, req.String())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error post logs": "Check your request"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
	})

	authorized.GET("/logs/:querySearch", func(c *gin.Context) {
		querySearch := c.Params.ByName("querySearch")
		rtn, err := esClient.SearchForEntries(querySearch)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": "Check your request"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, rtn)
	})
	return engine
}
