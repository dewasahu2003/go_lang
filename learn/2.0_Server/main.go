package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Server 🙂")
	get("")
	fmt.Println("=================🤗==============")
	postJson("post")
	fmt.Println("=================🤗==============")
	postForm("postform")
	fmt.Println("=================🤗==============")

}

func postForm(path string) {
	//post in URLencoded format
	myUrl := fmt.Sprintf("http://localhost:1000/%+v", path)

	//making url encoded data
	data := url.Values{}
	data.Add("name", "dewa")
	data.Add("phno", "+918602996068")
	fmt.Println("url encoded data: ", data)

	//posting
	post_req, err := http.PostForm(myUrl, data)
	Err(err)
	defer post_req.Body.Close()

	//reading the req we posted
	bytedata, _ := ioutil.ReadAll(post_req.Body)

	//convert in readable format
	var byteconverter strings.Builder
	byteconverter.Write(bytedata)

	fmt.Println("The data posted is: ", byteconverter.String())
	fmt.Println("Status code: ", post_req.StatusCode)

}

func postJson(path string) {
	// post data in JSON||

	myUrl := fmt.Sprintf("http://localhost:1000/%+v", path)

	//data to post
	dataToPost := strings.NewReader(`
		{
			"fontStyle":"dotty",
			"fontSize":24
			
		}	
	`)

	req_post, err := http.Post(myUrl, "application/json", dataToPost)
	Err(err)
	defer req_post.Body.Close()

	bytedata, _ := ioutil.ReadAll(req_post.Body)
	data := string(bytedata)

	fmt.Println("The posted data is: ", data)
	fmt.Println("Status code: ", req_post.StatusCode)

}

func get(path string) {
	myUrl := fmt.Sprintf("http://localhost:1000/%+v", path)

	res, err := http.Get(myUrl)
	Err(err)
	defer res.Body.Close()

	bytedata, err := ioutil.ReadAll(res.Body)
	Err(err)
	// data := string(bytedata)    // one way to handle byte response

	var byteconverter strings.Builder
	len, _ := byteconverter.Write(bytedata)
	fmt.Println("data len: ", len)

	data := byteconverter.String()
	fmt.Println(data)
	fmt.Println("Status code: ", res.StatusCode)
}

func Err(err error) {
	if err != nil {
		panic(err)
	}

}
