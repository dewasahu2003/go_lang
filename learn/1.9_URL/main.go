package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://loc.dev:3000/learn?coursename=reactjs&payement=0xDsdfjodc5644"

// & is comma for url world

func main() {

	fmt.Println("URL:= for getting information from url and constructing them")
	result, _ := url.Parse(myurl)

	fmt.Println("The URL is: ", result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparameter := result.Query()
	fmt.Printf("qparameter is of %T \n", qparameter) // url.Values -> map
	fmt.Println(qparameter["coursename"])

	for key, val := range qparameter {
		fmt.Printf("%v : %v\n", key, val)
	}

	create_Url := url.URL{
		Scheme:  "https",
		Host:    "design.io",
		Path:    "/changetheorderdesign",
		RawPath: "user=dewa",
	}
	proper_url_format := create_Url.String()
	fmt.Println(proper_url_format)
}
