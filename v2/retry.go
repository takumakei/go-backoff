package backoff

import (
	"context"

	"github.com/cenkalti/backoff/v4"
)

// Retry the function fn until it does not return error or BackOff stops.
//
// BackOff is created by [New] with options.
func Retry(fn func() error, options ...Option) error {
	return backoff.Retry(fn, New(options...))
}

// RetryContext the function fn until it does not return error or BackOff stops.
//
// BackOff is created by [NewContext] with options and ctx.
func RetryContext(ctx context.Context, fn func() error, options ...Option) error {
	return backoff.Retry(fn, NewContext(ctx, options...))
}

// RetryR1 is an alias of [Retry].
func RetryR1(fn func() error, options ...Option) error {
	return Retry(fn, options...)
}

// RetryContextR1 is an alias of [RetryContext].
func RetryContextR1(ctx context.Context, fn func() error, options ...Option) error {
	return RetryContext(ctx, fn, options...)
}

// RetryR2 the function fn that returns 2 values until it does not return error or BackOff stops.
func RetryR2[R1 any](fn func() (R1, error), options ...Option) (r1 R1, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, err = fn()
			return
		},
		New(options...),
	)
	return
}

// RetryContextR2 the function fn that returns 2 values until it does not return error or BackOff stops.
func RetryContextR2[R1 any](ctx context.Context, fn func() (R1, error), options ...Option) (r1 R1, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, err = fn()
			return
		},
		NewContext(ctx, options...),
	)
	return
}

// RetryR3 the function fn that returns 3 values until it does not return error or BackOff stops.
func RetryR3[R1, R2 any](fn func() (R1, R2, error), options ...Option) (r1 R1, r2 R2, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, r2, err = fn()
			return
		},
		New(options...),
	)
	return
}

// RetryContextR3 the function fn that returns 3 values until it does not return error or BackOff stops.
func RetryContextR3[R1, R2 any](ctx context.Context, fn func() (R1, R2, error), options ...Option) (r1 R1, r2 R2, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, r2, err = fn()
			return
		},
		NewContext(ctx, options...),
	)
	return
}

// RetryR4 the function fn that returns 4 values until it does not return error or BackOff stops.
func RetryR4[R1, R2, R3 any](fn func() (R1, R2, R3, error), options ...Option) (r1 R1, r2 R2, r3 R3, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, r2, r3, err = fn()
			return
		},
		New(options...),
	)
	return
}

// RetryContextR4 the function fn that returns 4 values until it does not return error or BackOff stops.
func RetryContextR4[R1, R2, R3 any](ctx context.Context, fn func() (R1, R2, R3, error), options ...Option) (r1 R1, r2 R2, r3 R3, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, r2, r3, err = fn()
			return
		},
		NewContext(ctx, options...),
	)
	return
}

// RetryR5 the function fn that returns 5 values until it does not return error or BackOff stops.
func RetryR5[R1, R2, R3, R4 any](fn func() (R1, R2, R3, R4, error), options ...Option) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, r2, r3, r4, err = fn()
			return
		},
		New(options...),
	)
	return
}

// RetryContextR5 the function fn that returns 5 values until it does not return error or BackOff stops.
func RetryContextR5[R1, R2, R3, R4 any](ctx context.Context, fn func() (R1, R2, R3, R4, error), options ...Option) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
	err = backoff.Retry(
		func() (err error) {
			r1, r2, r3, r4, err = fn()
			return
		},
		NewContext(ctx, options...),
	)
	return
}
