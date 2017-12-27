package main

import "fmt"

func multiple3_or_5(number int) bool {
	return (number % 3 == 0 ) || (number % 5 == 0)
}

func main() {

	var total int

	for i:=1; i<1000; i++ {
		if multiple3_or_5(i) {
			total += i
		}
	}
	fmt.Println(total)
}

/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

 */