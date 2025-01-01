package funnel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFunnel(t *testing.T) {
	var sources []<-chan int

	for i := 0; i < 10; i++ {
		sourceChan := make(chan int)
		sources = append(sources, sourceChan)

		go func() { // Throw some junk data into our channels to simulate input
			defer close(sourceChan)

			for i := 0; i < 3; i++ {
				sourceChan <- rand.Int()
				time.Sleep(time.Second)
			}
		}()
	}

	destChan := Funnel(sources...)

	// Drain our destination channel
	for output := range destChan {
		fmt.Printf("We received int %d from our destination channel!\n", output)
	}
}
