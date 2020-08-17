package main

import (
	"fmt"
	bu "bufio"
	"os"
	stg "strings"
)

func main(){
	var input string

	reader := bu.NewReader(os.Stdin)
	fmt.Println("Enter your String: ")

	input, _ = reader.ReadString('\n')
	word := stg.TrimSuffix(input, "\n")
	word = stg.ReplaceAll(word, " ", "")
	word = stg.ToLower(word)

	fmt.Print("The string is: ", input)

	if stg.HasPrefix(word, "i") {
		if stg.Contains(word, "a") {
			if stg.HasSuffix(word, "n") {
				fmt.Println("Found! ")
			} else {
				fmt.Println("Not Found!")
			}
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Not Found!")
	}

}