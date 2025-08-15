package main


import (
	"fmt"
	"sync"
)


type SafeCounter struct {
	v map[string]bool
	mux sync.Mutex
}


type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}


var cnt SafeCounter = SafeCounter{v: make(map[string]bool)}


func Crawl(url string, depth int, fetcher Fetcher, exit chan bool) {

	if depth <= 0 {
		exit <- true
		return
	}
	
	cnt.mux.Lock()
	_, ok := cnt.v[url]
	if ok == false {
		cnt.v[url] = true
		cnt.mux.Unlock()
	} else {
		exit <- true
		cnt.mux.Unlock()
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {	
		fmt.Println(err)
		exit <- true
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	e := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, e)
	}

	for i := 0; i < len(urls); i++ {
		<-e
	}
	exit <- true
}

func main() {
	exit := make(chan bool)
	Crawl("https://golang.org/", 4, fetcher, exit)
	<-exit
}

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

var fetcher = fakeFetcher{
	"https://golang.org": &fakeResult {
		"The Go Programming Language", 
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg": &fakeResult {
		"Packages", 
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult {
		"Package fmt", 
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult {
		"Package fmt", 
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
