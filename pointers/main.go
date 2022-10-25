package main

import "fmt"

func chgno(o int) {
	o = 1200
}
func isReady() bool { return true }
func chg(l *int) {
	*l = 90
}

func main() {
	i, j := 12, 10
	p := &i
	chg(&i)
	fmt.Println(&i, i)
	chg(&j)
	chgno(j)
	fmt.Println(&j, j)
	fmt.Println(isReady())
	fmt.Println(*p)
	fmt.Println("ok")
}
