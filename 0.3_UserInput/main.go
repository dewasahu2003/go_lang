package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	rating := "How much you like the pizza"
	fmt.Println(rating)

	input, _ := reader.ReadString('\n')
	fmt.Println("ThankYou💓 for pizza: ", input)
	go changetheOrder()
	time.Sleep(2 * time.Second)
	fmt.Printf("The type of rating is: %T ", input)

}
func changetheOrder() {
	fmt.Println("🌍")
}
