package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files-in-GO")

	text := "change the order"

	file, err := os.Create("./CTO")

	if err != nil {
		panic(err)
	}

	len, err := file.WriteString(text)

	if err != nil {
		panic(err)
	}
	fmt.Println("The length of file is: ", len)

	defer file.Close()
	readFile("./CTO")

}

func readFile(filepath string) {
	dataByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println("The file data in byte[] : ", dataByte)
	fmt.Println("The file data in string : ", string(dataByte))
}
