package common

import "github.com/rhodeon/prettylog"

// Background launches operations in a goroutine and handles any panic that occurs.
func Background(fn func()) {
	go func() {
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
