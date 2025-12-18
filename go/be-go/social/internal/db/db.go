package db

import (
	"context"
	"database/sql"
	"time"
)

// not passign the struct here cuz internal package musnt know abt main stuff
func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	//set db configs
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//ping the db
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
