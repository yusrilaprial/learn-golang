package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muh"
	names[1] = "Yusril"
	names[2] = "Aprial"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{10, 20, 30, 40, 50, 100}
	fmt.Println(values)

	fmt.Println(len(values))
}