// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Cache safely keeps track of requested and fetched urls preventing
// multiple downloads. If an url (key) is present with a false boolean
// value, it hasn't been fetched.
type Cache struct {
	v map[string]bool
	m sync.Mutex
}

// Crawler uses fetcher to crawl pages starting with url, to a maximum of depth.
// It spawns as many concurrent crawlers as needed and waits for their completion
// in a WaitGroup. The cache garantees that each url is only fetched once.
func Crawler(url string, depth int, fetcher Fetcher) {
	group := &sync.WaitGroup{}
	cache := &Cache{v: map[string]bool{}}

	cache.Add(url) // update cache.
	group.Add(1)   // register in the group.
	go crawl(url, depth, fetcher, cache, group)
	group.Wait() // wait here for completion of all the spawned crawlers.

	return
}

// Cache().IsCached safely returns if the given url is in the cache.
func (c *Cache) IsCached(u string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	_, ok := c.v[u] // is u a key to c.v?

	return ok
}

// Cache().Add safely caches the given url.
func (c *Cache) Add(u string) {
	c.m.Lock()
	defer c.m.Unlock()

	c.v[u] = false // u is a key to c.v
}

// Cache().Mark safely marks the given url as fetched.
func (c *Cache) Mark(u string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	_, ok := c.v[u]
	if ok {
		c.v[u] = true
	}
	return ok
}

// crawl recursively spawns itself into more concurrent crawls.
// All crawls register in a WaitGroup.
func crawl(url string, depth int, fetcher Fetcher, cache *Cache, group *sync.WaitGroup) {
	defer group.Done() // handy! Code returns in various places.

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url) // fetch fake news.
	if err != nil {
		fmt.Println(err)
		return
	}
	cache.Mark(url) // update cache status.
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls { // spawn one crawler per url.
		if !cache.IsCached(u) {
			cache.Add(u) // update cache.
			group.Add(1) // register in the group.
			go crawl(u, depth-1, fetcher, cache, group)
			// no waiting here!
		}
	}

	return
}

func main() {
	start := time.Now()
	
	Crawler("https://golang.org/", 4, fetcher)
	
	elapsed := time.Since(start)
	fmt.Println("Crawler took ", elapsed)

	return
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
