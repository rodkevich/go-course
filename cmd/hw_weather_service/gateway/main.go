package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/google/uuid"
)

var (
	kibanaURL       = os.Getenv("KIBANAURL")
	historyURL      = os.Getenv("HISTORYURL")
	historyWriteURL = os.Getenv("HISTORYWRITEURL")
	weatherURL      = os.Getenv("WEATHERURL")
)

func main() {

	kibana, err := url.Parse(kibanaURL)
	// kibana, err := url.Parse("http://hw_weather_service_kib01:5601")
	// kibana, err := url.Parse("http://0.0.0.0:5601")
	if err != nil {
		log.Fatal(err)
	}

	logs, err := url.Parse(historyURL)
	// logs, err := url.Parse("http://app-history:9091")
	// logs, err := url.Parse("http://0.0.0.0:9091")
	if err != nil {
		log.Fatal(err)
	}

	weather, err := url.Parse(weatherURL)
	// weather, err := url.Parse("http://app-weather:9090")
	// weather, err := url.Parse("http://0.0.0.0:9090")
	if err != nil {
		log.Fatal(err)
	}
	// this one requires no basic auth
	proxyKibana := httputil.NewSingleHostReverseProxy(kibana)

	// this one requires gopher : historyService
	proxyLogs := httputil.NewSingleHostReverseProxy(logs)
	proxyLogs.Director = func(req *http.Request) {
		req.Header = map[string][]string{
			"Authorization": {"Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl"}, // set auth
		}
		req.Host = logs.Host
		req.URL.Scheme = logs.Scheme
		req.URL.Host = logs.Host
	}

	// this one requires gopher : weatherService
	proxyWeather := httputil.NewSingleHostReverseProxy(weather)
	proxyWeather.Director = func(req *http.Request) {
		ID, _ := uuid.NewUUID()
		req.Header = map[string][]string{
			"traceID":       {ID.String()},                          // set trace ID
			"Authorization": {"Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl"}, // set auth
		}
		req.Host = weather.Host
		req.URL.Scheme = weather.Scheme
		req.URL.Host = weather.Host
	}

	http.Handle("/", proxyKibana)
	http.Handle("/logs/", proxyLogs)
	http.Handle("/city/", proxyWeather)
	log.Fatal(http.ListenAndServe(":10000", nil))
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

/*
Requesting with NO auth won't call 401 status code error
because headers are added by gateway itself

curl --location --request GET 'http://localhost:10000/logs/this%20will%20be%20logged'
curl --location --request GET 'http://localhost:10000/city/Molodechno'


*/
