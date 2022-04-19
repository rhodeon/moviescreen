package main

import (
	_ "github.com/lib/pq"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/infrastructure/database"
	"github.com/rhodeon/prettylog"
	"sync"
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
			Tokens:      database.TokenController{Db: db},
			Movies:      database.MovieController{Db: db},
			Users:       database.UserController{Db: db},
			Permissions: database.PermissionController{Db: db},
		},
	}

	// establish waitgroup to ensure background tasks
	// are completed before shutting down the application
	backgroundWg := &sync.WaitGroup{}

	// start server
	err = serveApp(app, backgroundWg)
	if err != nil {
		prettylog.FatalError(err)
	}
}
