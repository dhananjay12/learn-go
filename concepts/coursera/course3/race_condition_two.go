/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
*/
package main

import (
	"fmt"
	"time"
)

func printWord(s string) {
	for _, c := range s {
		fmt.Print(string(c), " ")
		time.Sleep(10 * time.Millisecond)
	}
}

// Script containing two goroutines.

func main() {
	// run two different goroutine
	go printWord("Hello")
	go printWord("World")

	// prevent program termination before goroutines return
	time.Sleep(1 * time.Second)
}
