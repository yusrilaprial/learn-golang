package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	var c = 3
	var d = a + b - c * 2
	fmt.Println(d)
	fmt.Println(10 / 2)
	fmt.Println(10 % 2)

	// Augmented Assigment (+= -= *= /= %=)
	var i = 10
	i +=10
	fmt.Println(i)

	i +=5
	fmt.Println(i)

	// Unary Operator
	var j = 1
	j++
	j++
	fmt.Println(j)
	j--
	j--
	fmt.Println(j)
}