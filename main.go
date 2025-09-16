package main
import (
	"fmt"
	"projet_red/character" // Importing character package
	"projet_red/item"
	karl "projet_red/Karl"
)

func main() {
	fmt.Println("Bienvenue dans RPG ULTIME!")
	var current *character.Character
	for {
		fmt.Println("\n================ Sélection du personnage ================")
		fmt.Println("1. Matheo (Elfe)")
		fmt.Println("2. Ilian (Humain)")
		fmt.Println("3. Karl (Nain)")
		fmt.Println("4. Créer un personnage personnalisé")
		fmt.Println("0. Quitter le jeu")
		fmt.Print("Votre choix : ")
		var choixPerso int
		fmt.Scan(&choixPerso)
		switch choixPerso {
		case 1:
			current = &character.Matheo
		case 2:
			for {
				fmt.Println("\n================ Sélection du personnage ================")
				fmt.Println("1. Matheo (Elfe)")
				fmt.Println("2. Ilian (Humain)")
				fmt.Println("3. Karl (Nain)")
				fmt.Println("4. Créer un personnage personnalisé")
				fmt.Println("0. Quitter le jeu")
				fmt.Print("Votre choix : ")
				var choixPerso int
				fmt.Scan(&choixPerso)
				switch choixPerso {
				case 1:
					current = &character.Matheo
				case 2:
					current = &character.Ilian
				case 3:
					current = &character.Karl
				case 4:
					custom := character.CharacterCreation()
					current = &custom
				case 0:
					fmt.Println("Merci d'avoir joué à RPG ULTIME ! Au revoir !")
					return
				default:
					fmt.Println("Choix invalide.")
					continue
				}
			}
			switch choix {
			case 1:
				character.DisplayInfo(current)
			case 2:
				character.AccessInventory(current)
			case 3:
				item.ShopMenu(current)
			case 4:
				karl.MenuForgeron(current)
			case 5:
				fmt.Println("Choisissez votre adversaire :")
				fmt.Println("1. Gobelin d'entrainement")
				fmt.Println("2. Ogre")
				fmt.Println("3. Loup-Garou")
				fmt.Println("4. Dragon")
				fmt.Print("Votre choix : ")
				var monstreChoix int
				fmt.Scan(&monstreChoix)
				var monstre karl.Monster
				fmt.Println("6. Livres de sorts")
				fmt.Println("9. Changer de personnage")
				fmt.Println("0. Quitter le jeu")
				fmt.Print("Votre choix : ")
				var choix int
				fmt.Scan(&choix)
				switch choix {
				case 1:
					ShowDetailedCharacter(current)
				case 2:
					character.AccessInventory(current)
				case 3:
					item.ShopMenu(current)
				case 4:
					karl.MenuForgeron(current)
				case 5:
					importCombat(current)
				case 6:
					fmt.Println("Forgeron non implémenté dans ce menu.")
				case 0:
					fmt.Println("Merci d'avoir joué à RPG ULTIME ! Au revoir !")
					return
				default:
					fmt.Println("Choix invalide.")
				}
				if choix == 9 {
					break
				}
		default:
			fmt.Println("Choix invalide.")
			continue
		}
	}
		for {
			fmt.Println("\n================ MENU PRINCIPAL ================")
			fmt.Println("1. Voir les détails du personnage")
			fmt.Println("2. Accéder à l'inventaire")
			fmt.Println("3. Boutique")
			fmt.Println("4. Forgeron")
			fmt.Println("5. Combat d'entraînement")
			fmt.Println("6. Livres de sorts")
			fmt.Println("9. Changer de personnage")
			fmt.Println("0. Quitter le jeu")
			fmt.Print("Votre choix : ")
			var choix int
			fmt.Scan(&choix)
			switch choix {
			case 1:
				ShowDetailedCharacter(current)
			case 2:
				character.AccessInventory(current)
			case 3:
				importItemShop(current)
			case 4:
				importForgeron(current)
			case 5:
				importCombat(current)
			case 6:
		case 3:
			item.ShopMenu(current)
				break
			fmt.Println("Forgeron non implémenté dans ce menu.")
				fmt.Println("Merci d'avoir joué à RPG ULTIME ! Au revoir !")
				return
			default:
				fmt.Println("Choix invalide.")
			}
			if choix == 9 {
				break
			}
		}
	}

	// ShowDetailedCharacter displays all details of the character with explanations
	func ShowDetailedCharacter(c *character.Character) {
		fmt.Println("================ Détails du Personnage ================")
		fmt.Printf("Nom: %s\n", c.Nom)
		fmt.Printf("Classe: %s\n", c.Classe)
		fmt.Printf("Niveau: %d\t// Niveau d'expérience du personnage\n", c.Niveau)
		fmt.Printf("Points de vie (HP): %d/%d\t// Points de vie actuels / maximum\n", c.HPactuel, c.MaxHP)
		fmt.Printf("Or: %d\t// Richesse du personnage\n", c.Gold)
		fmt.Printf("Inventaire:\n")
		if len(c.Inventaire) == 0 {
			fmt.Println("  (vide)")
		} else {
			for _, it := range c.Inventaire {
				fmt.Printf("  - %s\n    Quantité: %d\n    Rareté: %d\n    Prix: %d\n", it.Name, it.Quantity, it.Rarity, it.Price)
			}
		}
	}

	func importCombat(c *character.Character) {
		// Combat d'entraînement contre un monstre
		fmt.Println("Choisissez votre adversaire :")
		fmt.Println("1. Gobelin d'entrainement")
		fmt.Println("2. Ogre")
		fmt.Println("3. Loup-Garou")
		fmt.Println("4. Dragon")
		fmt.Print("Votre choix : ")
		var monstreChoix int
		fmt.Scan(&monstreChoix)
		var m karl.Monster
		switch monstreChoix {
		case 1:
			m = karl.InitGoblin()
		case 2:
			m = karl.InitOgre()
		case 3:
			m = karl.InitLoupGarou()
		case 4:
			m = karl.InitDragon()
		default:
			fmt.Println("Choix invalide, combat contre le Gobelin par défaut.")
			m = karl.InitGoblin()
		}
		// You can call a combat function here if needed
		fmt.Println("Fin du combat d'entraînement. Retour au menu principal.")
	}
func importCombat(c *character.Character) {
	// Combat d'entraînement contre un monstre
	fmt.Println("Choisissez votre adversaire :")
	fmt.Println("1. Gobelin d'entrainement")
	fmt.Println("2. Ogre")
	fmt.Println("3. Loup-Garou")
	fmt.Println("4. Dragon")
	fmt.Print("Votre choix : ")
	var monstreChoix int
	fmt.Scan(&monstreChoix)
	var m Monster
	switch monstreChoix {
	case 1:
		m = initGoblin()
	case 2:
		m = initOgre()
	case 3:
		m = initLoupGarou()
	case 4:
		m = initDragon()
	default:
		fmt.Println("Choix invalide, combat contre le Gobelin par défaut.")
		m = initGoblin()
	}
	TrainingFight(c, &m)
	fmt.Println("Fin du combat d'entraînement. Retour au menu principal.")
}


// ShowDetailedCharacter displays all details of the character with explanations
func ShowDetailedCharacter(c *character.Character) {
	fmt.Println("================ Détails du Personnage ================")
	fmt.Printf("Nom: %s\n", c.Nom)
	fmt.Printf("Classe: %s\n", c.Classe)
	fmt.Printf("Niveau: %d\t// Niveau d'expérience du personnage\n", c.Niveau)
	fmt.Printf("Points de vie (HP): %d/%d\t// Points de vie actuels / maximum\n", c.HPactuel, c.MaxHP)
	fmt.Printf("Or: %d\t// Richesse du personnage\n", c.Gold)
	fmt.Printf("Inventaire:\n")
	if len(c.Inventaire) == 0 {
		fmt.Println("  (vide)")
	} else {
		for _, it := range c.Inventaire {
			fmt.Printf("  - %s\n    Quantité: %d\n    Rareté: %d\n    Prix: %d\n", it.Name, it.Quantity, it.Rarity, it.Price)
		}
	}
}
