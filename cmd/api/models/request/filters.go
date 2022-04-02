package request

// Filters represents the filters as query parameters in a request url.
type Filters struct {
	Page  int
	Limit int
	Sort  string
}
