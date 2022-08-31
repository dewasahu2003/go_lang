package main

import "fmt"

func main() {
	fmt.Println("Functions")
	result := add(1, 9)
	fmt.Printf("The addition is: %+v\n", result)

	proresult, comment, score := proadder(1, 2, 3, 2, 2)
	fmt.Printf("Proadder adds to: %v with comment: %v and a score😀 of: %v", proresult, comment, score)

}

func add(x int, y int) int {
	return x + y
}

func proadder(x ...int) (int, string, int) { //variadic func
	total := 0
	for val := range x {
		total += val
	}
	return total, "wow!", 10 //tuple sytx like py
}
