package main

import (
	_ "github.com/lib/pq"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/infrastructure/database"
	"github.com/rhodeon/prettylog"
)

func main() {
	// setup server configuration
	config := common.Config{}
	config.Parse()
	err := config.Validate()
	if err != nil {
		prettylog.FatalError(err)
	}

	// open database connection
	db, err := openDb(config)
	if err != nil {
		prettylog.FatalError(err)
	}
	defer db.Close()
	prettylog.InfoF("Database connection pool established")

	app := internal.Application{
		Config: config,
		Repositories: repository.Repositories{
			Movies: database.MovieController{Db: db},
		},
	}

	err = serveApp(app)
	if err != nil {
		prettylog.FatalError(err)
	}
}
