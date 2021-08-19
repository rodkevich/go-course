package main

import (
	"fmt"
	"sync"
	"time"
)

// Fetcher ...
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cacheCheckChannel chan *map[string]interface{}) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start).Nanoseconds())
	}()
	time.Sleep(time.Nanosecond * 2)

	type (
		memo struct {
			mu      sync.Mutex
			content map[string]interface{}
		}
		fetchResult struct {
			url   string
			body  string
			urls  []string
			err   error
			depth int
		}
	)

	funcContextChannel := make(chan *fetchResult)
	cache := memo{
		mu:      sync.Mutex{},
		content: map[string]interface{}{},
	}

	fetchClosure := func(url string, depth int) {
		body, urls, err := fetcher.Fetch(url)
		funcContextChannel <- &fetchResult{url, body, urls, err, depth}
	}
	go fetchClosure(url, depth)
	cache.content[url] = nil
	for i := 1; i > 0; i-- {
		res := <-funcContextChannel
		if res.err != nil {
			fmt.Println(res.err)
			continue
		}
		fmt.Printf("found: %s %q\n", res.url, res.body)
		if res.depth > 0 {
			for _, entry := range res.urls {
				if _, presented := cache.content[entry]; !presented {
					cache.mu.Lock()
					i++
					go fetchClosure(entry, res.depth-1)
					cache.content[entry] = nil
					cache.mu.Unlock()
				}
			}
		}
	}
	close(funcContextChannel)
	cacheCheckChannel <- &cache.content
	close(cacheCheckChannel)
}

func main() {
	var mainChannel = make(chan *map[string]interface{})
	go Crawl("https://golang.org/", 4, fetcher, mainChannel)
	for i := range mainChannel {
		fmt.Println(i)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
