package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	animalMap := make(map[string]string)
	animalMap["🐈"] = "cat"
	animalMap["🐘"] = "elephant"
	animalMap["🐅"] = "tiger"
	animalMap["🐕"] = "dog"

	fmt.Println("Map: ", animalMap)

	//delete
	delete(animalMap, "🐅")
	fmt.Println("new map: ", animalMap)

	//for loop

	for key, value := range animalMap {
		fmt.Printf("For key: %v, value %v \n", key, value)
	}

}
