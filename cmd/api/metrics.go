package main

import (
	"database/sql"
	"expvar"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"runtime"
	"time"
)

func setMetrics(config common.Config, db *sql.DB) {
	expvar.NewString(common.MetricVersion).Set(config.Version)

	expvar.Publish(common.MetricTimestamp, expvar.Func(func() any {
		return time.Now().Unix()
	}))

	expvar.Publish(common.MetricGoroutines, expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	expvar.Publish(common.MetricDatabase, expvar.Func(func() any {
		return db.Stats()
	}))
}
