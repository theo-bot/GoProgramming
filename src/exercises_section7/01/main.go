package main

import "fmt"

func dual (one int) (int,bool) {
	return one / 2, one % 2 == 0
}

func main() {
	h, even := dual(5)
	fmt.Println(h, even)

}
