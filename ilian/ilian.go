package main

import (
	"fmt"
)

func main() {
	name()
	classe()
}

func name() string {
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

func classe() string {
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
		Classe = "Humain"
	} else if r == 2 {
		fmt.Println("Vous avez choisi la classe Loup-garou")
		Classe = "Loup-garou"
	} else if r == 3 {
		fmt.Println("Vous avez choisi la classe Hybride")
		Classe = "Hybride"
	} else if r == 4 {
		fmt.Println("Vous avez choisi la classe Nain")
		Classe = "Nain"
	} else if r == 5 {
		fmt.Println("Vous avez choisi la classe Ange")
		Classe = "Ange"
	} else if r == 6 {
		fmt.Println("Vous avez choisi la classe Démon")
		Classe = "Démon"
	}

	return ""
}
