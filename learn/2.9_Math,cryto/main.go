package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Math🧮,Random🌿,Crypto🪙")

	//rand_math

	// rand.Seed(time.Now().UnixNano())
	// rand_num := rand.Intn(5)
	// fmt.Println(rand_num)

	// rand_crypto

	//more accurate
	num, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(num)

}
