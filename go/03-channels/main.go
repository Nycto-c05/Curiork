package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func base() {
	c := make(chan int, 2) //adding size 2 makes it buffered

	// goroutine sends
	//main goroutine receives
	go func(i int) {
		c <- i
		c <- (i + 1)
	}(4)

	var i = <-c
	fmt.Println(i, <-c)
}

func loop() {
	datachan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			datachan <- i
		}
		close(datachan)
	}()

	for n := range datachan {
		fmt.Printf("n = %d\n", n)
	}
}

func worker() int {
	time.Sleep(time.Second)
	return rand.Intn(10)
}

func main() {

	// base()
	// loop()

	datachan := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				res := worker() // returns some int
				datachan <- res
			}()

		}
		wg.Wait()
		close(datachan)

	}()

	for n := range datachan {
		fmt.Printf("n = %d\n", n)
	}

}
