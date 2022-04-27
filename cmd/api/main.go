package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/infrastructure/database"
	"github.com/rhodeon/prettylog"
	"os"
	"sync"
)

func main() {
	// load environment variable
	err := godotenv.Load()
	if err != nil {
		prettylog.FatalError(err)
	}

	// setup server configuration
	config := common.Config{}
	config.Parse()
	config.Version = version
	err = config.Validate()
	if err != nil {
		prettylog.FatalError(err)
	}

	// display version and build time before exiting if the flag is set
	if config.DisplayVersion {
		displayVersion()
		os.Exit(0)
	}

	// open database connection
	db, err := openDb(config)
	if err != nil {
		prettylog.FatalError(err)
	}
	defer db.Close()
	prettylog.InfoF("Database connection pool established")

	setMetrics(config, db)

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
