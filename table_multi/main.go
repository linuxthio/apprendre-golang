package main

import "fmt"

// this function make a multiplication table of n

func table(n int, lim int) []string {
	var tb []string
	s := ""
	for i := 0; i < lim; i++ {
		r := n * i

		s = fmt.Sprintf("%d x %d = %d ", n, i, r)
		tb = append(tb, s)

	}
	return tb
}

func table_v2(n int) []string {
	return table(n, 11)

}

func main() {

	for _, v := range table(2, 20) {
		fmt.Println(v)
	}
}
