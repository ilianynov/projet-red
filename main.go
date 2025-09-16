package main

import (
	"fmt"
)

const (
	green = "\033[32m"
	cyan  = "\033[36m"
	reset = "\033[0m"
)

func main() {
	for {
		showMainMenu()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			createCharacter()
		case 2:
			displayCharacterInfo()
		case 3:
			accessInventory()
		case 4:
			shopMenu()
		case 5:
			blacksmithMenu()
		case 6:
			combatTraining()
		case 7:
			spellBooksMenu()
		case 8:
			specialMenu()
		case 0:
			fmt.Println("Merci d'avoir joué !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func showMainMenu() {
	fmt.Println("\n================ RPG ULTIME ================")
	fmt.Println("1. Créer un personnage")
	fmt.Println("2. Afficher les infos du personnage")
	fmt.Println("3. Accéder à l'inventaire")
	fmt.Println("4. Boutique (Marchand)")
	fmt.Println("5. Forgeron")
	fmt.Println("6. Entrainement (Combat)")
	fmt.Println("7. Livres de sorts")
	fmt.Println("8. Spécial (Potion/Poison)")
	fmt.Println("0. Quitter")
	fmt.Println("============================================")
	fmt.Print("Votre choix : ")
}

func createCharacter() {
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(cyan + "| Entrez le nom de votre personnage |" + reset)
	fmt.Println(green + "=====================================" + reset)
	var name string
	fmt.Scan(&name)
	fmt.Printf(cyan+"| Bienvenue dans notre monde, %s ! |\n"+reset, name)
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(cyan + "  _______________________________________________________________" + reset)
	fmt.Println(cyan + " /                                                               \\" + reset)
	fmt.Println(cyan + "|                    //===================\\                      |" + reset)
	fmt.Println(cyan + "|      CHOISISSEZ VOTRE CLASSE                                   |" + reset)
	fmt.Println(cyan + "|        1. Humain        4. Nain                                |" + reset)
	fmt.Println(cyan + "|        2. Loup-garou    5. Ange                                |" + reset)
	fmt.Println(cyan + "|        3. Hybride       6. Démon                               |" + reset)
	fmt.Println(cyan + "\\_______________________________________________________________/" + reset)
}

func displayCharacterInfo() {
	fmt.Println("Affichage des infos du personnage...")
	// Call character.DisplayInfo() here
}

func accessInventory() {
	fmt.Println("Accès à l'inventaire...")
	// Call character.AccessInventory() here
}

func shopMenu() {
	fmt.Println("================ BOUTIQUE =================")
	fmt.Println("Or actuel : <montant>")
	fmt.Println("-------------------------------------------")
	fmt.Println(" 1. <item> <prix> or  [Rareté: <niveau>]")
	fmt.Println(" ...")
	fmt.Println("-------------------------------------------")
	fmt.Println(" 0. Quitter")
}

func blacksmithMenu() {
	fmt.Println("=====================================")
	fmt.Println("|        Bienvenue chez le Forgeron !         |")
	fmt.Println("=====================================")
	fmt.Println("| 1. Épée basique (10 Or)           |")
	fmt.Println("| 2. Épée rare (20 Or)              |")
	fmt.Println("| 3. Épée épique (25 Or)            |")
	fmt.Println("| 4. Épée légendaire (40 Or)        |")
	fmt.Println("| 5. Épée démoniaque (60 Or)        |")
	fmt.Println("| 6. Épée angélique (60 Or)         |")
	fmt.Println("| 7. Armure rare (30 Or)            |")
	fmt.Println("| 8. Armure épique (50 Or)          |")
	fmt.Println("| 9. Armure légendaire (70 Or)      |")
	fmt.Println("| 10. Armure démoniaque (90 Or)     |")
	fmt.Println("| 11. Armure angélique (90 Or)      |")
	fmt.Println("| 12. Petite bombe (30 Or)          |")
	fmt.Println("| 13. Grosse bombe (50 Or)          |")
	fmt.Println("| 14. Bombe nucléaire (100 Or)      |")
	fmt.Println("| 15. Quitter                       |")
	fmt.Println("=====================================")
}

func combatTraining() {
	fmt.Println("=====================================")
	fmt.Println("|         Combat contre : <nom>     |")
	fmt.Println("=====================================")
	fmt.Println("| 1. Attaquer")
	fmt.Println("| 2. Inventaire")
	fmt.Println("=====================================")
	fmt.Print("Votre choix : ")
	// Add combat logic here
}

func spellBooksMenu() {
	fmt.Println("Menu des livres de sorts...")
	// Add spell book logic here
}

func specialMenu() {
	fmt.Println("Menu spécial...")
	// Add special menu logic here
}
