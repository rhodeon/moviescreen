package mock

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"sort"
	"strings"
)

// caseInsensitiveSubslice checks if the target slice contains the data slice.
func caseInsensitiveSubslice(data []string, target []string) bool {
	if len(data) > len(target) {
		return false
	}
	for _, element := range data {
		if !caseInsensitiveIn(element, target) {
			return false
		}
	}
	return true
}

// In returns true if the data is found in the target list.
// Both values are converted to lowercase for comparison.
func caseInsensitiveIn(data string, target []string) bool {
	for _, element := range target {
		if strings.EqualFold(element, data) {
			return true
		}
	}
	return false
}

func sortMoviesById(movies models.Movies, ascending bool) {
	if ascending {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Id < movies[j].Id
		})
	} else {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Id > movies[j].Id
		})
	}
}

func sortMoviesByTitle(movies models.Movies, ascending bool) {
	if ascending {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Title < movies[j].Title
		})
	} else {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Title > movies[j].Title
		})
	}
}

func sortMoviesByYear(movies models.Movies, ascending bool) {
	if ascending {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Year < movies[j].Year
		})
	} else {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Year > movies[j].Year
		})
	}
}

func sortMoviesByRuntime(movies models.Movies, ascending bool) {
	if ascending {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Runtime < movies[j].Runtime
		})
	} else {
		sort.Slice(movies, func(i, j int) bool {
			return movies[i].Runtime > movies[j].Runtime
		})
	}
}
