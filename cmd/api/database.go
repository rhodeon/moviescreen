package main

import (
	"context"
	"database/sql"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"time"
)

func openDb(config common.Config) (*sql.DB, error) {
	dbCfg := config.Db
	db, err := sql.Open("postgres", dbCfg.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbCfg.MaxOpenConns)
	db.SetMaxIdleConns(dbCfg.MaxIdleConns)

	idleTime, err := time.ParseDuration(dbCfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(idleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
