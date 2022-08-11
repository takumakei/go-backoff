package backoff

import "github.com/cenkalti/backoff/v4"

// Default values for ExponentialBackOff.
const (
	DefaultInitialInterval     = backoff.DefaultInitialInterval
	DefaultRandomizationFactor = backoff.DefaultRandomizationFactor
	DefaultMultiplier          = backoff.DefaultMultiplier
	DefaultMaxInterval         = backoff.DefaultMaxInterval
	DefaultMaxElapsedTime      = backoff.DefaultMaxElapsedTime
	DefaultStop                = backoff.Stop
)
