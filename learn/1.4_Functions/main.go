package main

import "fmt"

func main() {
	fmt.Println("Functions")
	result := add(1, 9)
	fmt.Printf("The addition is: %+v\n", result)

	proresult, comment, score := proadder(1, 2, 3, 2, 2)
	fmt.Printf("Proadder adds to: %v with comment: %v and a score😀 of: %v\n", proresult, comment, score)

	anonymus_function := func(s string) { //like js
		fmt.Printf("%+v I'm anonymous\n", s)
	}
	anonymus_function("hello")
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
