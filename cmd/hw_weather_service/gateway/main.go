package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	logs, err := url.Parse("http://0.0.0.0:9090")
	if err != nil {
		log.Fatal(err)
	}
	kibana, err := url.Parse("http://0.0.0.0:5601")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", httputil.NewSingleHostReverseProxy(kibana))
	http.Handle("/logs/", httputil.NewSingleHostReverseProxy(logs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
