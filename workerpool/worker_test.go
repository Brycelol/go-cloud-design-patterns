package workerpool

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	jobs := make(chan int, 20)
	results := make(chan int)
	wg := sync.WaitGroup{}

	// Spawn 6 workers
	for i := 0; i < 6; i++ {
		go worker(i, jobs, results)
	}

	// Time for work
	for i := 0; i < 20; i++ {
		wg.Add(1)
		jobs <- i
	}

	go func() { // Close channels once we've read everything
		wg.Wait()
		close(jobs)
		close(results)
	}()

	for result := range results { // Grab results and decrement our wait group.
		fmt.Println("Received result from our worker pool:", result)
		wg.Done()
	}
}
