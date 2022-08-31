package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" //parsing string to anything
	"strings" //making changes to string
)

func main() {
	fmt.Println("Enter age of 🐶")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Age of your dog: ", input)

	numAge, err := strconv.ParseFloat(strings.TrimSpace(input), 32)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 year to your dog's age", numAge+1)
	}

}
