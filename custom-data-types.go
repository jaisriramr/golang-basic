package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Jaisriram", Age: 25}
	fmt.Println(p.Name, p.Age)
}
