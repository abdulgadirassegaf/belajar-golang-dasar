package main

import "fmt"

func main() {
	type NoKTP string

	var ktpSaya NoKTP = "11111111"
	var contoh = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpSaya)
	fmt.Println(contohKtp)
}
