go-backoff
======================================================================

Package backoff provides wrapper functions for "[github.com/cenkalti/backoff/v4](https://pkg.go.dev/github.com/cenkalti/backoff/v4)".


examples
----------------------------------------------------------------------

```go
	mockAPI := func(ctx context.Context, name string) (string, error) { return name, nil }

	result, err := backoff.RetryContextR2(
		ctx,
		bind.P2R2(mockAPI, ctx, "hello"),
		backoff.MaxInterval(7*time.Second),
		backoff.Stop(7*time.Second),
		backoff.MaxRetries(7),
	)
```

see https://pkg.go.dev/github.com/takumakei/go-bind .
