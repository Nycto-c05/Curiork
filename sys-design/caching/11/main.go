package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

func mountRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

func mountPg(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres")
	if err != nil {
		log.Fatal("cannot connect:", err)
	}
	return conn
}

func redisGet(rdb *redis.Client, ctx context.Context) error {
	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		return err
	}
	_ = val // just consume
	return nil
}

func pgGet(conn *pgx.Conn, ctx context.Context) error {
	var n string
	err := conn.QueryRow(ctx, "SELECT name FROM demo WHERE name='sara';").Scan(&n)
	if err != nil {
		return err
	}
	_ = n
	return nil
}

func benchmark(name string, fn func() error, iterations int) {
	var min, max, total time.Duration
	min = time.Hour // initialize very high

	for i := 0; i < iterations; i++ {
		start := time.Now()
		if err := fn(); err != nil {
			log.Fatal(name, "error:", err)
		}
		elapsed := time.Since(start)
		total += elapsed
		if elapsed < min {
			min = elapsed
		}
		if elapsed > max {
			max = elapsed
		}
	}

	fmt.Printf("%s - iterations: %d, avg: %v, min: %v, max: %v\n",
		name, iterations, total/time.Duration(iterations), min, max)
}

func main() {
	ctx := context.Background()

	// Setup
	rdb := mountRedis()
	defer rdb.Close()

	// Ensure Redis key exists
	if err := rdb.Set(ctx, "foo", "bar", 0).Err(); err != nil {
		log.Fatal(err)
	}

	pg := mountPg(ctx)
	defer pg.Close(ctx)

	// Run benchmark
	iterations := 1000

	benchmark("Redis GET", func() error { return redisGet(rdb, ctx) }, iterations)
	benchmark("Postgres SELECT", func() error { return pgGet(pg, ctx) }, iterations)
}
