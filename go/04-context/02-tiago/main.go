package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var print = fmt.Println

// ----------------------
func doWork(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(2 * time.Second):
		print("Work done")
	case <-ctx.Done():
		print("Cancelled ", ctx.Err())
	}
}

func one() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup //or wait for go routine to return by adding timer at end

	wg.Add(1)
	go doWork(ctx, &wg)
	wg.Wait()
}

// ----------------------

func main() {
	one()
}

/*
so when the context timeouts the done channel becomes ready to read ie is closes the done channel, and if it does not timeout, and i call cancel() then i force the done channel to become available to read by clsing it
*/
