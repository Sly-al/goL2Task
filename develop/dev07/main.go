package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	merged := make(chan interface{})
	wg.Add(len(channels))

	output := func(inp <-chan interface{}) {
		defer wg.Done()
		for now := range inp {
			merged <- now
		}
	}
	for _, optChan := range channels {
		go output(optChan)
	}
	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(sig(10*time.Second), sig(4*time.Second), sig(time.Second))
	fmt.Printf("fone after %v", time.Since(start))
}
