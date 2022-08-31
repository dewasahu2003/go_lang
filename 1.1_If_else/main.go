package main

import "fmt"

func main() {

	fmt.Println("if-else")

	age := 10
	if age < 10 {
		fmt.Println("is a kid")
	} else if age < 15 || age > 10 {
		fmt.Println("is teenager")
	} else {
		fmt.Println("is adult")
	}

	if num := 10; num < 5 {
		fmt.Println("num is less than 5")
	} else {
		fmt.Println("num is grater than 5")
	}

}
