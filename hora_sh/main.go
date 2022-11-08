package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
	"strconv"
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
				ss:=fmt.Sprintf("%v",g.grille[l][c])
				fmt.Print(ss)
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

func gethora() (int, int, int) {
	// h:=time.Hour
	// m:=time.Minute
	// s:=time.Second

	h, m, s := time.Time.Clock(time.Now())

	return h, m, s
}

func main() {
	gg := Grille{
		nom: "g",
		c:   20,
		l:   10,
		pts: []Pointd{{0, 0, 0}},
		grille: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	var o Pointd
	o.c = 0
	o.x = 5
	o.y = 5

	gg.pts = append(gg.pts, o)

	var p Pointd
	p.c = 1
	p.x = 8
	p.y = 4

	gg.pts = append(gg.pts, p)
	x := 0
	teta := 0.0
	for {

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		h, m, s := gethora()
		// display hour minute and second
		

		// pm0 := ""
		// pm1 := ""

		// ps0 := ""
		// ps1 := ""
		// ph0 := ""
		ph1 := ""
		hh := fmt.Sprintf("%v", h)
		sh := strings.Split(hh, "")
		ph0 := ""
		if h > 9 {
			ph0=sh[0]
			ph1=fmt.Sprintf("%v",sh[1])
		}else{
			ph0="0"
			ph1=sh[0]
		}

		ps1 := ""
		ss := fmt.Sprintf("%v", s)
		sss := strings.Split(ss, "")
		ps0 := ""
		if s > 9 {
			ps0 = sss[0]
			ps1=fmt.Sprintf("%v",sss[1])
		}else{
			ps1=sss[0]
			ps0="0"
		}

		pm1 := ""
		mm := fmt.Sprintf("%v", m)
		mmm := strings.Split(mm, "")
		pm0 := ""
		if m > 9 {
			pm0 = mmm[0]
			pm1=fmt.Sprintf("%v",mmm[1])
		}else{
			pm1=mmm[0]
			pm0="0"
		}


		var pth0 Pointd
		var pth1 Pointd

		var ptm0 Pointd
		var ptm1 Pointd
		
		var pts0 Pointd
		var pts1 Pointd

		xhora,yhora:= 3,9

		inth0,err:=strconv.ParseInt(ph0,0,8)

		if err!=nil{
			pth0.c=int(inth0)
			pth0.x=xhora
			pth0.y=yhora
			gg.pts = append(gg.pts, pth0)
		}

		inth1,err:=strconv.ParseInt(ph1,0,8)

		if err!=nil{
			pth1.c=int(inth1)
			pth1.x=xhora+1
			pth1.y=yhora
			gg.pts = append(gg.pts, pth1)
		}
		intm0,err:=strconv.ParseInt(pm0,0,8)

		if err!=nil{
			ptm0.c=int(intm0)
			ptm0.x=xhora+2
			ptm0.y=yhora
			gg.pts = append(gg.pts, ptm0)
		}

		intm1,err:=strconv.ParseInt(pm1,0,8)

		if err!=nil{
			ptm1.c=int(intm1)
			ptm1.x=xhora+3
			ptm1.y=yhora
			gg.pts = append(gg.pts, ptm1)
		}

		ints0,err:=strconv.ParseInt(ps0,0,8)

		if err!=nil{
			pts0.c=int(ints0)
			pts0.x=xhora+4
			pts0.y=yhora
			gg.pts = append(gg.pts, pts0)
		}

		ints1,err:=strconv.ParseInt(ps1,0,8)

		if err!=nil{
			pts1.c=int(ints1)
			pts1.x=xhora+5
			pts1.y=yhora
			gg.pts = append(gg.pts, pts1)
		}

		fmt.Println("h = ", ph0," h1", ph1)
		fmt.Println("m = ", pm0," m1", pm1)
		fmt.Println("s = ", ps0," s1", ps1)
		fmt.Println(h, m, s)
		gg.erase()
		gg.move(1, x, 1)
		r := 4.0
		teta = teta + 0.5
		a := 4 + r*math.Cos(teta)
		b := 4 + r*math.Sin(teta)

		gg.move(2, int(a), int(b))

		gg.remplir()
		gg.aff()
		x = x + 1
		if x > gg.c-1 {
			x = 0
		}

		fmt.Println(gg.pts)
		time.Sleep(100000000)
	}
}
