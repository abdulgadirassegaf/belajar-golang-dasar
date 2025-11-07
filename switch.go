package main

import "fmt"

func main() {
	name := "Gadir"

	switch name {
	case "Abdul":
		fmt.Println("Hello Abdul")
	case "Gadir":
		fmt.Println("Hello Gadir")
	default:
		fmt.Println("Ini Siapa Ya?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama kepanjang")
	case false:
		fmt.Println("Nama Pass")
	}

	name = "Abdullah"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah Pass")
	}

}
