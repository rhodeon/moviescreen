package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// MarshalJSON of Runtime appends the time unit in "mins" as a suffix,
// appending "min" if the runtime is 1.
func (r Runtime) MarshalJSON() ([]byte, error) {
	var runtime string
	if r == 1 {
		runtime = fmt.Sprintf("%d min", r)
	} else {
		runtime = fmt.Sprintf("%d mins", r)
	}

	quotedRuntime := strconv.Quote(runtime)
	return []byte(quotedRuntime), nil
}

// UnmarshalJSON trims out the "min" or "mins" suffix and sets
// the remaining integer as the Runtime.
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// remove quotes from JSON value
	unquotedRuntime, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return err
	}

	// separate the digit and the unit
	parts := strings.Fields(unquotedRuntime)
	if len(parts) != 2 || (parts[1] != "min" && parts[1] != "mins") {
		return ErrInvalidRuntimeFormat
	}

	// convert digit to int and set the value as Runtime
	runtime, err := strconv.Atoi(parts[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(runtime)
	return nil
}
