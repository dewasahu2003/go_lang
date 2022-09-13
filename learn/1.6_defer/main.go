package main

import "fmt"

func main() {
	fmt.Println("defer")

	defer fmt.Println("one") //this will run last
	fmt.Println("two")       //this will go first

	fmt.Println("========================================")

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer fmt.Println("four")
	fmt.Println("five") // 54321
	make_defer()        //5,43210,four,three,two,one,one

}

func make_defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
