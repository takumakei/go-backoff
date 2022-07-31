package backoff_test

import (
	"context"
	"fmt"
	"time"

	backoff "github.com/takumakei/go-backoff"
)

func Example() {
	// mockAPI takes 2 parameters, returns 2 values.
	mockAPI := func(ctx context.Context, name string) (string, error) { return name, nil }

	ctx := context.Background()

	// - P2 means that the target function takes 2 parameters.
	// - R2 means that the target function returns 2 values.
	// - types of result and err is inferred by the target function.
	result, err := backoff.RetryP2R2(
		// the target function to retry when an error occurs
		mockAPI,
		// the 1st argument for the target function
		ctx,
		// the 2nd argument for the target function
		"hello",

		// parameters for ExponentialBackOff
		backoff.MaxInterval(7*time.Second),
		backoff.Stop(7*time.Second),
		backoff.Context(ctx),
		backoff.MaxRetries(7),
	)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Println(result)
	}

	// Output: hello
}
