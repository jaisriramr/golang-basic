package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("test.txt", []byte("Heello, Go!"), 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	data, _ := os.ReadFile("test.txt")
	fmt.Println(string(data))
}
