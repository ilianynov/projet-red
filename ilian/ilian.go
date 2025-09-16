package main

import (
	"fmt"
)

func main() {
	Name()
	Classe()
}

func Name() string {
	var i string

	fmt.Println("=====================================")
	fmt.Println("| Entrez le nom de votre personnage |")
	fmt.Println("=====================================")
	fmt.Scan(&i)
	fmt.Println("========================================")
	fmt.Println("| Bienvenue dans notre monde,", i, " ! |")
	fmt.Println("========================================")
	return ""
}

var ClasseValue string

func Classe() string {
	var r int

	fmt.Println("  _______________________________________________________________")
	fmt.Println(" /                                                               \\")
	fmt.Println("|      						      		  |")
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
	fmt.Println("|		     1. Humain        4. Nain                     |")
	fmt.Println("|   	             2. Loup-garou    5. Ange                     |")
	fmt.Println("|   		     3. Hybride       6. Démon                    |")
	fmt.Println(" \\_______________________________________________________________/")

	fmt.Scan(&r)

	if r == 1 {
		fmt.Println("Vous avez choisi la classe Humain")
		ClasseValue = "Humain"
	} else if r == 2 {
		fmt.Println("Vous avez choisi la classe Loup-garou")
		ClasseValue = "Loup-garou"
	} else if r == 3 {
		fmt.Println("Vous avez choisi la classe Hybride")
		ClasseValue = "Hybride"
	} else if r == 4 {
		fmt.Println("Vous avez choisi la classe Nain")
		ClasseValue = "Nain"
	} else if r == 5 {
		fmt.Println("Vous avez choisi la classe Ange")
		ClasseValue = "Ange"
	} else if r == 6 {
		fmt.Println("Vous avez choisi la classe Démon")
		ClasseValue = "Démon"
	}

	return ""
}
