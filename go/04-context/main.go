package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var print = fmt.Println

func one() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	//api that takes some time
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("API data returned")
	case <-ctx.Done():
		fmt.Println("Too long, killed", ctx.Err())
	}
}

func two() {
	//withValue to pass metadata (not huge structs)
	type userId string
	const user userId = "userID"

	ctx := context.WithValue(context.Background(), user, "123") // should not use string as key (not encouraged), custom object only

	if user, ok := ctx.Value(user).(string); ok { //Value returns any, coerce it to .string() and  ;ok means if ok, just compact
		print(user)
	} else {
		print("This is a protected route")
	}

}

func three(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		print("API Resp")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "API response")

	case <-ctx.Done():
		print("Context expired")
		http.Error(w, "Context expired lol", http.StatusGatewayTimeout)
	}
}

func main() {
	// one()
	// two()

	func() {
		http.HandleFunc("GET /", three)
		http.ListenAndServe(":8080", nil)
	}()
}
