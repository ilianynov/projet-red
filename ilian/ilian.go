package main

import (
	"fmt"
)

func main() {
	name()
	classe()
}

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxHP      int
	HPactuel   int
	Inventaire []string
}

func name() string {
	var i string

	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scan(&i)
	fmt.Println("Bienvenue dans notre monde", i)
	return ""
}

func classe() string {
	var r int

	fmt.Print("Choisissez votre classe : \n 1. Humain \n 2. Loupgarou \n 3. Hybride \n 4. Nain \n 5. Ange \n 6. Démon \n")
	fmt.Scan(&r)
	if r == 1 {
		fmt.Println("Vous avez choisi la classe Humain")
	} else if r == 2 {
		fmt.Println("Vous avez choisi la classe Loupgarou")
	} else if r == 3 {
		fmt.Println("Vous avez choisi la classe Hybride")
	} else if r == 4 {
		fmt.Println("Vous avez choisi la classe Nain")
	} else if r == 5 {
		fmt.Println("Vous avez choisi la classe Ange")
	} else if r == 6 {
		fmt.Println("Vous avez choisi la classe Démon")
	}

	return ""
}
