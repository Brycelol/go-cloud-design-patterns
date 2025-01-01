package circuitbreaker

import (
	"context"
	"errors"
	"sync"
	"time"
)

const CircuitBrokenMessage = "service is currently unavailable"

type Circuit func(ctx context.Context) (string, error)

func Breaker(circuit Circuit, threshold int) Circuit {
	var failures int
	var last = time.Now()
	var mutex sync.RWMutex

	return func(ctx context.Context) (string, error) {
		mutex.RLock()

		d := failures - threshold

		if d >= 0 {
			shouldRetryAt := last.Add((2 << d) * time.Second)

			if !time.Now().After(shouldRetryAt) {
				mutex.RUnlock()
				return "", errors.New(CircuitBrokenMessage)
			}
		}

		mutex.RUnlock()

		response, err := circuit(ctx) // Issue underlying request

		mutex.Lock() // Lock shared resources
		defer mutex.Unlock()

		last = time.Now() // retrieve time of attempt

		if err != nil {
			failures++
			return response, err
		}

		failures = 0 // We've had a successful invocation, reset counter.

		return response, nil
	}
}
