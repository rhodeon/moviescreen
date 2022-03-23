package main

import (
	"fmt"
	"github.com/rhodeon/moviescreen/cmd/api/internal"
	"log"
	"net/http"
	"time"

	"github.com/rhodeon/prettylog"
)

func main() {
	app := internal.Application{}
	app.Config.Parse()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      app.Router(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	prettylog.InfoF("Starting %s server on %s", app.Config.Env, srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
