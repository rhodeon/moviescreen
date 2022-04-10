package common

import (
	"github.com/rhodeon/prettylog"
	"sync"
)

// Background launches operations in a goroutine and handles any panic that occurs.
func Background(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)

	go func() {
		// decrement waitgroup when the background task is complete
		defer wg.Done()

		// run a deferred function which catches any panic in the goroutine
		defer func() {
			if err := recover(); err != nil {
				prettylog.Error(err)
			}
		}()

		// execute the function
		fn()
	}()
}
