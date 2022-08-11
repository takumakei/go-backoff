package backoff

import (
	"context"

	"github.com/cenkalti/backoff/v4"
)

// New creates an [backoff.ExponentialBackOff] by calling [backoff.NewExponentialBackOff], applies options to it, returns it as [backoff.BackOff].
func New(options ...Option) backoff.BackOff {
	return Apply(backoff.NewExponentialBackOff(), options...)
}

// NewContext creates an [backoff.ExponentialBackOff] by calling [backoff.NewExponentialBackOff], applies options to it, wraps by [backoff.WithContext] with ctx, returns it as [backoff.BackOff].
func NewContext(ctx context.Context, options ...Option) backoff.BackOff {
	return backoff.WithContext(New(options...), ctx)
}
