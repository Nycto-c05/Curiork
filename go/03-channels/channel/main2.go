package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var print = fmt.Println

func one() {
	// wont work, unbuff need a receiving routine
	c := make(chan int)
	c <- 1
	i := <-c
	fmt.Println(i)
}

func two() {
	var wg sync.WaitGroup
	//here, routine sends, and main rotine receives
	// when pushin from main routine and recieve in another routine
	// the recieving routine must be ready, so c<-2 is below
	//else deadlock

	//had to add wg, cuz exits before printing
	c := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		val := <-c // waits and listens for val
		fmt.Println(val)
	}()

	time.Sleep(2 * time.Second)
	c <- 2
	wg.Wait()
}

func three() {
	//main routine recieves n go routine sends
	c := make(chan string)

	// wg.Add(1)
	go func(s string) {
		// defer wg.Done()
		//This below will block until main goroutine is ready to receive.
		//if all other routines are asleep/blocked, then deadlock
		c <- s
	}("ligma")
	// wg.Wait()//blocks main

	time.Sleep(2 * time.Second) //even tho main is blocked, there is a possible path, hence no deadlock
	fmt.Println(<-c)

}

func four() {
	c := make(chan int)
	go func() {
		defer close(c) // deadlock w/o this line
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	// val, ok := <-c
	// if !ok {
	// 	print("error")
	// } else {
	// 	print(val, "0th val")
	// }
	for i := range c {
		print(i) //keeps waiting here even when obove routine has completed
		//no one to send, so deadlock
		//hence need the sending routine to close out and notify anyone using channel that no mo vals coming to channel from that routine

	}
}
func main() {
	// one()
	// two()
	// three()
	// four()
}
