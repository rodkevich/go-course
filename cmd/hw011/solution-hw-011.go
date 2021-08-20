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
func Crawl(url string, depth int, fetcher Fetcher, mainChannel chan *map[string]interface{}) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start).Nanoseconds())
	}()
	time.Sleep(time.Nanosecond * 2)

	type memo struct {
		*sync.Mutex
		content map[string]interface{}
	}
	type fetchResult struct {
		url   string
		err   error
		depth int
		*fakeResult
	}

	inFunctionChannel := make(chan *fetchResult)
	cache := &memo{
		Mutex:   &sync.Mutex{},
		content: map[string]interface{}{},
	}

	fetchClosure := func(url string, depth int) {
		body, urls, err := fetcher.Fetch(url)
		inFunctionChannel <- &fetchResult{
			url:   url,
			err:   err,
			depth: depth,
			fakeResult: &fakeResult{
				body: body,
				urls: urls,
			},
		}
	}

	var alreadyMet interface{} = nil // alias just for readability
	// deal with a first fetch
	go fetchClosure(url, depth)
	cache.content[url] = alreadyMet

	for i := 1; i > 0; i-- {
		res := <-inFunctionChannel
		if res.err != nil {
			fmt.Println(res.err)
			continue
		}
		fmt.Printf("found: %s %q\n", res.url, res.body)
		if res.depth > 0 {
			for _, entry := range res.urls {
				if _, presented := cache.content[entry]; !presented {
					cache.Lock()
					i++
					go fetchClosure(entry, res.depth-1)
					cache.content[entry] = alreadyMet
					cache.Unlock()
				}
			}
		}
	}
	close(inFunctionChannel)
	mainChannel <- &cache.content
	close(mainChannel)
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
