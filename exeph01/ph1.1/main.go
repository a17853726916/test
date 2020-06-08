package main

import "fmt"

func plusTwo() func(int) int {
	return func(y int) int {
		return y + 2
	}
}

func main() {
	f := plusTwo()
	fmt.Println(f(6))
}
