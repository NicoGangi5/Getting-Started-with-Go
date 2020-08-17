package main

import (
	bu "bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxLength = 20

type name struct {
	fname string
	lname string
}

func main()  {
	inputReader := bu.NewReader(os.Stdin)

	fmt.Print("Enter the name of the file (please include .txt at the end): ")
	flname, err := inputReader.ReadString('\n')
	flname = strings.Trim(flname, "\n")
	if err != nil {
		log.Fatal(err)
	}

	fl, err := os.Open(flname)
	if err != nil {
		log.Fatal(err)
	}

	names := make([]name, 0, 3)
	data := bu.NewScanner(fl)
	for data.Scan() {
		slide := strings.Split(data.Text(), " ")
		var flName name
		flName.fname, flName.lname = slide[0], slide[1]

		if len(flName.fname) > maxLength {
			flName.fname = flName.fname[:maxLength]
		}

		if len(flName.lname) > maxLength {
			flName.lname = flName.lname[:maxLength]
		}

		names = append(names, flName)
	}

	fmt.Println("The struct:")
	for _, v := range names {
		fmt.Println("First name:", v.fname, " ------------ Last name:", v.lname)
	}
}