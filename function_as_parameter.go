package main

import "fmt"

type Filter func(string) string

func sayHellowithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	} else {
		return name
	}
}

func main() {
	sayHellowithFilter("Gadir", spamFilter)

	filter := spamFilter
	sayHellowithFilter("Anjing", filter)
}
