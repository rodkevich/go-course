package main

import (
	"net/http"
	"os"

	historyService "github.com/rodkevich/go-course/homework/hw_weather_service/history"
	weatherService "github.com/rodkevich/go-course/homework/hw_weather_service/weather"
	"github.com/rodkevich/go-course/homework/hw_weather_service/weather/types"

	"github.com/gin-gonic/gin"
)

const serviceName = "final_weather_service"

var (
	port          = os.Getenv("WEATHERSERVICEPORT")
	weather       *weatherService.Client
	history       *historyService.Client
)
var users = make(map[string]string)

func init() {
	weather = weatherService.NewOpenWeatherClient(
		"https://api.openweathermap.org",
		serviceName,
		"5e829817fc17ff78e515f532fa65302e",
		"metric",
	)
	history = historyService.NewEsClient(serviceName)
}

func setupRouter() (engine *gin.Engine) {
	engine = gin.Default()

	engine.GET("/city/:name", func(c *gin.Context) {
		// var query string
		// if query, _ = c.GetQuery("name"); query == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "no search query present"})
		// 	return
		// }
		cityName := c.Param("name")
		rtn, err := weather.GetByCityName(cityName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error get city": "check your request"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
	})

	engine.POST("/logs/create", func(c *gin.Context) {
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

	engine.GET("/logs/:querySearch", func(c *gin.Context) {
		querySearch := c.Params.ByName("querySearch")
		rtn, _ := history.SearchForEntries(querySearch)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"status": "400", "error": "Check your request"})
		// 	return
		// }
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
	})

	// Get user value
	engine.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := users[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	// }))
	authorized := engine.Group("/", gin.BasicAuth(gin.Accounts{
		"foo": "bar", // user:foo password:bar
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")
		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			users[user] = json.Value
			cityName := "name"
			rtn, err := weather.GetByCityName(cityName)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "check your request"})
				return
			}
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
		}
	})
	return
}

func main() {
	port = "9090"
	r := setupRouter()
	r.Run(":" + port)
}
