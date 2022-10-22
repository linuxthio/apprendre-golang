package main

import "fmt"

type perso struct {
	prenom string
	nom    string
	age    int
}

func (p perso) aff() {
	pn := "Prenom"
	nn := "Nom"
	ag := "Age"
	fmt.Printf("%9s : %9s \n%9s : %9s\n%9s : %9d\n", pn, p.prenom, nn, p.nom, ag, p.age)

}

func aff(p perso) {
	pn := "Prenom"
	nn := "Nom"
	ag := "Age"
	fmt.Printf("%9s : %9s \n%9s : %9s\n%9s : %9d\n", pn, p.prenom, nn, p.nom, ag, p.age)

}

func main() {
	moi := perso{"djib", "thio", 43}
	aff(moi)
}
