package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Stock")

}

func stock(s string) {
	point := fmt.Sprintf("https://query1.finance.yahoo.com/v11/finance/quoteSummary/%v", s)
	res, err := http.Get(point)
	if err != nil {
		log.Fatal(err)
	}
}
