package main

import "fmt"

func main() {
	fmt.Println("ARRAY")

	var fruitList = [3]string{"🍎", "🥭", "🍇"}
	fmt.Println(fruitList)
	fmt.Println(len(fruitList)) //the length will be what you define the array while creating it
	//array is not used that much in go
}
