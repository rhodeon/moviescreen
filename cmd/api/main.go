package main

import (
	"fmt"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/handlers"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/rhodeon/prettylog"
)

func main() {
	app := internal.Application{}

	// setup server configuration
	app.Config.Parse()
	err := app.Config.Validate()
	if err != nil {
		prettylog.FatalError(err)
	}

	routeHandlers := common.RouteHandlers{
		Error:  handlers.NewErrorHandler(),
		Misc:   handlers.NewMiscHandler(app.Config),
		Movies: handlers.NewMovieHandler(app.Config),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      app.Router(routeHandlers),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// open database connection
	db, err := openDb(app.Config)
	if err != nil {
		prettylog.FatalError(err)
	}
	defer db.Close()
	prettylog.InfoF("Database connection pool established")

	// start server
	prettylog.InfoF("Starting %s server on %s", app.Config.Env, srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
