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

	fmt.Println("=====================================")
	fmt.Println("| Entrez le nom de votre personnage |")
	fmt.Println("=====================================")
	fmt.Scan(&i)
	fmt.Println("================================")
	fmt.Println("| Bienvenue dans notre monde ! |", i)
	fmt.Println("================================")
	return ""
}

func classe() string {
	var r int

	fmt.Println("  _______________________________________________________________")
	fmt.Println(" /                                                               \\")
	fmt.Println("|      						                                   |")
	fmt.Println("|                    //===================\\                       |")
	fmt.Println("|                    |      CHOISISSEZ     |                      |")
	fmt.Println("|                    |                     |                      |")
	fmt.Println("|                    |         VOTRE       |                      |")
	fmt.Println("|                    |                     |                      |")
	fmt.Println("|                    |        CLASSE       |                      |")
	fmt.Println("|                    \\===================//                       |")
	fmt.Println("|                                                                 |")
	fmt.Println("|                                                                 |")
	fmt.Println("|_________________________________________________________________|")
	fmt.Println(" \\                                                               /")
	fmt.Println("  )                                                             ( ")
	fmt.Println(" /                                                               \\")
	fmt.Println("|			1. Humain        4. Nain                     |")
	fmt.Println("|   		2. Loup-garou    5. Ange                     |")
	fmt.Println("|   		3. Hybride       6. Démon                    |")
	fmt.Println(" \\_______________________________________________________________/")

	fmt.Scan(&r)

	if r == 1 {
		fmt.Println("Vous avez choisi la classe Humain")
	} else if r == 2 {
		fmt.Println("Vous avez choisi la classe Loup-garou")
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
<<<<<<< HEAD
=======

func DisplayInfo() {
	fmt.Println("Nom :", i)
	fmt.Println("Classe :", r)
	fmt.Println("Niveau :", Niveau)
	fmt.Println("HP Actuel :", HPactuel)
	fmt.Println("Max HP :", MaxHP)
	fmt.Println("Inventaire :", Inventaire)
}
>>>>>>> d6d8fe7d9f8404d876eea34df41c5a38879aa6bb
