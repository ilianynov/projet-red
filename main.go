package main

import (
	"fmt"
	karl "projet_red/Karl"
)

var characterChosen bool

func main() {
	level := 1
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
			combatTraining(level)
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
	if characterChosen {
		fmt.Println("Un personnage a déjà été choisi. Vous ne pouvez pas en changer.")
		return
	}

	fmt.Println("=====================================")
	fmt.Println("| Entrez le nom de votre personnage |")
	fmt.Println("=====================================")
	var name string
	fmt.Scan(&name)
	fmt.Printf("| Bienvenue dans notre monde, %s ! |\n", name)
	fmt.Println("=====================================")
	fmt.Println("  _______________________________________________________________")
	fmt.Println(" /                                                               \\")
	fmt.Println("|                    //===================\\                      |")
	fmt.Println("|      CHOISISSEZ VOTRE CLASSE                                   |")
	fmt.Println("|        1. Humain        4. Nain                                |")
	fmt.Println("|        2. Loup-garou    5. Ange                                |")
	fmt.Println("|        3. Hybride       6. Démon                               |")
	fmt.Println("\\_______________________________________________________________/")

	var classChoice int
	for {
		fmt.Print("Choisissez votre classe (1-6) : ")
		fmt.Scan(&classChoice)
		switch classChoice {
		case 1:
			fmt.Println("Vous avez choisi : Humain")
			characterChosen = true
			return
		case 2:
			fmt.Println("Vous avez choisi : Loup-garou")
			characterChosen = true
			return
		case 3:
			fmt.Println("Vous avez choisi : Hybride")
			characterChosen = true
			return
		case 4:
			fmt.Println("Vous avez choisi : Nain")
			characterChosen = true
			return
		case 5:
			fmt.Println("Vous avez choisi : Ange")
			characterChosen = true
			return
		case 6:
			fmt.Println("Vous avez choisi : Démon")
			characterChosen = true
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func displayCharacterInfo() {
	fmt.Println("Affichage des infos du personnage...")
}

func accessInventory() {
	fmt.Println("Accès à l'inventaire...")
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

func calculateGoldReward(monster karl.Monster) int {
	switch monster.Nom {
	case "Gobelin d'entrainement":
		return 10
	case "Ogre":
		return 20
	case "Loup-Garou":
		return 30
	case "Dragon":
		return 50
	default:
		return 0
	}
}

func getMonsterByLevel(level int) karl.Monster {
	switch {
	case level < 5:
		return InitGoblin()
	case level < 10:
		return InitOgre()
	case level < 15:
		return InitLoupGarou()
	default:
		return InitDragon()
	}
}

func combatTraining(level int) {
	monster := getMonsterByLevel(level)
	characterName := "VotrePersonnage"
	characterHP := 100
	gold := 0

	fmt.Printf("\033[36m=====================================\033[0m\n")
	fmt.Printf("\033[33m| Combat entre %s et %s !\033[0m\n", characterName, monster.Nom)
	fmt.Printf("\033[36m=====================================\033[0m\n")

	for {
		fmt.Printf("\033[31m%s : %d/%d PV\033[0m\n", monster.Nom, monster.HPactuel, monster.MaxHP)
		fmt.Printf("\033[32m%s : %d PV\033[0m\n", characterName, characterHP)
		fmt.Println("| 1. Attaquer")
		fmt.Println("| 2. Inventaire")
		fmt.Println("| 0. Quitter le combat")
		fmt.Printf("\033[36m=====================================\033[0m\n")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:

			damage := 15
			monster.HPactuel -= damage
			fmt.Printf("\033[32m%s attaque %s et inflige %d dégâts !\033[0m\n", characterName, monster.Nom, damage)
			if monster.HPactuel <= 0 {
				fmt.Printf("\033[33m%s a été vaincu !\033[0m\n", monster.Nom)
				goldReward := calculateGoldReward(monster)
				gold += goldReward
				fmt.Printf("\033[33mVous avez gagné %d pièces d'or !\033[0m\n", goldReward)
				fmt.Printf("\033[33mOr total : %d\033[0m\n", gold)
				return
			}

			fmt.Printf("\033[31m%s contre-attaque et inflige %d dégâts !\033[0m\n", monster.Nom, monster.Attaque)
			characterHP -= monster.Attaque
			if characterHP <= 0 {
				fmt.Printf("\033[31m%s a été vaincu par %s !\033[0m\n", characterName, monster.Nom)
				return
			}
		case 2:
			fmt.Println("Ouverture de l'inventaire...")
			accessInventory()
		case 0:
			fmt.Println("Vous quittez le combat.")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func spellBooksMenu() {
	fmt.Println("Menu des livres de sorts...")
}

func specialMenu() {
	fmt.Println("Menu spécial...")
}

func InitGoblin() karl.Monster {
	return karl.InitGoblin()
}

func InitOgre() karl.Monster {
	return karl.InitOgre()
}

func InitLoupGarou() karl.Monster {
	return karl.InitLoupGarou()
}

func InitDragon() karl.Monster {
	return karl.InitDragon()
}
