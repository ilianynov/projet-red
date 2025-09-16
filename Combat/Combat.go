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

func TrainingFight(c *character.Character, m *Monster) {
	tour := 1
	for c.HPactuel > 0 && m.HPactuel > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		fmt.Printf("%s : %d/%d PV\n", c.Nom, c.HPactuel, c.MaxHP)
		fmt.Printf("%s : %d/%d PV\n", m.Nom, m.HPactuel, m.MaxHP)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire (non implémenté)")
		fmt.Print("Votre choix : ")
		var choix int
		fmt.Scan(&choix)
		if choix == 1 {
			degats := 10 // dégâts du joueur
			m.HPactuel -= degats
			fmt.Printf("Vous attaquez et infligez %d dégâts à %s!\n", degats, m.Nom)
		} else {
			fmt.Println("Inventaire non implémenté, vous perdez votre tour!")
		}
		if m.HPactuel <= 0 {
			fmt.Printf("Vous avez vaincu %s !\n", m.Nom)
			break
		}
		// Tour du monstre
		degatsMonstre := m.Attaque
		c.HPactuel -= degatsMonstre
		fmt.Printf("%s attaque et inflige %d dégâts à %s!\n", m.Nom, degatsMonstre, c.Nom)
		if c.HPactuel <= 0 {
			fmt.Println("Vous avez été vaincu !")
			break
		}
		tour++
	}
	fmt.Println("Fin du combat d'entraînement. Retour au menu principal.")
}
