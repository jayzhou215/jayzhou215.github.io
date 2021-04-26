package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	mu          sync.Mutex
	fetchedUrls map[string]interface{}
	wg          *sync.WaitGroup
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (cr *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if cr.setNxUrl(u) {
			cr.wg.Add(1)
			go func() {
				defer cr.wg.Done()
				cr.Crawl(u, depth-1, fetcher)
			}()
		}
	}
	return
}

func (cr *Crawler) setNxUrl(u string) bool {
	cr.mu.Lock()
	defer func() {
		cr.mu.Unlock() // 仅对map数组的操作进行lock, unlock
	}()
	if _, ok := cr.fetchedUrls[u]; ok {
		return false
	}
	cr.fetchedUrls[u] = nil
	return true
}

func main() {
	cr := Crawler{
		mu:          sync.Mutex{},
		fetchedUrls: make(map[string]interface{}),
		wg:          &sync.WaitGroup{},
	}
	cr.Crawl("https://golang.org/", 4, fetcher)
	cr.wg.Wait()
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
