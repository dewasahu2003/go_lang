package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("web-request")

	res, err := http.Get(url) //getting the url

	checkerr(err)

	fmt.Println("The type of response is: ", res)
	defer res.Body.Close()                     //IMPORTANT TO CLOSE
	data_byte, err := ioutil.ReadAll(res.Body) //reading from the body of url
	checkerr(err)

	content := string(data_byte)

	fmt.Println("The data in byte is: ", content)

}
func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
