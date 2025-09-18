package main

import (
	"fmt"
	karl "projet_red/Karl"
	"projet_red/character"
	"projet_red/item"
)

var player character.Character
var characterChosen bool
var xp int = 0
var level int = 1

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
	var classe string
	for {
		fmt.Print("Choisissez votre classe (1-6) : ")
		fmt.Scan(&classChoice)
		switch classChoice {
		case 1:
			classe = "Humain"
			fmt.Println("Vous avez choisi : Humain")
			characterChosen = true
			break
		case 2:
			classe = "Loup-garou"
			fmt.Println("Vous avez choisi : Loup-garou")
			characterChosen = true
			break
		case 3:
			classe = "Hybride"
			fmt.Println("Vous avez choisi : Hybride")
			characterChosen = true
			break
		case 4:
			classe = "Nain"
			fmt.Println("Vous avez choisi : Nain")
			characterChosen = true
			break
		case 5:
			classe = "Ange"
			fmt.Println("Vous avez choisi : Ange")
			characterChosen = true
			break
		case 6:
			classe = "Démon"
			fmt.Println("Vous avez choisi : Démon")
			characterChosen = true
			break
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
		if characterChosen {
			break
		}
	}
	player = character.InitCharacter(name, classe, 1, 100, 100, []character.Item{})
}

func displayCharacterInfo() {
	character.DisplayInfo(&player)
}

func accessInventory() {
	character.AccessInventory(&player)
}

func shopMenu() {
	item.ShopMenu(&player)
}

func blacksmithMenu() {
	fmt.Println("Fonctionnalité forgeron à compléter : achat et amélioration d'objets.")
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

func calculateXPReward(monster karl.Monster) int {
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

func levelUp() {
	xpNeeded := level * 100
	if xp >= xpNeeded {
		level++
		xp -= xpNeeded
		fmt.Printf("\033[34m[Level Up] Félicitations ! Vous êtes maintenant niveau %d !\033[0m\n", level)
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

func chatGoldReward(gold int) {
	fmt.Printf("\033[35m[Chat] Félicitations ! Vous avez gagné %d pièces d'or !\033[0m\n", gold)
}

func combatTraining(level int) {
	monster := getMonsterByLevel(level)
	characterName := player.Nom
	characterHP := player.HPactuel
	gold := player.Gold

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
				player.Gold = gold
				chatGoldReward(goldReward)
				xpReward := calculateXPReward(monster)
				xp += xpReward
				fmt.Printf("\033[35m[XP] Vous avez gagné %d XP !\033[0m\n", xpReward)
				levelUp()
				fmt.Printf("\033[33mOr total : %d\033[0m\n", gold)
				return
			}

			fmt.Printf("\033[31m%s contre-attaque et inflige %d dégâts !\033[0m\n", monster.Nom, monster.Attaque)
			characterHP -= monster.Attaque
			player.HPactuel = characterHP
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
