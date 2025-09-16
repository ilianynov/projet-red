<<<<<<< HEAD
package main

import (
	"fmt"
	karl "projet_red/Karl"
	"projet_red/character"
	"projet_red/item"
)

type Monster struct {
	Nom      string
	MaxHP    int
	HPactuel int
	Attaque  int
}

func ShowMainMenu() {
=======
func ShowAllDesigns() {
	// ANSI color codes for PowerShell (works for most terminals)
>>>>>>> parent of 455598b (0ughwfea)
	cyan := "\x1b[36m"
	green := "\x1b[32m"
	red := "\x1b[31m"
	reset := "\x1b[0m"
	fmt.Println(cyan + "\n      ______\n   .-        -.\n  /            \\n |,  .-.  .-.  ,|\n | )(_o/  \\o_)( |\n |/     /\\     \\|\n (_     ^^     _)\n  \\__|IIIIII|__/\n   | \\IIIIII/ |\n   \\          /\n    `--------`\nje vous croyais plus fort... Résurrection ?\n" + reset)
	fmt.Println(green + "\n================ RPG ULTIME ================\n" + reset)
	fmt.Println(cyan + "1. Créer un personnage" + reset)
	fmt.Println(cyan + "2. Afficher les infos du personnage" + reset)
	fmt.Println(cyan + "3. Accéder à l'inventaire" + reset)
	fmt.Println(cyan + "4. Boutique (Marchand)" + reset)
	fmt.Println(cyan + "5. Forgeron" + reset)
	fmt.Println(cyan + "6. Entrainement (Combat)" + reset)
	fmt.Println(cyan + "7. Livres de sorts" + reset)
	fmt.Println(cyan + "8. Spécial (Potion/Poison)" + reset)
	fmt.Println(cyan + "9. Démo fonctions avancées" + reset)
	fmt.Println(cyan + "0. Quitter" + reset)
	fmt.Println(green + "============================================" + reset)
	fmt.Println(cyan + "Astuce: Tapez 0 à tout moment pour quitter un menu." + reset)
	fmt.Println(green + "\n=====================================" + reset)
	fmt.Println(cyan + "| Entrez le nom de votre personnage |" + reset)
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(cyan + "| Bienvenue dans notre monde, <nom> ! |" + reset)
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(cyan + "  _______________________________________________________________" + reset)
	fmt.Println(cyan + " /                                                               \\" + reset)
	fmt.Println(cyan + "|                    //===================\\                      |" + reset)
	fmt.Println(cyan + "|      CHOISISSEZ VOTRE CLASSE                                   |" + reset)
	fmt.Println(cyan + "|        1. Humain        4. Nain                                |" + reset)
	fmt.Println(cyan + "|        2. Loup-garou    5. Ange                                |" + reset)
	fmt.Println(cyan + "|        3. Hybride       6. Démon                               |" + reset)
	fmt.Println(cyan + "\\_______________________________________________________________/" + reset)
	fmt.Println(green + "\n=====================================" + reset)
	fmt.Println(cyan + "|        Bienvenue chez le Forgeron !         |" + reset)
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(cyan + "| 1. Épée basique (10 Or)           |" + reset)
	fmt.Println(cyan + "| 2. Épée rare (20 Or)              |" + reset)
	fmt.Println(cyan + "| 3. Épée épique (25 Or)            |" + reset)
	fmt.Println(cyan + "| 4. Épée légendaire (40 Or)        |" + reset)
	fmt.Println(cyan + "| 5. Épée démoniaque (60 Or)        |" + reset)
	fmt.Println(cyan + "| 6. Épée angélique (60 Or)         |" + reset)
	fmt.Println(cyan + "| 7. Armure rare (30 Or)            |" + reset)
	fmt.Println(cyan + "| 8. Armure épique (50 Or)          |" + reset)
	fmt.Println(cyan + "| 9. Armure légendaire (70 Or)      |" + reset)
	fmt.Println(cyan + "| 10. Armure démoniaque (90 Or)     |" + reset)
	fmt.Println(cyan + "| 11. Armure angélique (90 Or)      |" + reset)
	fmt.Println(cyan + "| 12. Petite bombe (30 Or)          |" + reset)
	fmt.Println(cyan + "| 13. Grosse bombe (50 Or)          |" + reset)
	fmt.Println(cyan + "| 14. Bombe nucléaire (100 Or)      |" + reset)
	fmt.Println(cyan + "| 15. Quitter                       |" + reset)
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(green + "\n================ BOUTIQUE =================" + reset)
	fmt.Println(cyan + "Or actuel : <montant>" + reset)
	fmt.Println(cyan + "-------------------------------------------" + reset)
	fmt.Println(cyan + " 1. <item> <prix> or  [Rareté: <niveau>]" + reset)
	fmt.Println(cyan + " ..." + reset)
	fmt.Println(cyan + "-------------------------------------------" + reset)
	fmt.Println(cyan + " 0. Quitter" + reset)
	fmt.Println(green + "\n=====================================" + reset)
	fmt.Println(cyan + "|         Combat contre : <nom>     |" + reset)
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(cyan + "| 1. Attaquer" + reset)
	fmt.Println(cyan + "| 2. Inventaire" + reset)
	fmt.Println(green + "=====================================" + reset)
	fmt.Println(cyan + "Votre choix :" + reset)
}
}
package main

