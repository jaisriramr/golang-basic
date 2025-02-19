package main

import "fmt"

func main() {

	day := "Friday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("Weekend is near!")
	default:
		fmt.Println("It's a normal day.")
	}

}
