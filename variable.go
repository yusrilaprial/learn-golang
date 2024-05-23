package main

import "fmt"

func main() {
	var name string

	name = "Yusril"
	fmt.Println(name)

	name = "M Yusril Aprial"
	fmt.Println(name)

	// Simplefied Variable
	job := "Programmer"
	fmt.Println(job)

	job = "Full Stack"
	fmt.Println(job)

	// Membuat variable sekaligus
	var (
		firstName = "Muhammad"
		middleName = "Yusril"
		lastName = "Aprial"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}