package weather

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rodkevich/go-course/homework/hw_weather_service/history/client"
	"github.com/rodkevich/go-course/homework/hw_weather_service/history/types"
	"github.com/rodkevich/go-course/homework/hw_weather_service/weather/openweather"
)

const serviceName = "weather_service"

var (
	openWeatherBaseURL = os.Getenv("OPENWEATHERBASEURL")
	openWeatherAPIKey  = os.Getenv("OPENWEATHERAPIKEY")
	clientES           *client.Client
	clientOW           *openweather.Client
)

func init() {
	clientOW = openweather.NewOpenWeatherClient(
		openWeatherBaseURL,
		serviceName,
		openWeatherAPIKey,
		"metric",
	)
	clientES = client.NewEsClient(serviceName)
}

// SetupService ...
func SetupService() (engine *gin.Engine) {
	engine = gin.Default()

	engine.GET("/city/:name", func(c *gin.Context) {
		var traceID = "Unknown user request"
		if t := c.Request.Header.Get("traceID"); t != "" {
			traceID = t
		}
		start := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
		body := types.LogPostRequest{
			Title:     serviceName,
			TraceID:   traceID,
			Timestamp: start,
			Body:      "Very useful information about request from: " + c.Request.RemoteAddr,
		}
		_, err := clientES.SaveWithIndex(serviceName, body.String())
		if err != nil {
			log.Println("error 1st save /city/:name :", err)
		}

		// get weather
		cityName := c.Param("name")
		city, err := clientOW.GetByCityName(cityName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error get city": "such is life"})
		}

		// log again
		finish := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
		body = types.LogPostRequest{
			Title:     serviceName,
			TraceID:   traceID,
			Timestamp: finish,
			Body:      city,
		}
		_, err = clientES.SaveWithIndex(serviceName, body.String())
		if err != nil {
			log.Println("error 2nd save /city/:name :", err)
			return
		}
		log.Printf("saved data about: %v start: %v, finish: %v", traceID, start, finish)

		// return result
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, city)
	},
	)
	return engine
}
