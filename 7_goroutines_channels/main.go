package main

import (
	"fmt"
	"sync"
	timepkg "time"
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
  1) Create a channel of strings.
  2) Start a WaitGroup and launch N goroutines that sleep for a random-ish time (use time.Duration).
  3) Each goroutine sends a message like "result from worker i".
  4) Close the channel when all workers are done; range over channel to print results.
*/
func main() {
	// TODO: implement
	_ = sync.WaitGroup{}
	_ = timepkg.Sleep
	fmt.Println("TODO")
}
