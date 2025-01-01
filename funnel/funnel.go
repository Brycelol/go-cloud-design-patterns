package funnel

import "sync"

// Funnel will multiplex multiple input channels onto a single output channel.
func Funnel(sources ...<-chan int) <-chan int {
	dest := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(len(sources))

	// Iterate our sources
	for _, sChan := range sources {
		go func(<-chan int) {
			defer wg.Done()
			for sChanVal := range sChan { // spool channel until it is closed.
				dest <- sChanVal
			}
		}(dest)
	}

	// Close destination when sources are closed.
	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}
