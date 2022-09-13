package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	//no inheritance in go

	tuffy := Animal{name: "tuffy", Age: 9, Weight: 24}
	fmt.Printf("Tuffy complete:%+v \n", tuffy)
	fmt.Printf("Name: %v,Age: %v,Weight: %v", tuffy.name, tuffy.Age, tuffy.Weight)

}

type Animal struct {
	name   string
	Age    int
	Weight int
}
