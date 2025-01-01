package retry

import (
	"context"
	"fmt"
	"time"
)

type Effector func(ctx context.Context) (string, error)

func Retryable(effector Effector, retries int, delay time.Duration) Effector {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			fmt.Printf("Attempt %d of %d failed. Trying again in %v\n", r+1, retries, delay)

			select {
			case <-time.After(delay): // Sleep for our delay
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}
