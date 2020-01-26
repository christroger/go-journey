package main

/*
There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times
The philosophers pick up the chopsticks in any order, not lowest-numbered first
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

*/

import(
	"sync"
	"fmt"
)

type ChopS struct { sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	number int
	permission chan int
	finished chan int
}

func main() {
	chopsticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(ChopS)
	}

	permission := make(chan int, 2)
	finished := make(chan int, 2)

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{chopsticks[i], chopsticks[(i+1)%5], i+1, permission, finished}
	}

	go host(permission, finished)
	var wg sync.WaitGroup
	for i := 0; i<5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg)
	}
	wg.Wait()
}

func (p Philo) eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		<- p.permission
		p.leftCS.Lock()
		p.rightCS.Lock()
	
		fmt.Printf("starting to eat %v\n", p.number)
		fmt.Printf("finished eating %v\n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		p.finished <- 1
	}

	wg.Done()
}

func host(permission, finished chan int) {
	permission <- 1
	permission <- 1
	for {
		<- finished
		permission <- 1
	}
}