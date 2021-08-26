package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/google/uuid"
)

var (
	kibanaURL   = os.Getenv("KIBANAURL")
	historyURL  = os.Getenv("HISTORYURL")
	weatherURL  = os.Getenv("WEATHERURL")
	gatewayPort = os.Getenv("GATEWAYPORT")
)

func main() {

	kibana, err := url.Parse(kibanaURL)
	if err != nil {
		log.Fatal(err)
	}

	logs, err := url.Parse(historyURL)
	if err != nil {
		log.Fatal(err)
	}

	weather, err := url.Parse(weatherURL)
	if err != nil {
		log.Fatal(err)
	}

	// this one requires no basic auth
	proxyKibana := httputil.NewSingleHostReverseProxy(kibana)

	// this one requires basic auth - gopher : historyService
	proxyLogs := httputil.NewSingleHostReverseProxy(logs)

	proxyLogs.Director = func(r *http.Request) {
		request, err := httputil.DumpRequest(r, true)
		if err != nil {
			return
		}
		fmt.Printf("%q", request)
		// set auth
		r.Header = map[string][]string{
			"Authorization": {"Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl"}}
		r.Host = logs.Host
		r.URL.Scheme = logs.Scheme
		r.URL.Host = logs.Host
	}

	// this one requires basic auth - gopher : weatherService
	proxyWeather := httputil.NewSingleHostReverseProxy(weather)

	proxyWeather.Director = func(req *http.Request) {
		ID, _ := uuid.NewUUID()
		req.Header = map[string][]string{
			// set trace ID
			"traceID": {ID.String()},
			// set auth
			"Authorization": {"Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl"}}
		req.Host = weather.Host
		req.URL.Scheme = weather.Scheme
		req.URL.Host = weather.Host
	}

	http.Handle("/", proxyKibana)
	http.Handle("/logs/", proxyLogs)
	http.Handle("/city/", proxyWeather)

	log.Fatal(http.ListenAndServe(":"+gatewayPort, nil))
}

/*
Requesting through gateway with NO auth won't call 401 status code error
because headers are added by itself

to log smth:
curl -X POST 'http://localhost:10000/logs/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "this is something",
    "traceID": "17d9eeae-0692-11ec-b323-0242ac180008",
    "timestamp": "2021-08-26T17:12:15.425Z",
    "body": "that will be logged"
}'

to search logs:
curl -X GET 'http://localhost:10000/logs/this%20will%20be%20logged'
curl -X GET 'http://localhost:10000/logs/17d9eeae-0692-11ec-b323-0242ac180008'

to get weather or 9090 it isn't protected:
curl -X GET 'http://localhost:10000/city/Chicago'
then curl for logs
*/
