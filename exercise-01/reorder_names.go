package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 || len(args) > 4 {
		fmt.Println("Usage: reorder_names.go <first name> <last name> [<middle name>] <country code>")
		return
	}

	firstName := args[0]
	lastName := args[1]
	middleName := ""
	countryCode := args[len(args)-1]

	if len(args) == 4 {
		middleName = args[2]
	}

	switch countryCode {
	case "US":
		// first name, middle initial, and last name
		fmt.Println(firstName, middleName, lastName)
	case "VN":
		// last name, middle initial, and first name
		fmt.Println(lastName, middleName, firstName)
	// Add more country formats here
	default:
		fmt.Println("Unsupported country code")
	}
}
