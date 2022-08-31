package main

import "fmt"

const Logintoken string = "0xddeioiha564a"

func main() {

	//explicit
	var isDone bool = true
	fmt.Print(isDone)

	var name string = "dewa"
	fmt.Println(name)

	var age int = 18
	fmt.Println(age)

	//implicit

	var name2 = "machoman"
	fmt.Println(name2)

	fmt.Println(Logintoken)
	fmt.Printf("The type of logintoken is : %T", Logintoken)

	isHappy := false
	fmt.Println(isHappy)
}
