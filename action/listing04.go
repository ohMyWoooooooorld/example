package action

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}
	fmt.Println("Completed:", prefix)
}

func Goroutines4() {
	runtime.GOMAXPROCS(2)

	wg.Add(2)
	go printPrime("Goroutine A")
	go printPrime("Goroutine B")

	fmt.Println("Main goroutine: waiting for goroutines to complete")
	wg.Wait()
	fmt.Println("Main goroutine: all goroutines completed")
}
