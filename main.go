package main

import (
	"fmt"
	"math/rand"
	karl "projet_red/Karl"
	"projet_red/character"
	"projet_red/graphic"
	"projet_red/item"
	"time"
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
		case 0:
			fmt.Println("Merci d'avoir joué !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func showMainMenu() {
	fmt.Println("\n================ RrdYnovLegende ================")
	fmt.Println("1. Créer un personnage")
	fmt.Println("2. Afficher les infos du personnage")
	fmt.Println("3. Accéder à l'inventaire")
	fmt.Println("4. Boutique (Marchand)")
	fmt.Println("5. Forgeron")
	fmt.Println("6. Entrainement (Combat)")
	fmt.Println("7. Livres de sorts")
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
			continue
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
	fmt.Println("Bienvenue chez le marchand !")
	for i, it := range item.ShopItems {
		fmt.Printf("%d. %s (Prix : %d or)\n", i+1, it.Name, it.Price)
	}
	fmt.Println("0. Quitter")
	fmt.Print("Votre choix : ")
	var choix int
	fmt.Scan(&choix)
	if choix > 0 && choix <= len(item.ShopItems) {
		itemAchat := item.ShopItems[choix-1]
		if player.Gold >= itemAchat.Price {
			player.Gold -= itemAchat.Price
			player.Inventaire = append(player.Inventaire, character.Item{Name: itemAchat.Name, Quantity: 1, Rarity: itemAchat.Rarity})
			fmt.Printf("Vous avez acheté : %s\n", itemAchat.Name)
		} else {
			fmt.Println("Vous n'avez pas assez d'or.")
		}
	}
}

func blacksmithMenu() {
	karl.MenuForgeron(&player)
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

func getRandomMonster() karl.Monster {
	monsters := []karl.Monster{
		InitGoblin(),
		InitOgre(),
		InitLoupGarou(),
		InitDragon(),
	}
	rand.Seed(time.Now().UnixNano())
	return monsters[rand.Intn(len(monsters))]
}

func chatGoldReward(gold int) {
	fmt.Printf("\033[35m[Chat] Félicitations ! Vous avez gagné %d pièces d'or !\033[0m\n", gold)
}

func getMonsterMaterial(monster karl.Monster) string {
	switch monster.Nom {
	case "Gobelin d'entrainement":
		return "Peau de gobelin"
	case "Ogre":
		return "Dent d'ogre"
	case "Loup-Garou":
		return "Griffe de loup-garou"
	case "Dragon":
		return "Écaille de dragon"
	default:
		return "Inconnu"
	}
}

func combatTraining(level int) {
	monster := getRandomMonster()
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
		fmt.Println("| 3. Sorts")
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
				material := getMonsterMaterial(monster)
				player.Inventaire = append(player.Inventaire, character.Item{Name: material, Quantity: 1, Rarity: 1})
				fmt.Printf("\033[36mVous avez obtenu : %s\033[0m\n", material)
				player.HPactuel = player.MaxHP // Reset la vie du personnage après victoire
				return
			}

			fmt.Printf("\033[31m%s contre-attaque et inflige %d dégâts !\033[0m\n", monster.Nom, monster.Attaque)
			characterHP -= monster.Attaque
			player.HPactuel = characterHP
			if characterHP <= 0 {
				graphic.ShowDeathArt()
				fmt.Println("Vous êtes mort !")
				fmt.Printf("\033[31m%s a été vaincu par %s !\033[0m\n", characterName, monster.Nom)
				player.HPactuel = player.MaxHP // Reset la vie du personnage après défaite
				return
			}
		case 2:
			fmt.Println("Ouverture de l'inventaire...")
			accessInventory()
		case 3:
			useSpell(&monster, &characterHP, &player)
		case 0:
			fmt.Println("Vous quittez le combat.")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func spellBooksMenu() {
	fmt.Println("========== Livres de sorts ==========")
	fmt.Println("Voici les sorts disponibles :")
	fmt.Println("1. Coup de poing : Inflige 8 dégâts à l'adversaire.")
	fmt.Println("2. Boule de feu : Inflige 18 dégâts à l'adversaire.")
	fmt.Println("3. Soin : Rend 15 PV.")
	fmt.Println("0. Retour")
	fmt.Print("Entrez le numéro d'un sort pour lire sa description, ou 0 pour revenir : ")

	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:
		fmt.Println("Coup de poing : Un sort simple qui inflige 8 dégâts à l'adversaire.")
	case 2:
		fmt.Println("Boule de feu : Un sort puissant qui inflige 18 dégâts à l'adversaire.")
	case 3:
		fmt.Println("Soin : Un sort qui rend 15 points de vie à votre personnage.")
	case 0:
		fmt.Println("Retour au menu principal.")
	default:
		fmt.Println("Choix invalide.")
	}
}

func useSpell(monster *karl.Monster, characterHP *int, player *character.Character) {
	fmt.Println("========== Livres de sorts ==========")
	fmt.Println("1. Coup de poing : Inflige 8 dégâts à l'adversaire.")
	fmt.Println("2. Boule de feu : Inflige 18 dégâts à l'adversaire.")
	fmt.Println("3. Soin : Rend 15 PV.")
	fmt.Println("0. Annuler")
	fmt.Print("Choisissez un sort : ")
	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:
		monster.HPactuel -= 8
		fmt.Printf("Vous lancez Coup de poing et infligez 8 dégâts à %s !\n", monster.Nom)
	case 2:
		monster.HPactuel -= 18
		fmt.Printf("Vous lancez Boule de feu et infligez 18 dégâts à %s !\n", monster.Nom)
	case 3:
		*characterHP += 15
		if *characterHP > player.MaxHP {
			*characterHP = player.MaxHP
		}
		player.HPactuel = *characterHP
		fmt.Printf("Vous lancez Soin et récupérez 15 PV !\n")
	case 0:
		fmt.Println("Sort annulé.")
	default:
		fmt.Println("Choix invalide.")
	}
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
