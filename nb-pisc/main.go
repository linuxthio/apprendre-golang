package main

import "fmt"

func nb(a int, b int, c int) {

	n := 140
	aa := n - a
	bb := n - b
	cc := n - c
	t := aa + bb + cc
	r := 400 - t
	fmt.Printf("\npisc 1: %d \npisc 2: %d \npisc 3 : %d \ntotal : %d \nreste : %d\n", aa, bb, cc, t, r)

}

func main() {
	nb(0, 67, 87)

	var a int
	a = 0
	fmt.Printf("valeur de a :")
	i, e := fmt.Scanf("%d", &a)
	if i > 0 {
		fmt.Println("ok ")
	}
	if e == nil {
		fmt.Printf("%d \n", a)
	}

}
