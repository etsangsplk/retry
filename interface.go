package retry

import (
	"time"
)

// Set this to override how now is discovered and how sleeping is done
// This is mostly useful for testing, but you never know
var TimeFunc = time.Now
var SleepFunc = time.Sleep

// This is the main interface around which this library is
// built.  It defines a very simple interface for abstracting retry
// logic in your application.
type RetryStrategy interface {
	Next() bool
	HasNext() bool
}

// Retry strategy expanded with reset functionality.
type ResettableRetryStrategy interface {
	RetryStrategy
	Reset()
}

// Useful helper method.  Calls action until it returns true or
// the retry strategy returns false.
func Do(strategy RetryStrategy, action func() bool) bool {
	for strategy.Next() {
		if succ := action(); succ {
			return succ
		}
	}
	return false
}
