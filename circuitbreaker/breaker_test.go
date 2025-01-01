package circuitbreaker

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreakerThreshold(t *testing.T) {
	badFunction := Breaker(func(ctx context.Context) (string, error) { return "", errors.New("failed") }, 1)

	_, errA := badFunction(context.Background())
	_, errB := badFunction(context.Background())

	assert.EqualError(t, errA, "failed")
	assert.EqualError(t, errB, CircuitBrokenMessage)
}

func TestBreakerGoodInvocation(t *testing.T) {
	goodFunction := Breaker(func(ctx context.Context) (string, error) { return "good", nil }, 1)

	result, _ := goodFunction(context.Background())

	assert.Equal(t, "good", result)
}
