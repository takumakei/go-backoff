package backoff

import (
	"github.com/cenkalti/backoff/v4"
)

// Apply applies options to exp, returns it as [backoff].[BackOff].
//
// [backoff]: https://pkg.go.dev/github.com/cenkalti/backoff/v4
// [BackOff]: https://pkg.go.dev/github.com/cenkalti/backoff/v4#BackOff
func Apply(exp *backoff.ExponentialBackOff, options ...Option) backoff.BackOff {
	bu := &builder{exp: exp}
	for _, opt := range options {
		opt(bu)
	}
	return bu.build()
}
