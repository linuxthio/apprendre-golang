package main

import "fmt"

type Humain struct {
	prenom  string
	nom     string
	date    string
	salaire int
}

func (h *Humain) set(p, n, d string, s int) {
	h.prenom = p
	h.nom = n
	h.date = d
	h.salaire = s
}

func (h Humain) get() {
	fmt.Println("-------------------------------------------")
	fmt.Println("Prenom : ", h.prenom)

	fmt.Println("Nom : ", h.nom)
	fmt.Println("Date : ", h.date)
	fmt.Println("Salaire : ", h.salaire)
}

func main() {
	var hum Humain
	hum.set(
		"djibril",
		"thiong",
		"16/01/78",
		450000,
	)
	//hum.get()

	fem := Humain{prenom: "penda", nom: "mbengue", date: "12/2/87", salaire: 120000}
	//fem.get()

	humains := []Humain{}

	humains = append(humains, hum)
	humains = append(humains, fem)

	for i := 0; i < len(humains); i++ {

		humains[i].get()
	}

}
