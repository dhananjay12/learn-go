/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
*/
package main

import (
	"fmt"
	"time"
)

var x int

// Containe 2 goroutines and main executes these concurrently
// 1st goroutine increments x
// 2nd goroutine prints x
// If the 2nd goroutine is executed before the 1st goroutine, the value of x will be 0
// If the 2nd goroutine is executed after the 1st goroutine, the value of x will be 1

// Run program multiple times , it will print 0 and 1 randomly
func main() {
	go func() {
		x++
	}()
	go func() {
		fmt.Println("x = ", x)
	}()
	// Sleep for 1 second to allow the goroutines to execute
	time.Sleep(1 * time.Second)
}
