package main

import "fmt"

func main() {
	type NoKTP string

	var ktpYusril NoKTP = "1111111111111"

	var contoh string = "2222222222222"
	var ktpContoh NoKTP = NoKTP(contoh)

	fmt.Println(ktpYusril)
	fmt.Println(ktpContoh)
}