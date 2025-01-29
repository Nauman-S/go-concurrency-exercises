//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	mu.Lock()
	c := time.Tick(1 * time.Second)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		<-c
		mu.Unlock()
		return
	}
	<-c
	mu.Unlock()

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently

		val := u
		go Crawl(val, depth-1, wg, mu)
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex = sync.Mutex{}

	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg, &mu)
	wg.Wait()
}
