package workerpool

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "running job ", j)
		time.Sleep(time.Second)
		results <- rand.Int() // Place a random int to result set
	}
}
