package main

import "fmt"

func main() {
	name := "Assegaf"

	if name == "Adbul" {
		fmt.Println("Hello Abdul")
	} else if name == "Gadir" {
		fmt.Println("Hello Gadir")
	} else if name == "Assegaf" {
		fmt.Println("Hello Assegaf")
	} else {
		fmt.Println("Ini Siapa Ya?")
	}

	if leng := len(name); leng > 5 {
		fmt.Println("Nama Kepanjangan")
	} else {
		fmt.Println("Nama Pass")
	}
}
