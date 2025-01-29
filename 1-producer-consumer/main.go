//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, c chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(c)
			return
		}
		c <- tweet
	}
}

func consumer(c <-chan *Tweet) {
	for {
		tweet := <-c
		if tweet == nil {
			return
		}
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	var ch chan *Tweet = make(chan *Tweet, 100)

	// Producer
	go producer(stream, ch)

	// Consumer
	consumer(ch)

	fmt.Printf("Process took %s\n", time.Since(start))
}
