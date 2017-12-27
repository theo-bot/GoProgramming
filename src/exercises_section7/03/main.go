package main

import "fmt"

func max(numbers ...int) int {
	fmt.Printf("%T\n", numbers)
	var largest int
	for _,v := range  numbers {
		if v > largest {
			largest = v
		}
	}

	return largest
}

func main() {
	greatest := max(6,8,2,9)
	fmt.Println(greatest)
}
