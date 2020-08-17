package main

import (
	"fmt"
)

func main(){
	var input float32

	fmt.Println("enter your number to be truncated")

	fmt.Scan(&input)
	fmt.Println("Your number truncated is: ", int(input))
}