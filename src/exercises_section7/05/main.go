package main

import "fmt"

func foo(list ...int) {
	for _, v := range list {
		fmt.Println(v)
	}
}

func main() {
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
	foo()
}
