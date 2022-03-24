package response

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Id      int      `json:"id,omitempty"`
	Title   string   `json:"title,omitempty"`
	Year    int      `json:"year,omitempty"`
	Runtime Runtime  `json:"runtime,omitempty"`
	Genres  []string `json:"genres,omitempty"`
	Version int      `json:"version,omitempty"`
}

type Runtime int

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
