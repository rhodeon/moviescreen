package main

import (
	"database/sql"
	"expvar"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"runtime"
	"time"
)

func setMetrics(config common.Config, db *sql.DB) {
	expvar.NewString("version").Set(config.Version)

	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
}
