package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	go player("A", court)
	go player("B", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Printf("Winner: %s\n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Fail: %s\n", name)

			close(court)
			return
		}

		fmt.Printf("%s hit %d's ball!", name, ball)
		ball++

		court <- ball
	}
}
