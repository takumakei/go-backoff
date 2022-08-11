package backoff_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/takumakei/go-backoff/v2"
	"github.com/takumakei/go-bind"
)

func Example() {
	// mockAPI takes 2 parameters, returns 2 values.
	mockAPI := func(ctx context.Context, name string) (string, error) { return name, nil }

	ctx := context.Background()

	result, err := backoff.RetryContextR2(
		ctx,
		bind.P2R2(mockAPI, ctx, "hello"),
		backoff.MaxInterval(7*time.Second),
		backoff.Stop(7*time.Second),
		backoff.MaxRetries(7),
	)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Println(result)
	}

	// Output: hello
}

func TestRetry(t *testing.T) {
	n := 0
	never := errors.New("never")
	err := backoff.Retry(
		func() error {
			n++
			return never
		},
		backoff.MaxInterval(1),
		backoff.MaxRetries(3),
	)
	assert.ErrorIs(t, err, never)
	// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
	assert.Equal(t, 4, n)
}

func TestRetryContext(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		n := 0
		never := errors.New("never")
		err := backoff.RetryContext(
			context.Background(),
			func() error {
				n++
				return never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.ErrorIs(t, err, never)
		// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
		assert.Equal(t, 4, n)
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // キャンセル

		n := 0
		never := errors.New("never")
		err := backoff.RetryContext(
			ctx,
			func() error {
				n++
				return never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.ErrorIs(t, err, context.Canceled)
		// fn は必ず1回実行される
		assert.Equal(t, 1, n)
	})
}

func TestRetryR1(t *testing.T) {
	n := 0
	never := errors.New("never")
	err := backoff.RetryR1(
		func() error {
			n++
			return never
		},
		backoff.MaxInterval(1),
		backoff.MaxRetries(3),
	)
	assert.ErrorIs(t, err, never)
	// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
	assert.Equal(t, 4, n)
}

func TestRetryContextR1(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		n := 0
		never := errors.New("never")
		err := backoff.RetryContextR1(
			context.Background(),
			func() error {
				n++
				return never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.ErrorIs(t, err, never)
		// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
		assert.Equal(t, 4, n)
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // キャンセル

		n := 0
		never := errors.New("never")
		err := backoff.RetryContextR1(
			ctx,
			func() error {
				n++
				return never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.ErrorIs(t, err, context.Canceled)
		// fn は必ず1回実行される
		assert.Equal(t, 1, n)
	})
}

func TestRetryR2(t *testing.T) {
	n := 0
	never := errors.New("never")
	r, err := backoff.RetryR2(
		func() (int, error) {
			n++
			return 42, never
		},
		backoff.MaxInterval(1),
		backoff.MaxRetries(3),
	)
	assert.ErrorIs(t, err, never)
	assert.Equal(t, 42, r)
	// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
	assert.Equal(t, 4, n)
}

func TestRetryContextR2(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		n := 0
		never := errors.New("never")
		r, err := backoff.RetryContextR2(
			context.Background(),
			func() (int, error) {
				n++
				return 42, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, 42, r)
		assert.ErrorIs(t, err, never)
		// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
		assert.Equal(t, 4, n)
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // キャンセル

		n := 0
		never := errors.New("never")
		r, err := backoff.RetryContextR2(
			ctx,
			func() (int, error) {
				n++
				return 42, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, 42, r)
		assert.ErrorIs(t, err, context.Canceled)
		// fn は必ず1回実行される
		assert.Equal(t, 1, n)
	})
}

func TestRetryR3(t *testing.T) {
	n := 0
	never := errors.New("never")
	s, r, err := backoff.RetryR3(
		func() (string, int, error) {
			n++
			return "hello", 42, never
		},
		backoff.MaxInterval(1),
		backoff.MaxRetries(3),
	)
	assert.Equal(t, "hello", s)
	assert.Equal(t, 42, r)
	assert.ErrorIs(t, err, never)
	// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
	assert.Equal(t, 4, n)
}

func TestRetryContextR3(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		n := 0
		never := errors.New("never")
		s, r, err := backoff.RetryContextR3(
			context.Background(),
			func() (string, int, error) {
				n++
				return "hello", 42, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, "hello", s)
		assert.Equal(t, 42, r)
		assert.ErrorIs(t, err, never)
		// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
		assert.Equal(t, 4, n)
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // キャンセル

		n := 0
		never := errors.New("never")
		s, r, err := backoff.RetryContextR3(
			ctx,
			func() (string, int, error) {
				n++
				return "hello", 42, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, "hello", s)
		assert.Equal(t, 42, r)
		assert.ErrorIs(t, err, context.Canceled)
		// fn は必ず1回実行される
		assert.Equal(t, 1, n)
	})
}

func TestRetryR4(t *testing.T) {
	n := 0
	never := errors.New("never")
	s, r, b, err := backoff.RetryR4(
		func() (string, int, bool, error) {
			n++
			return "hello", 42, true, never
		},
		backoff.MaxInterval(1),
		backoff.MaxRetries(3),
	)
	assert.Equal(t, "hello", s)
	assert.Equal(t, 42, r)
	assert.True(t, b)
	assert.ErrorIs(t, err, never)
	// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
	assert.Equal(t, 4, n)
}

func TestRetryContextR4(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		n := 0
		never := errors.New("never")
		s, r, b, err := backoff.RetryContextR4(
			context.Background(),
			func() (string, int, bool, error) {
				n++
				return "hello", 42, true, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, "hello", s)
		assert.Equal(t, 42, r)
		assert.True(t, b)
		assert.ErrorIs(t, err, never)
		// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
		assert.Equal(t, 4, n)
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // キャンセル

		n := 0
		never := errors.New("never")
		s, r, b, err := backoff.RetryContextR4(
			ctx,
			func() (string, int, bool, error) {
				n++
				return "hello", 42, true, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, "hello", s)
		assert.Equal(t, 42, r)
		assert.True(t, b)
		assert.ErrorIs(t, err, context.Canceled)
		// fn は必ず1回実行される
		assert.Equal(t, 1, n)
	})
}

func TestRetryR5(t *testing.T) {
	n := 0
	never := errors.New("never")
	s, r, b, i, err := backoff.RetryR5(
		func() (string, int, bool, complex128, error) {
			n++
			return "hello", 42, true, 3 + 4i, never
		},
		backoff.MaxInterval(1),
		backoff.MaxRetries(3),
	)
	assert.Equal(t, "hello", s)
	assert.Equal(t, 42, r)
	assert.True(t, b)
	assert.Equal(t, 3+4i, i)
	assert.ErrorIs(t, err, never)
	// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
	assert.Equal(t, 4, n)
}

func TestRetryContextR5(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		n := 0
		never := errors.New("never")
		s, r, b, i, err := backoff.RetryContextR5(
			context.Background(),
			func() (string, int, bool, complex128, error) {
				n++
				return "hello", 42, true, 3 + 4i, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, "hello", s)
		assert.Equal(t, 42, r)
		assert.True(t, b)
		assert.Equal(t, 3+4i, i)
		assert.ErrorIs(t, err, never)
		// fn は必ず1回実行され、err != nil ならば最大 MaxRetries 回リトライ実行する.
		assert.Equal(t, 4, n)
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // キャンセル

		n := 0
		never := errors.New("never")
		s, r, b, i, err := backoff.RetryContextR5(
			ctx,
			func() (string, int, bool, complex128, error) {
				n++
				return "hello", 42, true, 3 + 4i, never
			},
			backoff.MaxInterval(1),
			backoff.MaxRetries(3),
		)
		assert.Equal(t, "hello", s)
		assert.Equal(t, 42, r)
		assert.True(t, b)
		assert.Equal(t, 3+4i, i)
		assert.ErrorIs(t, err, context.Canceled)
		// fn は必ず1回実行される
		assert.Equal(t, 1, n)
	})
}
