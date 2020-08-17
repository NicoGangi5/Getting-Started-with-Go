package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main()  {

	var input string
	var err error
	var number, lap int = 0, 0

	sli := make([]int, 3)

	for err == nil {

		fmt.Println("Enter your number (X to exit):")
		fmt.Scan(&input)

		number, err = strconv.Atoi(input)

		if lap < 3 {
			sli[lap] = number
			lap++
		} else {
			sli = append(sli, number)
			sort.Ints(sli)
		}

		fmt.Println("The Slice is: ", sli)

	}
}
