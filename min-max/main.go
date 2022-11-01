package main

import "fmt"

func min_max(tab []int) (int, int) {
	var lmin int
	var lmax int
	lmin = tab[0]
	lmax = -1

	for i := 0; i < len(tab); i++ {
		if tab[i] <= lmin {
			lmin = tab[i]
		}
		if tab[i] > lmax {
			lmax = tab[i]
		}
	}

	return lmin, lmax

}

func main() {
	fmt.Println("vim-go")
	tb := []int{12, 33, 3, 7, 4, 46}
	m, mx := min_max(tb)
	fmt.Println(tb)
	fmt.Println(m)

	fmt.Println(mx)
}
