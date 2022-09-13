package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Modifying time⌛")

	prsentTime := time.Now()
	fmt.Println(prsentTime)

	//formatting time

	modifiedTime := prsentTime.Format("1-2-2006 sunday")
	fmt.Println(modifiedTime)

	createdDate := time.Date(2021, time.October, 28, 23, 23, 23, 23, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
