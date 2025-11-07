package main

import "fmt"

func getFullName() (string, string) {
	return "Abdul Gadir", "Assegaf"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
