package request

import (
	"github.com/rhodeon/moviescreen/internal/validator"
	"github.com/rhodeon/moviescreen/internal/validator/rules"
	"strings"
)

// Filters represents the filters as query parameters in a request url.
type Filters struct {
	Page       int
	Limit      int
	Sort       string
	ValidSorts []string
}

const (
	FilterFieldPage  = "page"
	FilterFieldLimit = "limit"
	FilterFieldSort  = "sort"
)

func (f Filters) Validate() *validator.Validator {
	v := validator.New("filter")

	// check that the page and page_size parameters contain sensible values
	v.Check(f.Page > 0, FilterFieldPage, "must be greater than zero")
	v.Check(f.Page <= 10_000_000, FilterFieldPage, "must be a maximum of 10 million")
	v.Check(f.Limit > 0, FilterFieldLimit, "must be greater than zero")
	v.Check(f.Limit <= 100, FilterFieldLimit, "must be a maximum of 100")

	// check that the sort parameter matches a value in the valid list
	v.Check(rules.In(strings.TrimPrefix(f.Sort, "-"), f.ValidSorts), FilterFieldSort, "invalid sort value")

	return v
}

// SortColumn checks that the base form of the sort filter exists in the list of valid sorts,
// and returns the base form if so.
// This is to enable passing in a valid column name as an SQL order.
func (f Filters) SortColumn(defaultSort string) string {
	baseSort := strings.TrimPrefix(f.Sort, "-")
	for _, validSort := range f.ValidSorts {
		if validSort == baseSort {
			return baseSort
		}
	}
	// this should never occur due to previous sort validation
	return defaultSort
}

// SortDirection returns the order direction of an SQL query ("ASC" or "DESC")
// based on the prefix of the sort filter.
func (f Filters) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

// Offset calculates and returns the position of the data to start selecting from.
// It depends on the page number and limit.
func (f Filters) Offset() int {
	return (f.Page - 1) * f.Limit
}
