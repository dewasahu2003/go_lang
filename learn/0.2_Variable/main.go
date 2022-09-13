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

	//var
	//It is used to declare and initialize the variables
	//inside and outside the functions.
	//need not be init

	//:=
	//It is used to declare and initialize the variables only inside the functions.
	//need to be init
}
