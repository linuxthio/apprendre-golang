package main

import "fmt"

type Grille struct {
	nom    string
	c      int
	l      int
	pts []Point
	grille [][]int
}

type Point struct{
	c string
	x int
	y int
}

func (g Grille) aff(){
	pt:=[]string{"+","o","-"}
	for l:= range g.grille{
		fmt.Print("\t\t\t")
		for  c :=range g.grille[l]{
			switch g.grille[l][c]{
			case 0:
				fmt.Print(pt[0])
			case 1:
				fmt.Print(pt[1])
			default:
				fmt.Print(pt[2])

			}
		}
		fmt.Print("\n")
	}
}

func (g *Grille) addItem(c int, l int, a int) {
	g.grille[l][c] = a
}
func move(x,y int){


}
func main() {
	// tb := [][]int{
	// 	{0, 0},
	// 	{0, 0},
	// 	{0, 0},
	// }
	gg := Grille{
		nom:    "g",
		c:      5,
		l:      5,
		pts: []Point{{".",0,0}},
		grille: [][]int{{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0},{0,0,0,0,0}},
	}
	// tb[0][0] = 0
	// tb[0][1] = 1

	// tb[1][0] = 2
	// tb[1][1] = 3

	// tb[2][0] = 4
	// tb[2][1] = 5

	r:=0
	for i:=range gg.grille{

		gg.addItem(0, r, i)
		gg.aff()
	}
	
	// fmt.Println(tb)
	// fmt.Println(gg)
}
