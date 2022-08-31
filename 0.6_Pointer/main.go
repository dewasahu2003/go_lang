package main

import "fmt"

func main() {

	fmt.Println("POINTER")

	var number *int // when pointer unassigned it is asigned to nil
	fmt.Println(number)

	myNum := 24
	var ptr = &myNum

	*ptr = myNum + 1
	myNum = myNum + 2
	fmt.Println("The pointer plain", ptr)       //give address
	fmt.Println("The pointer with value", *ptr) // * give value

	fmt.Println("New value 🧮", *ptr)
	fmt.Println("New value 👑", myNum)
	//make change to *ptr (holding pointer) or to main var the thing will be same

}
