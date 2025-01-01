package retry

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFailedRetryable(t *testing.T) {

	retry := Retryable(func(ctx context.Context) (string, error) { return "", errors.New("failed") }, 5, 1)

	_, err := retry(context.Background())

	assert.EqualError(t, err, "failed")
}
