package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Star goroutines")

	go printCounts("A")

	go printCounts("B")

	fmt.Println("Waiting to Finish")

	wg.Wait()
	fmt.Println("\nFinished the Execution")
}

func printCounts(label string) {

	defer wg.Done()

	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}
