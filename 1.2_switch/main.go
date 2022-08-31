package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch")

	rand.Seed(time.Now().UnixMilli())
	num := rand.Intn(7)

	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough // if want to continue after conditionis satisfied
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("none amoung them")
	}

}
