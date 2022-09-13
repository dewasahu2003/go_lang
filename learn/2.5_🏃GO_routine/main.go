package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// GO_ROUTINE ☄️🏃
//1. run task parallely
//2. flexible stack:= 2 Kb
//3. diffrent than async/await as async/await execute concurrently but GO_routine work parallely

var wg sync.WaitGroup //pointer
var mutx sync.Mutex   //pointer
var signals = []string{"test"}

func main() {
	fmt.Println("GOroutine ☄️🏃")

	// go greeter("A") //running this parallely
	// greeter("Z")

	websites := []string{
		"https://github.com", "https://google.com",
		"https://apple.com", "https://microsoft.com",
	}

	for _, web := range websites {
		go get_statusCode(web) //go routine starts (they will run aggresively in background according to themselves)
		//will consume the memory|until finished
		fmt.Println("🚀🚀🚀")
		wg.Add(1) // we added 4 GO_routine
		fmt.Println("⚡⚡⚡")

	}

	fmt.Println("☄️☄️☄️")
	wg.Wait() //"wait for myguys they are just coming"
	// wait is necessary to hold the goroutine
	// that are in queue when we lock the memory for a particular routine
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

/*
creating request to multiple website using GO_ROUTINE
*/

func get_statusCode(endpoint string) {
	defer wg.Done() //as we completed task we ,,removed the GO_routine
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("something went wrong")
	}
	mutx.Lock()
	signals = append(signals, endpoint)
	mutx.Unlock()
	//if we dont lock the memory then the operation would have been a hactik
	// as all GO_routine we declare will be fighting for the memory
	fmt.Println("🔥🔥🔥")
	fmt.Printf("status code %v: %d\n", endpoint, res.StatusCode)
	fmt.Println(time.Now())
}
