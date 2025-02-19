package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	nums = append(nums, 40)
	fmt.Println(nums)
}
