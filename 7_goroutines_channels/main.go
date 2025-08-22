package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Problem 7 — Goroutines & Channels (Concurrency)

Goal:

	Spawn N goroutines (e.g., 5) that simulate "fetching" from URLs (sleep & return a string).
	Each goroutine sends its result on a channel. The main goroutine collects and prints them.

Why:

	Teaches goroutines (`go`), channels, and sync.WaitGroup fan-out/fan-in pattern.

Run:

	go run ./7_goroutines_channels

TODOs:
 1. Create a channel of strings.
 2. Start a WaitGroup and launch N goroutines that sleep for a random-ish time (use time.Duration).
 3. Each goroutine sends a message like "result from worker i".
 4. Close the channel when all workers are done; range over channel to print results.
*/
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			task(i)
		}(i)
	}

	wg.Wait()
}

func task(i int) {
	fmt.Printf("Starting task... %d\n", i)
	startTime := time.Now()
	time.Sleep(time.Duration(i))
	endTime := time.Now()
	timeElapsed := endTime.Sub(startTime)
	messages := make(chan string)
	go func() { messages <- fmt.Sprintf("Time elapsed for task %d -- %s", i, timeElapsed) }()
	msg := <-messages
	fmt.Println(msg)

}
