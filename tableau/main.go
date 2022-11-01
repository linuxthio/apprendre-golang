package main

import "fmt"

func main() {
	var tb [3][2]int

	tb[0][0] = 0
	tb[0][1] = 1

	tb[1][0] = 2
	tb[1][1] = 3

	tb[2][0] = 4
	tb[2][1] = 5

	for k, a := range tb {
		fmt.Printf("[%v] = %d \n", k, a)
	}
}
