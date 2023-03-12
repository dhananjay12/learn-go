/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
)

const numPhilosophers = 5
const numEatsPerPhilosopher = 3

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *Chopstick
	host                          chan bool
}

func (p Philosopher) eat() {
	for i := 0; i < numEatsPerPhilosopher; i++ {
		// Request permission from the host to eat (blocks until permission is granted)
		<-p.host

		// Pick up chopsticks in any order
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Println("starting to eat", p.id)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		fmt.Println("finishing eating", p.id)

		// Release permission to the host to allow another philosopher to eat (unblocks the host)
		p.host <- true
	}
}

func main() {
	// Create chopsticks
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// Create host channel to allow only two philosophers to eat concurrently
	host := make(chan bool, 2)

	// Create philosophers
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
			host:           host,
		}
	}

	// Start philosophers
	var wg sync.WaitGroup
	wg.Add(numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		// Use a closure to pass the current philosopher to the goroutine (otherwise all goroutines will use the last philosopher)
		go func(p *Philosopher) {
			p.eat()
			wg.Done()
		}(philosophers[i])
	}

	// Allow first two philosophers to start eating
	host <- true
	host <- true

	wg.Wait()
}
