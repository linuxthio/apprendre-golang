package main

import "fmt"

type Grille struct {
	nom    string
	c      int
	l      int
	grille [][]int
}

func (g Grille) aff(){
	for l:= range g.grille{
		for  c :=range g.grille[l]{
			fmt.Print(g.grille[l][c])
		}
		fmt.Print("\n")
	}
}

func (g *Grille) addItem(c int, l int, a int) {
	g.grille[l][c] = a
}

func main() {
	tb := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
	}
	gg := Grille{
		nom:    "g",
		c:      5,
		l:      5,
		grille: [][]int{{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0}},
	}
	tb[0][0] = 0
	tb[0][1] = 1

	tb[1][0] = 2
	tb[1][1] = 3

	tb[2][0] = 4
	tb[2][1] = 5

	gg.addItem(0, 0, 3)
	gg.aff()
	for k, a := range tb {
		fmt.Printf("[%v] = %d \n", k, a)
	}
	fmt.Println(tb)
	fmt.Println(gg)
}