import (
	"fmt"
	karl "projet_red/Karl"
	"projet_red/character"
	"projet_red/combat"
	"projet_red/item"
)

func main() {
	var c1 character.Character
	var created bool = false
		ShowAllDesigns()
		cyan := "\x1b[36m"
		green := "\x1b[32m"
		red := "\x1b[31m"
		reset := "\x1b[0m"
<<<<<<< HEAD
		fmt.Print(cyan + "Votre choix : " + reset)
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			c1 = character.CharacterCreation()
			created = true
		case 2:
			if created {
				character.DisplayInfo(&c1)
			} else {
				fmt.Println("Créez d'abord un personnage !")
			}
		case 3:
			if created {
				character.AccessInventory(&c1)
			} else {
				fmt.Println("Créez d'abord un personnage !")
			}
		case 4:
			if created {
				item.ShopMenu(&c1)
			} else {
				fmt.Println("Créez d'abord un personnage !")
			}
		case 5:
			if created {
				karl.MenuForgeron(&c1)
			} else {
				fmt.Println("Créez d'abord un personnage !")
			}
		case 6:
			if created {
				fmt.Println(green + "\nChoisissez votre adversaire :" + reset)
				fmt.Println(cyan + "1. Gobelin d'entrainement" + reset)
				fmt.Println(cyan + "2. Ogre" + reset)
				fmt.Println(cyan + "3. Loup-Garou" + reset)
				fmt.Println(cyan + "4. Dragon" + reset)
				fmt.Print(cyan + "Votre choix : " + reset)
				var monstreChoix int
				fmt.Scan(&monstreChoix)
				var monstre Monster
				switch monstreChoix {
				case 1:
					monstre = Monster{"Gobelin d'entrainement", 40, 40, 5}
				case 2:
					monstre = Monster{"Ogre", 60, 60, 10}
				case 3:
					monstre = Monster{"Loup-Garou", 80, 80, 15}
				case 4:
					monstre = Monster{"Dragon", 150, 150, 25}
				default:
					fmt.Println("Choix invalide, combat contre le Gobelin par défaut.")
					monstre = Monster{"Gobelin d'entrainement", 40, 40, 5}
				}
				fmt.Println(green+"Combat contre : "+reset, monstre.Nom)
				// combat.TrainingFight(&c1, &monstre) // Uncomment if compatible
			} else {
				fmt.Println("Créez d'abord un personnage !")
			}
		case 7:
			if created {
				character.ShowSpellBooksMenu(&c1)
			} else {
				fmt.Println("Créez d'abord un personnage !")
			}
		case 8:
			if created {
				fmt.Println(cyan + "1. Utiliser une potion empoisonnée" + reset)
				fmt.Println(cyan + "2. Vérifier si le personnage est mort" + reset)
				fmt.Print(cyan + "Votre choix : " + reset)
				var sub int
				fmt.Scan(&sub)
				switch sub {
				case 1:
					fmt.Println("Potion empoisonnée non implémentée dans ce menu.")
				case 2:
					fmt.Println("Fonction de vérification de mort non implémentée dans ce menu.")
				}
			} else {
				fmt.Println("Créez d'abord un personnage !")
			}
		case 9:
			fmt.Println(green + "\n--- Fonctions avancées et tests ---" + reset)
			fmt.Println(cyan + "1. Gérer l'or (Spend/Earn)" + reset)
			fmt.Println(cyan + "2. Ajouter un livre de sort à l'inventaire" + reset)
			fmt.Println(cyan + "3. Ajouter un item à l'inventaire" + reset)
			fmt.Println(cyan + "4. Bonus équipement (Karl)" + reset)
			fmt.Println(cyan + "5. Quitter ce menu" + reset)
=======
		for {
			fmt.Println(green + "\n================ RPG ULTIME ================" + reset)
			fmt.Println(cyan + "1. Créer un personnage" + reset)
			fmt.Println(cyan + "2. Afficher les infos du personnage" + reset)
			fmt.Println(cyan + "3. Accéder à l'inventaire" + reset)
			fmt.Println(cyan + "4. Boutique (Marchand)" + reset)
			fmt.Println(cyan + "5. Forgeron" + reset)
			fmt.Println(cyan + "6. Entrainement (Combat)" + reset)
			fmt.Println(cyan + "7. Livres de sorts" + reset)
			fmt.Println(cyan + "8. Spécial (Potion/Poison)" + reset)
			fmt.Println(cyan + "9. Démo fonctions avancées" + reset)
			fmt.Println(cyan + "0. Quitter" + reset)
			fmt.Println(green + "============================================" + reset)
			fmt.Println(cyan + "Astuce: Tapez 0 à tout moment pour quitter un menu." + reset)
>>>>>>> parent of 455598b (0ughwfea)
			fmt.Print(cyan + "Votre choix : " + reset)
			var choix int
			fmt.Scan(&choix)

			switch choix {
			case 1:
				c1 = character.CharacterCreation()
				created = true
			case 2:
				if created {
					character.DisplayInfo(&c1)
				} else {
					fmt.Println(red + "Créez d'abord un personnage !" + reset)
				}
			case 3:
				if created {
					character.AccessInventory(&c1)
				} else {
					fmt.Println(red + "Créez d'abord un personnage !" + reset)
				}
			case 4:
<<<<<<< HEAD
				fmt.Println(green+"Bonus équipement : "+reset, karl.BonusEquipement(karl.Equipment{}))
=======
				if created {
					item.ShopMenu(&c1)
				} else {
					fmt.Println(red + "Créez d'abord un personnage !" + reset)
				}
			case 5:
				if created {
					karl.MenuForgeron(&c1)
				} else {
					fmt.Println(red + "Créez d'abord un personnage !" + reset)
				}
			case 6:
				if created {
					gob := combat.InitGoblin()
					combat.TrainingFight(&c1, &gob)
				} else {
					fmt.Println(red + "Créez d'abord un personnage !" + reset)
				}
			case 7:
				if created {
					character.ShowSpellBooksMenu(&c1)
				} else {
					fmt.Println(red + "Créez d'abord un personnage !" + reset)
				}
			case 8:
				if created {
					fmt.Println(cyan + "1. Utiliser une potion empoisonnée" + reset)
					fmt.Println(cyan + "2. Vérifier si le personnage est mort" + reset)
					fmt.Print(cyan + "Votre choix : " + reset)
					var sub int
					fmt.Scan(&sub)
					switch sub {
					case 1:
						fmt.Println(red + "Potion empoisonnée non implémentée dans ce menu." + reset)
					case 2:
						fmt.Println(red + "Fonction de vérification de mort non implémentée dans ce menu." + reset)
					}
				} else {
					fmt.Println(red + "Créez d'abord un personnage !" + reset)
				}
			case 9:
				fmt.Println(green + "\n--- Fonctions avancées et tests ---" + reset)
				fmt.Println(cyan + "1. Gérer l'or (Spend/Earn)" + reset)
				fmt.Println(cyan + "2. Ajouter un livre de sort à l'inventaire" + reset)
				fmt.Println(cyan + "3. Ajouter un item à l'inventaire" + reset)
				fmt.Println(cyan + "4. Bonus équipement (Karl)" + reset)
				fmt.Println(cyan + "5. Quitter ce menu" + reset)
				fmt.Print(cyan + "Votre choix : " + reset)
				var adv int
				fmt.Scan(&adv)
				switch adv {
				case 1:
					character.SpendGold(&c1, 10)
					character.EarnGold(&c1, 20)
				case 2:
					character.AddSpellBookToInventory(&c1, character.MagicStaff)
				case 3:
					item.AddItemToInventory(&c1, item.HealthPotion)
				case 4:
					fmt.Println(green + "Bonus équipement : " + reset, karl.BonusEquipement(karl.Equipment{}))
				}
			case 0:
				fmt.Println(green + "\n      ______\n   .-        -.\n  /            \\n |,  .-.  .-.  ,|\n | )(_o/  \\o_)( |\n |/     /\\     \\|\n (_     ^^     _)\n  \\__|IIIIII|__/\n   | \\IIIIII/ |\n   \\          /\n    `--------`\nMerci d'avoir joué à RPG ULTIME !" + reset)
				fmt.Println(green + "Au revoir !" + reset)
				return
			default:
				fmt.Println(red + "Choix invalide." + reset)
>>>>>>> parent of 455598b (0ughwfea)
			}
		}
}
