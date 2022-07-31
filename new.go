package backoff

import (
	"github.com/cenkalti/backoff/v4"
)

// New creates an backoff.[ExponentialBackOff], applies options to it, returns it as [backoff].[BackOff].
//
// [ExponentialBackOff]: https://pkg.go.dev/github.com/cenkalti/backoff/v4#ExponentialBackOff
// [backoff]: https://pkg.go.dev/github.com/cenkalti/backoff/v4
// [BackOff]: https://pkg.go.dev/github.com/cenkalti/backoff/v4#BackOff
func New(options ...Option) backoff.BackOff {
	return Apply(backoff.NewExponentialBackOff(), options...)
}
