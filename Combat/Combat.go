package combat

import (
	"fmt"
	karl "projet_red/Karl"
	"projet_red/character"
)

type Monster = karl.Monster

const (
	red   = "\033[31m"
	green = "\033[32m"
	cyan  = "\033[36m"
	reset = "\033[0m"
)

func InitGoblin() Monster {
	return karl.InitGoblin()
}

func TrainingFight(c *character.Character, m *Monster) {
	tour := 1
	for c.HPactuel > 0 && m.HPactuel > 0 {
		fmt.Printf(cyan+"\n--- Tour %d ---\n"+reset, tour)
		charTurn(c, m)
	}

	if c.HPactuel > 0 {
		fmt.Println("Vous avez vaincu l'ennemi !")
		goldEarned := calculateGoldReward(m)
		c.Gold += goldEarned
		fmt.Printf("Vous gagnez %d pièces d'or !\n", goldEarned)
	} else {
		fmt.Println("Vous avez été vaincu...")
	}
}

func calculateGoldReward(m *Monster) int {
	switch m.Nom {
	case "Gobelin d'entrainement":
		return 10
	case "Ogre":
		return 20
	case "Loup-Garou":
		return 30
	case "Dragon":
		return 50
	default:
		return 5
	}
}

func charTurn(j *character.Character, m *Monster) {
	var choix int
	for {
		fmt.Println(green + "=====================================" + reset)
		fmt.Printf(green+"|         Combat contre : %s         |\n"+reset, m.Nom)
		fmt.Println(green + "=====================================" + reset)
		fmt.Println(cyan + "| 1. Attaquer" + reset)
		fmt.Println(cyan + "| 2. Inventaire" + reset)
		fmt.Println(green + "=====================================" + reset)
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			degats := 5
			m.HPactuel -= degats
			if m.HPactuel < 0 {
				m.HPactuel = 0
			}
			fmt.Printf(red+"%s utilise Attaque basique et inflige %d dégâts à %s\n"+reset, j.Nom, degats, m.Nom)
			fmt.Printf(red+"%s : %d/%d PV\n"+reset, m.Nom, m.HPactuel, m.MaxHP)
			return // Fin du tour du joueur
		case 2:
			character.AccessInventory(j)
			fmt.Print("Choisissez un objet à utiliser (0 pour annuler) : ")
			var choixObj int
			fmt.Scan(&choixObj)
			if choixObj == 0 {
				continue
			}
			if choixObj > 0 && choixObj <= len(j.Inventaire) {
				obj := j.Inventaire[choixObj-1]
				fmt.Printf(cyan+"Vous utilisez %s.\n"+reset, obj.Name)
				// Implement item effects here
				return // Fin du tour du joueur
			} else {
				fmt.Println(red + "Choix invalide." + reset)
				continue
			}
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}
