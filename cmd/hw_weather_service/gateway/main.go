package main

import (
	"fmt"
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
		// body := copy(body, req.Body)
		// b, _ := ioutil.ReadAll(body)
		// fmt.Println(string(b))
		fmt.Println(req.BasicAuth())
		req.Header = map[string][]string{
			"Accept-Encoding": {"gzip, deflate"},
			"Accept-Language": {"en-us"},
			"Authorization":   {"Basic Zm9vOmJhcg=="},
		}
		req.Host = logs.Host
		req.URL.Scheme = logs.Scheme
		req.URL.Host = logs.Host
		// req.URL.Path = c.Param("proxyPath")
	}
	proxyKibana := httputil.NewSingleHostReverseProxy(kibana)

	http.Handle("/", proxyKibana)
	http.Handle("/logs/", proxyLogs)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
