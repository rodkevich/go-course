package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rodkevich/go-course/homework/hw_weather_service/history"
	"github.com/rodkevich/go-course/homework/hw_weather_service/weather"
)

const serviceName = "weather_service"

var (
	openWeatherBaseURL = os.Getenv("OPENWEATHERBASEURL")
	openWeatherAPIKey  = os.Getenv("OPENWEATHERAPIKEY")
	historyWriteURL    = os.Getenv("HISTORYWRITEURL")
	weatherPort        = os.Getenv("WEATHERPORT")
	esClient           *history.Client
	wsc                *weather.Client
)

func init() {
	wsc = weather.NewOpenWeatherClient(
		openWeatherBaseURL,
		serviceName,
		openWeatherAPIKey,
		"metric",
	)
	esClient = history.NewEsClient(serviceName)

}

func setupRouter() (engine *gin.Engine) {
	engine = gin.Default()

	engine.GET("/city/:name", func(c *gin.Context) {
		cityName := c.Param("name")
		rtn, err := wsc.GetByCityName(cityName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error get city": "check your request"})
			return
		}
		// log from gateway
		if traceID := c.Request.Header.Get("traceID"); traceID != "" {
			now := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
			body := `
					{
                    "title": "` + serviceName + `",
                    "traceID": "` + traceID + `",
                    "timestamp": "` + now + `",
                    "body": "Very useful information"
                	}
					`
			// err = logToHistory(body)
			rtn, err := esClient.Save(serviceName, body)
			if err != nil {
				log.Println("error esClient.Save logging:", err)
			}
			log.Println("saving request with traceID: ", traceID, " ",rtn)
			if err != nil {
				log.Println("error /city/:name logging:", err)
			}
		}
		// log NOT from gateway
		if traceID := c.Request.Header.Get("traceID"); traceID == "" {
			now := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
			body := `
					{
                    "title": "` + serviceName + `",
                    "traceID": "User: Unknown",
                    "timestamp": "` + now + `",
                    "body": "Very useful information about Unknown User Request"
                	}
					`
			rtn, err := esClient.Save(serviceName, body)
			log.Println("saving UnknownUserRequest", rtn)
			if err != nil {
				log.Println("error /city/:name logging:", err)
			}
		}
		c.Header("Content-Type", "application/json")
		log.Println(rtn)
		c.String(http.StatusOK, rtn)
	})
	return engine
}

func main() {
	r := setupRouter()
	err := r.Run(":" + weatherPort)
	if err != nil {
		log.Fatal(err)
	}
}

// func logToHistory(text string) (err error) {
//
// 	url := historyWriteURL
// 	method := "POST"
//
// 	payload := strings.NewReader(`{
//     "title": "whyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
//     "traceID": "nnnnnnnnnnneeeeeeeee",
//     "timestamp": "rrraaabooooooooooootaaaaaaaeeettttttt",
//     "body": "ggggggggggggggggg"
// }`)
//
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)
//
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	req.Header.Add("Content-Type", "application/json")
//
// 	res, err := client.Do(req)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()
//
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	log.Println(string(body))
// 	return
// }

// func logToHistory(text string) (err error) {
// 	body, err := json.Marshal(text)
// 	log.Println(body)
// 	if err != nil {
// 		return
// 	}
// 	resp, err := http.Post(
// 		historyWriteURL,
// 		"application/json",
// 		bytes.NewBuffer(body),
// 	)
// 	log.Println(resp)
// 	if err != nil {
// 		return
// 	}
// 	resp.Body.Close()
// 	return
// }
