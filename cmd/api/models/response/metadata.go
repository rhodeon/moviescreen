package response

import "math"

// Metadata holds the response metadata.
type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageLimit    int `json:"page_limit,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

// CalculateMetadata generates a metadata from the given page, limit
// and total number of records.
func CalculateMetadata(page int, pageLimit int, totalRecords int) Metadata {
	// return an empty metadata if no records are found
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageLimit:    pageLimit,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageLimit))),
		TotalRecords: totalRecords,
	}
}
