package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rodkevich/go-course/homework/hw_weather_service/history"
	"github.com/rodkevich/go-course/homework/hw_weather_service/weather"

	"github.com/gin-gonic/gin"
)

var (
	port          = os.Getenv("WEATHERSERVICEPORT")
	weatherApiKey = os.Getenv("WEATHERSERVICEAPIKEY")
	clientOW      *weather.Client
	clientH       *history.Client
)
var db = make(map[string]string)

func init() {
	clientOW = weather.NewOpenWeatherClient(
		"https://api.openweathermap.org",
		"weather_service",
		"5e829817fc17ff78e515f532fa65302e",
		"metric",
	)
	clientH = history.NewClient()
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
		rtn, err := clientOW.GetWeatherByCityName(cityName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "check your request"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"status": "ok", "entry": rtn})
	})

	// engine.GET("/logs/:text", func(c *gin.Context) {
	engine.GET("/logs", func(c *gin.Context) {
		t := time.Now()
		defer func() {
			fmt.Println(time.Since(t).Seconds())
		}()
		var text string
		// text = c.Param("text")
		text = "weather_service_ES_logging"

		rtn := clientH.WriteToIndex("test", text)
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
		value, ok := db[user]
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
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
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
			db[user] = json.Value
			cityName := "name"
			rtn, err := clientOW.GetWeatherByCityName(cityName)
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
