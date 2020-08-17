package main

import (
	"encoding/json"
	"fmt"
	"os"
	bu "bufio"
	stg "strings"
)



func main()  {
	var nameI, addI string
	person := make(map[string]interface{})
	personJson := make(map[string]interface{})

	reader := bu.NewReader(os.Stdin)

	fmt.Println("Enter a name:")
	nameI, _ = reader.ReadString('\n')
	nameI = stg.TrimSuffix(nameI, "\n")
	person["name"] = nameI

	fmt.Println("Enter a address:")
	addI, _ = reader.ReadString('\n')
	addI = stg.TrimSuffix(addI, "\n")
	person["address"] = addI

	fmt.Println("The name ", nameI, "and the address ", addI)
	fmt.Println("Map before Json: ", person)

	dataJson, err := json.Marshal(person)
	if err == nil{
		stgJson := string(dataJson)
		fmt.Println("Data in Json formart: ", stgJson)
	}

	err = json.Unmarshal(dataJson, &personJson)
	if err == nil {
		fmt.Println("Map after Json: ", personJson)
	}

}
