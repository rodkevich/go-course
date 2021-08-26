package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rodkevich/go-course/homework/hw_weather_service/history"
	"github.com/rodkevich/go-course/homework/hw_weather_service/history/types"

	"github.com/gin-gonic/gin"
)

const serviceName = "history_service_001"

var (
	historyPort = os.Getenv("HISTORYPORT")
	esClient    *history.Client
)

func init() {
	esClient = history.NewEsClient(serviceName)
}

func setupRouter() (engine *gin.Engine) {
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
		rtn, err := esClient.Save(serviceName, req.String())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error post logs": "Check your request"})
			return
		}
		c.Header("Content-Type", "application/json")
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
		// c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
		c.String(http.StatusOK, rtn)
	})
	return engine
}

func main() {
	r := setupRouter()
	err := r.Run(":" + historyPort)
	if err != nil {
		log.Fatal(err)
	}
}

/*
For NO auth just remove headers. You'll get 401 status code error

GET With auth:

curl --location --request GET 'http://localhost:9091/logs/this%20will%20be%20logged' \
--header 'Authorization: Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl'

POST With auth:

curl --location --request POST 'http://localhost:9091/logs/create' \
--header 'Authorization: Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "this text will be logged"
}'

*/
