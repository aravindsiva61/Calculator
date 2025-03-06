package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}
func main() {
	a := 5
	b := 3

	fmt.Println("Addition - ", add(a, b))

	fmt.Println("Subtraction - ", sub(a, b))

	fmt.Println("Multiplication - ", mul(a, b))

}
