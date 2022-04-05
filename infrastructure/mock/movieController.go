package mock

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/request"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"strings"
	"time"
)

var movies = []models.Movie{
	{
		Id:      1,
		Title:   "Bullet Train",
		Year:    2022,
		Runtime: 108,
		Genres:  []string{"Action", "Comedy"},
		Version: 1,
		Created: time.Now(),
	},
	{
		Id:      2,
		Title:   "Hamilton",
		Year:    2020,
		Runtime: 140,
		Genres:  []string{"Musical", "Drama"},
		Version: 1,
		Created: time.Now(),
	},
}

type MovieController struct{}

func (m MovieController) Create(movie *models.Movie) error {
	movie.Id = 3
	movie.Version = 1
	movie.Created = time.Now()
	return nil
}

func (m MovieController) Get(id int) (models.Movie, error) {
	for _, movie := range movies {
		if movie.Id == id {
			return movie, nil
		}
	}

	return models.Movie{}, repository.ErrRecordNotFound
}

func (m MovieController) List(title string, genres []string, filters request.Filters) (models.Movies, response.Metadata, error) {
	movieList := models.Movies{}

	// add movies which match both the title and the genres
	for _, movie := range movies {
		if strings.Contains(movie.Title, title) && caseInsensitiveSubslice(genres, movie.Genres) {
			movieList = append(movieList, movie)
		}
	}

	// determine ending index based on page limit
	stop := filters.Offset() + filters.Limit
	if stop > len(movieList) {
		stop = len(movieList)
	}

	// get the total number of found records before reassigning
	// the movie list based on the page number and limit
	totalRecords := len(movieList)
	metadata := response.CalculateMetadata(filters.Page, filters.Limit, totalRecords)
	movieList = movieList[filters.Offset():stop]

	// sort based on the filter
	switch filters.Sort {
	case "id":
		sortMoviesById(movieList, true)
	case "-id":
		sortMoviesById(movieList, false)

	case "title":
		sortMoviesByTitle(movieList, true)
	case "-title":
		sortMoviesByTitle(movieList, false)

	case "year":
		sortMoviesByYear(movieList, true)
	case "-year":
		sortMoviesByYear(movieList, false)

	case "runtime":
		sortMoviesByRuntime(movieList, true)
	case "-runtime":
		sortMoviesByRuntime(movieList, false)
	}

	return movieList, metadata, nil
}

func (m MovieController) Update(movie *models.Movie) error {
	for _, mov := range movies {
		if mov.Id == movie.Id {
			movie.Version = mov.Version + 1
			return nil
		}
	}
	return repository.ErrRecordNotFound
}

func (m MovieController) Delete(id int) error {
	for _, movie := range movies {
		if movie.Id == id {
			// delete nothing as mock data is not persistent
			return nil
		}
	}
	return repository.ErrRecordNotFound
}
