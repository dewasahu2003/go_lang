package main

import (
	"fmt"
	"sort"
)

func main() {
	//two ways to allocate memory
	// 1.new:= create|no init|zero memory
	// 2.make:= create|init|non-zero memory
	fmt.Println("Slices 🦊🦊🦊")
	//slices are fully functional array

	fruitList := []string{"🍅", "🥭", "🍇"}
	fmt.Println("The friutList: ", fruitList)

	fruitList = append(fruitList, "🍉")
	fmt.Println("New friutList: ", fruitList)

	//can perform operation like py for array[1:3] and so on

	fmt.Println("fruitList: ", fruitList[1:3])

	//using memory allocation

	animals := make([]string, 3) //this is default memory allocation
	animals[0] = "🐕"
	animals[1] = "🐈"
	animals[2] = "🦒"

	fmt.Println("Animal: ", animals)

	animals = append(animals, "🐅", "🐘") // re-memory allocation
	fmt.Println("New Animals: ", animals)

	//sorting
	sort.Strings(animals)
	fmt.Println("sorted: ", animals)
	isSorted := sort.StringsAreSorted(animals)
	fmt.Println("is List Sorted: ", isSorted)

	//to delete an element from array
	index_to_remove := 2
	animalList := append(animals[:index_to_remove], animals[index_to_remove+1:]...)
	fmt.Println("The new animal List: ", animalList)

}
