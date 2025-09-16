package combat

import (
	"fmt"
	karl "projet_red/Karl"
	"projet_red/character"
)

type Monster = karl.Monster

func InitGoblin() Monster {
	return karl.InitGoblin()
}

func goblinPattern(j *character.Character) {
	karl.GoblinPattern(j)
}

func TrainingFight(c *character.Character, m *Monster) {
	tour := 1
	for c.HPactuel > 0 && m.HPactuel > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		charTurn(c, m)
	}
}

func charTurn(j *character.Character, m *Monster) {
	var choix int
	for {
		fmt.Println("=====================================")
		fmt.Println("|         Combat contre :", m.Nom)
		fmt.Println("=====================================")
		fmt.Println("| 1. Attaquer")
		fmt.Println("| 2. Inventaire")
		fmt.Println("=====================================")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			degats := 5
			m.HPactuel -= degats
			if m.HPactuel < 0 {
				m.HPactuel = 0
			}
			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s\n", j.Nom, degats, m.Nom)
			fmt.Printf("%s : %d/%d PV\n", m.Nom, m.HPactuel, m.MaxHP)
			return // Fin du tour du joueur
		case 2:
			if len(j.Inventaire) == 0 {
				fmt.Println("Votre inventaire est vide.")
				continue
			}
			fmt.Println("Inventaire :")
			for i, obj := range j.Inventaire {
				fmt.Printf("%d. %s\n", i+1, obj.Name)
			}
			fmt.Print("Choisissez un objet à utiliser (0 pour annuler) : ")
			var choixObj int
			fmt.Scan(&choixObj)
			if choixObj == 0 {
				continue
			}
			if choixObj > 0 && choixObj <= len(j.Inventaire) {
				karl.UtiliserObjet(j, m, j.Inventaire[choixObj-1])
				return // Fin du tour du joueur
			} else {
				fmt.Println("Choix invalide.")
				continue
			}
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
