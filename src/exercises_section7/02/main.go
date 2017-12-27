package main

import "fmt"

func main() {

	half := func (one int) (int,bool) {
		return one / 2, one % 2 == 0
	}

	h, even := half(6)
	fmt.Println(h, even)

}
