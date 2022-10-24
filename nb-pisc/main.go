package main

import "fmt"

func nb(a int, b int, c int) {

	n := 140
	aa := n - a
	bb := n - b
	cc := n - c
	t := aa + bb + cc
	r := 400 - t
	fmt.Println("\n\nLe nombre de place restant pour la piscine zone-01:")
	fmt.Printf("\npisc 1: %d \npisc 2: %d \npisc 3 : %d \n\ntotal : %d \n\nreste : %d\n", aa, bb, cc, t, r)

}

func setval(nv string) int {
	var n int

	fmt.Printf("Entrer la valeur de %s : ", nv)
	fmt.Scanf("%d", &n)

	return n

}

func main() {

	n1 := setval("n1")
	n2 := setval("n2")
	n3 := setval("n3")
	nb(n1, n2, n3)

	/*
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
	*/

}
