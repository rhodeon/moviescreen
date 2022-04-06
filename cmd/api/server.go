package main

import (
	"fmt"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/handlers"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"github.com/rhodeon/prettylog"
	"net/http"
	"time"
)

// starts up the server with the app data
func serveApp(app internal.Application) error {
	routeHandlers := common.RouteHandlers{
		Error:  handlers.NewErrorHandler(),
		Misc:   handlers.NewMiscHandler(app.Config),
		Movies: handlers.NewMovieHandler(app.Config, app.Repositories),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      app.Router(routeHandlers),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start server
	prettylog.InfoF("Starting %s server on %s", app.Config.Env, srv.Addr)
	return srv.ListenAndServe()
}
