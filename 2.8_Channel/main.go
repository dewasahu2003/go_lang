package main

import (
	"fmt"
	"sync"
)

// CHANNELS 📺 := pipeline(way) through which goroutineS can talk to each other
// channel will only take value if & only if it has listeners attached to it
func main() {
	fmt.Println("Channels 📺")

	//giving length to channel doesnt work with channel
	myChannel := make(chan int, 1)

	wg := &sync.WaitGroup{}

	// myChannel <- 5                       //this creates deadlock condition
	// fmt.Println(<-myChannel)

	wg.Add(2)

	// <-chan := reading/getting value from channel

	go func(CH <-chan int, WG *sync.WaitGroup) {
		value, isChannelOpen := <-myChannel
		fmt.Printf("change open:%+v || channel value:%+v\n", isChannelOpen, value)

		fmt.Printf("channel🚿🤽 value: %+v\n", <-myChannel)
		fmt.Printf("channel🚿🤽 value: %+v\n", <-myChannel)
		defer wg.Done()
	}(myChannel, wg)

	// chan<- := inserting/seeding value in channel

	go func(CH chan<- int, WG *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 10
		myChannel <- 20

		close(myChannel)

		//myChannel <- 6
		//close(myChannel)
		defer wg.Done()
	}(myChannel, wg)

	wg.Wait()
}

/* closing the channel before seeding value with close the channel BUT
it will have a value =0 ,
so to tackle this we do
value,isChannelOpen:= <-myChannel
*/
