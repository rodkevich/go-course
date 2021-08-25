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
	proxyLogs := httputil.NewSingleHostReverseProxy(logs)
	proxyLogs.Director = func(req *http.Request) {
		req.Header = map[string][]string{
			"Accept-Encoding": {"gzip, deflate"},
			"Accept-Language": {"en-us"},
			"Authorization":   {"Basic Zm9vOmJhcg=="}, // set auth
		}
		req.Host = logs.Host
		req.URL.Scheme = logs.Scheme
		req.URL.Host = logs.Host
	}
	proxyKibana := httputil.NewSingleHostReverseProxy(kibana)

	http.Handle("/", proxyKibana)
	http.Handle("/logs/", proxyLogs)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
