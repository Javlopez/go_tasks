package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	count := make(chan int)

	wg.Add(2)

	fmt.Println("Starting the goroutines\n")

	go printCounts("A", count)

	go printCounts("B", count)

	fmt.Println("Channel begin")

	count <- 1

	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("Finished the Execution")

}

func printCounts(label string, count chan int) {

	defer wg.Done()

	for {
		//receive message from channel
		val, ok := <-count

		if !ok {
			fmt.Println("The channel was closed")
			return
		}

		fmt.Printf("Count: %d received from %s \n", val, label)

		if val == 10 {
			fmt.Printf("Channel closed from %s\n", label)
			//closing channel
			close(count)
			return
		}
		val++
		//send count to back to another goroutine
		count <- val
	}
}
