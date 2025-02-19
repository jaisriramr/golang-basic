package main

import "fmt"

func main() {
	user := map[string]string{
		"name": "Jaisriram",
		"role": "Developer",
	}
	fmt.Println(user["role"])
}
