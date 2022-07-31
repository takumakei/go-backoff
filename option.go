package backoff

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Option is type of setting parameters of ExponentialBackOff.
type Option func(*builder)

type builder struct {
	exp *backoff.ExponentialBackOff
	max *uint64
	ctx context.Context
}

func (bu *builder) build() (b backoff.BackOff) {
	b = bu.exp
	if bu.max != nil {
		b = backoff.WithMaxRetries(b, *bu.max)
	}
	if bu.ctx != nil {
		b = backoff.WithContext(b, bu.ctx)
	}
	return
}

// InitialInterval uses d as InitialInterval.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
func InitialInterval(d time.Duration) Option {
	return func(bu *builder) { bu.exp.InitialInterval = d }
}

// RandomizationFactor uses f as RandomizationFactor.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
func RandomizationFactor(f float64) Option {
	return func(bu *builder) { bu.exp.RandomizationFactor = f }
}

// Multiplier uses f as Multiplier.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
func Multiplier(f float64) Option {
	return func(bu *builder) { bu.exp.Multiplier = f }
}

// MaxInterval uses d as MaxInterval.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
func MaxInterval(d time.Duration) Option {
	return func(bu *builder) { bu.exp.MaxInterval = d }
}

// MaxElapsedTime uses t as MaxElapsedTime.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
func MaxElapsedTime(d time.Duration) Option {
	return func(bu *builder) { bu.exp.MaxElapsedTime = d }
}

// Stop uses d as Stop.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
func Stop(d time.Duration) Option {
	return func(bu *builder) { bu.exp.Stop = d }
}

// Clock uses clock as Clock.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
func Clock(clock backoff.Clock) Option {
	return func(bu *builder) { bu.exp.Clock = clock }
}

// Context applies backoff.WithContext with ctx.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#WithContext
func Context(ctx context.Context) Option {
	return func(bu *builder) { bu.ctx = ctx }
}

// MaxRetries applies backoff.WithMaxRetries with max.
//
// see: https://pkg.go.dev/github.com/cenkalti/backoff/v4#WithMaxRetries
func MaxRetries(max uint64) Option {
	return func(bu *builder) { bu.max = &max }
}
