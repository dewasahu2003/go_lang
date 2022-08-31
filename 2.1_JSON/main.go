package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Animals struct {
	Name     string `json:"name"` //change the Name -> name
	Weight   int
	Color    string   `json:"-"`                  // will not show the user
	Category []string `json:"category,omitempty"` //if not nil->category,or->omitempty
}

func main() {
	fmt.Println("-JSON")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	animals := []Animals{
		{Name: "tuffy", Weight: 10, Color: "white", Category: []string{"dog", "pet", "domestic-animal"}},
		{Name: "martin", Weight: 250, Color: "orangered", Category: []string{"tiger", "pet", "wild-domestic-animal"}},
		{Name: "lio", Weight: 220, Color: "yellowred", Category: []string{"lion", "pet", "wild-domestic-animal"}},
		{Name: "dastan", Weight: 0, Color: "--", Category: nil},
	}

	bytedata, err := json.MarshalIndent(animals, "", "\t") // GO object _to_ json

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", bytedata)
	file, _ := os.Create("./json-encoded")
	file.WriteString(string(bytedata))

}

func DecodeJson() {
	bytedata, _ := ioutil.ReadFile("./json-encoded") //taking json
	isValid := json.Valid(bytedata)                  //check-whether json in correct format

	var animals []Animals // making a variable to put the json into

	if isValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(bytedata, &animals) // json _to_ GO object
		fmt.Printf("%#v\n", animals)       // %#v := speical for printing var
	} else {
		fmt.Println("JSON is not-valid")
	}

	fmt.Println("========================🤗======================")

	var mapOfAnimals []map[string]interface{}
	json.Unmarshal(bytedata, &mapOfAnimals)
	fmt.Printf("%#v\n", mapOfAnimals)

	fmt.Println("========================🤗======================")

	for key, val := range mapOfAnimals {
		fmt.Printf("%v:%v,(%T,%T)\n", key, val, key, val)
	}
}
