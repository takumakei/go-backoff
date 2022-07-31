// Package backoff provides wrapper functions for "github.com/cenkalti/backoff/v4".
package backoff

import (
	"github.com/cenkalti/backoff/v4"
	"github.com/takumakei/go-bind"
)

// RetryP0R1 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has zero parameter. (P0)
//  - fn returns one value of error. (R1)
func RetryP0R1(fn func() error, options ...Option) error {
	return backoff.Retry(fn, New(options...))
}

// RetryP0R2 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has zero parameter. (P0)
//  - fn returns 2 values. (R2)
func RetryP0R2[R0 any](fn func() (R0, error), options ...Option) (r0 R0, err error) {
	err = backoff.Retry(
		func() (err error) {
			r0, err = fn()
			return
		},
		New(options...),
	)
	return
}

// RetryP0R3 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has zero parameter. (P0)
//  - fn returns 3 values. (R3)
func RetryP0R3[R0, R1 any](fn func() (R0, R1, error), options ...Option) (r0 R0, r1 R1, err error) {
	err = backoff.Retry(
		func() (err error) {
			r0, r1, err = fn()
			return
		},
		New(options...),
	)
	return
}

// RetryP1R1 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has one parameter. (P1)
//  - fn returns one value. (R1)
func RetryP1R1[P0 any](fn func(P0) error, p0 P0, options ...Option) error {
	return RetryP0R1(bind.P1R1(fn, p0), options...)
}

// RetryP1R2 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has one parameter. (P1)
//  - fn returns 2 values. (R2)
func RetryP1R2[P0, R0 any](fn func(P0) (R0, error), p0 P0, options ...Option) (R0, error) {
	return RetryP0R2(bind.P1R2(fn, p0), options...)
}

// RetryP1R3 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has one parameter. (P1)
//  - fn returns 3 values. (R3)
func RetryP1R3[P0, R0, R1 any](fn func(P0) (R0, R1, error), p0 P0, options ...Option) (R0, R1, error) {
	return RetryP0R3(bind.P1R3(fn, p0), options...)
}

// RetryP2R1 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has 2 parameters. (P2)
//  - fn returns one value. (R1)
func RetryP2R1[P0, P1 any](fn func(P0, P1) error, p0 P0, p1 P1, options ...Option) error {
	return RetryP0R1(bind.P2R1(fn, p0, p1), options...)
}

// RetryP2R2 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has 2 parameters. (P2)
//  - fn returns 2 values. (R2)
func RetryP2R2[P0, P1, R0 any](fn func(P0, P1) (R0, error), p0 P0, p1 P1, options ...Option) (R0, error) {
	return RetryP0R2(bind.P2R2(fn, p0, p1), options...)
}

// RetryP2R3 calls fn under backoff.Retry with ExponentialBackOff with options.
//
//  - fn has 2 parameters. (P2)
//  - fn returns 3 values. (R3)
func RetryP2R3[P0, P1, R0, R1 any](fn func(P0, P1) (R0, R1, error), p0 P0, p1 P1, options ...Option) (R0, R1, error) {
	return RetryP0R3(bind.P2R3(fn, p0, p1), options...)
}
