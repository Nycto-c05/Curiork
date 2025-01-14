package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData []string = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	var t0 time.Time = time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1) //add 1 to counter
		go dbCall(i)

	}
	wg.Wait() // like a counter that if goes to 0, means proceed with ahead code
	fmt.Printf("\n\nExecution Time: %v\n", time.Since(t0))
	fmt.Println("The result from the DB is ", results)

}

func dbCall(i int) {
	defer wg.Done() //reduce counter

	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	save(dbData[i])
	log()

}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nCurrent results are: %v", results)
	m.RUnlock()
}
