package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	weatherService "github.com/rodkevich/go-course/homework/hw_weather_service/weather"
)

const serviceName = "weather_service"

var (
	openWeatherBaseURL = os.Getenv("OPENWEATHERBASEURL")
	openWeatherApiKey  = os.Getenv("OPENWEATHERAPIKEY")
	historyWriteURL    = os.Getenv("HISTORYWRITEURL")
	weatherPort        = os.Getenv("GATEWAYPORT")

	weather *weatherService.Client
)

func init() {
	weather = weatherService.NewOpenWeatherClient(
		openWeatherBaseURL,
		serviceName,
		openWeatherApiKey,
		"metric",
	)
}

func setupRouter() (engine *gin.Engine) {
	engine = gin.Default()

	engine.GET("/city/:name", func(c *gin.Context) {
		cityName := c.Param("name")
		rtn, err := weather.GetByCityName(cityName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error get city": "check your request"})
			return
		}
		// log from gateway
		if traceID := c.Request.Header.Get("traceID"); traceID != "" {
			err = logToHistory(traceID + " and other useful information")
			log.Println("saving request with traceID: ", traceID)
			if err != nil {
				log.Println(err)
			}
		}
		// log NOT from gateway
		if traceID := c.Request.Header.Get("traceID"); traceID == "" {
			err = logToHistory("UnknownUserRequest" + "and other useful information")
			log.Println("saving UnknownUserRequest")
			if err != nil {
				log.Println(err)
			}
		}
		c.Header("Content-Type", "application/json")
		log.Println(rtn)
		c.String(http.StatusOK, rtn)
	})
	return
}

func logToHistory(text string) (err error) {
	body, err := json.Marshal(map[string]string{"title": text})
	if err != nil {
		return
	}
	resp, err := http.Post(
		historyWriteURL,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

func main() {
	// port = "9090"
	r := setupRouter()
	err := r.Run(":" + weatherPort)
	if err != nil {
		log.Fatal(err)
	}
}
