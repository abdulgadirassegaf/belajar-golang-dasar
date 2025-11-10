package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var gadir Customer
	fmt.Println(gadir)

	gadir.Name = "Abdul Gadir Assegaf"
	gadir.Address = "Indonesia"
	gadir.Age = 28
	fmt.Println(gadir)
	fmt.Println(gadir.Name)
	fmt.Println(gadir.Address)
	fmt.Println(gadir.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	gadir.sayHello("Agus")
	joko.sayHello("Agus")
}
