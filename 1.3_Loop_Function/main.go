package main

import "fmt"

func main() {
	fmt.Println("Loop-functioning")

	days := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	//1.
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("============================================")

	//2.
	for d := range days {
		fmt.Println(days[d])
	}
	fmt.Println("============================================")

	//3
	for index, day := range days {
		fmt.Printf("%+v is at index %v\n", day, index)
	}

	//4 kinda....while-loop

	randomVal := 10

	for randomVal < 20 {
		if randomVal == 14 {
			goto home //goto statement
		}

		if randomVal == 15 {
			//break
			randomVal++
			continue
		}

		randomVal++
		fmt.Println(randomVal)

	}
home: //label
	fmt.Println("jumped to home")
}
