// +build exclude

/*

 Go Routine Race Condition Example
 run it with go race detector:
   $ go run -race cmain.go

	Example Output:
  --------------------------------------------------------
  WARNING: DATA RACE
	Read at 0x00c0000160d8 by goroutine 8:
		t.RaceCondition.func2()
				[...]racendition.go:19 +0x38

	Previous write at 0x00c0000160d8 by goroutine 7:
		t.RaceCondition.func1()
				[...]racendition.go:14 +0x38

	Goroutine 8 (running) created at:
		t.RaceCondition()
				[...]racendition.go:18 +0x10a
		main.main()
				[...]main.go:8 +0x2f

	Goroutine 7 (finished) created at:
		t.RaceCondition()
				[...]/racendition.go:13 +0xde
		main.main()
				[...]/main.go:8 +0x2f
	==================
	x is:  5
	Found 1 data race(s)
	exit status 66
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(RaceCondition()) // returns 4 or 5 depending by Go Scheduler
}

//RaceCondition is an example of a Race Condition in Go
func RaceCondition() int {
	x := 0 // shared variable, object of the race condition
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() { // create and run first concurrent goroutine
		x = 4
		wg.Done()
	}()

	go func() { // create and run second concurrent goroutine
		x++
		wg.Done()
	}()

	wg.Wait() // wait for both goroutine to finish
	return x
}
