package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(10, 1)
	fmt.Println("Sum:", result)
}
