package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	errors2 "github.com/rhodeon/moviescreen/cmd/api/errors"
	"github.com/rhodeon/moviescreen/cmd/api/handlers"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"github.com/rhodeon/prettylog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// serveApp starts up a server with the app data.
func serveApp(app internal.Application, backgroundWaitGroup *sync.WaitGroup) error {
	routeHandlers := common.RouteHandlers{
		Error:  errors2.NewErrorHandler(),
		Misc:   handlers.NewMiscHandler(app.Config),
		Movies: handlers.NewMovieHandler(app.Config, app.Repositories),
		Users:  handlers.NewUserHandler(app.Config, app.Repositories, backgroundWaitGroup),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      app.Router(routeHandlers),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start a background goroutine to intercept and handle shutdown events
	shutdownError := make(chan error)
	go handleShutdown(srv, shutdownError, backgroundWaitGroup)

	// start and listen on server until an error occurs
	prettylog.InfoF("starting %s server on %s", app.Config.Env, srv.Addr)
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		// return unsuccessful server shutdown errors
		return err
	}

	// block flow until the shutdown error channel is updated
	err = <-shutdownError
	if err != nil {
		return err
	}

	prettylog.Info("stopped server")
	return nil
}

// handleShutdown gracefully handles interruption and termination signals,
// giving ongoing request a 20-second leeway before shutting down the server.
// It should be run as a background goroutine.
func handleShutdown(server *http.Server, shutdownErr chan error, backgroundWg *sync.WaitGroup) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// wait until the quit channel is updated with a signal
	s := <-quit

	// 20-second timeout context to delay shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	prettylog.InfoF("shutting down server: %s", s)
	// shutdown the server and update the error channel
	// to resume execution on the main goroutine
	err := server.Shutdown(ctx)
	if err != nil {
		shutdownErr <- err
	}

	// wait for background tasks to complete before shutting down the application
	prettylog.Info("completing background tasks")
	backgroundWg.Wait()
	shutdownErr <- nil
}
