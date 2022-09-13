package main

import (
	"fmt"
	"sync"
)

// RACE_condition := when multiple goroutine run and we have only one memory //fight for memory
// || go run --race main.go || -> for checking race condition -> which can save cloud costing
func main() {
	fmt.Println("Race_Condtion")

	wg := &sync.WaitGroup{}
	mutx := &sync.RWMutex{}

	//RWMutex:= will stop all READ GOroutine when WRITE GOroutine is running

	score := []int{0}

	wg.Add(4) // adding the waitgroup for number of routine we have

	go func(WG *sync.WaitGroup, M *sync.RWMutex) { //if is (inline function)
		mutx.Lock()
		fmt.Println("One-running")
		score = append(score, 1)
		mutx.Unlock()
		defer wg.Done()
	}(wg, mutx)

	go func(WG *sync.WaitGroup, M *sync.RWMutex) {
		mutx.Lock()
		fmt.Println("Two-running")
		score = append(score, 2)
		mutx.Unlock()
		defer wg.Done()
	}(wg, mutx)

	go func(WG *sync.WaitGroup, M *sync.RWMutex) {
		mutx.Lock()
		fmt.Println("Three-running")
		score = append(score, 3)
		mutx.Unlock()
		defer wg.Done()
	}(wg, mutx)

	go func(WG *sync.WaitGroup, M *sync.RWMutex) {
		mutx.RLock()
		fmt.Println("Four-running")
		score = append(score, 4)
		fmt.Printf("%+v score from 4 ⚡\n", score)
		mutx.RUnlock()
		defer wg.Done()
	}(wg, mutx)
	wg.Wait()

	fmt.Println(score)
}
