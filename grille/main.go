package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"time"
)

type Grille struct {
	nom    string
	c      int
	l      int
	pts    []Pointd
	grille [][]int
}

type Pointd struct {
	c int
	x int
	y int
}

func (g Grille) aff() {
	pt := []string{".", "o", "-"}
	a := 0
	for l := range g.grille {
		fmt.Printf("%d\t", a)
		for c := range g.grille[l] {
			switch g.grille[l][c] {
			case 0:
				fmt.Print(pt[0])
			case 1:
				fmt.Print(pt[1])
			default:
				fmt.Print(pt[2])
			}
		}
		a = a + 1
		fmt.Print("\n")
	}
}

// erase all points
func (g *Grille) erase() {
	for c := range g.grille {
		for i := range g.grille[c] {
			g.grille[c][i] = 0
		}
	}
}
// add a point
func (g *Grille) addItem(c int, l int, a int) {
	g.grille[l][c] = a
}

func (g *Grille) remplir() {
	for i := 0; i < len(g.pts); i++ {
		l := g.pts[i].x
		c := g.pts[i].y
		g.grille[c][l] = g.pts[i].c
	}
}

func (g *Grille) move(i, a, b int) {
	// get the value of point in pts array
	// change the value of point in pts
	// erase the grille
	// set all points in the grille
	// show the new grille
	g.pts[i].x = a
	g.pts[i].y = b
}

func main() {
	gg := Grille{
		nom:    "g",
		c:      9,
		l:      9,
		pts:    []Pointd{{0, 0, 0}},
		grille: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 0, 0}, 
			{0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}
	
	var o Pointd
	o.c = 1
	o.x = 5
	o.y = 5


	gg.pts = append(gg.pts, o)


	var p Pointd
	p.c = 1
	p.x = 8
	p.y = 4


	gg.pts = append(gg.pts, p)
	x := 0
	teta:=0.0
	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		gg.erase()
		gg.move(1, x, 1)
		r:=4.0
		teta=teta+0.5
		a:=4+r*math.Cos(teta)
		b:=4+r*math.Sin(teta)
		
		gg.move(2, int(a), int(b))
		
		gg.remplir()
		gg.aff()
		x = x + 1
		if x > gg.c-1 {
			x = 0
		}
		time.Sleep(1000000000)
	}
}
