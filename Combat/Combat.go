// Fichier renommé en combat.go pour respecter la convention Go
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
	for c.HP > 0 && m.HPactuel > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		CharTurn(c, m)
		if m.HPactuel <= 0 {
			fmt.Printf("Vous avez vaincu %s !\n", m.Nom)
			break
		}
		GoblinPattern(c, m)
		if c.HP <= 0 {
			fmt.Println("Vous avez été vaincu !")
			break
		}
		tour++
	}
	fmt.Println("Fin du combat d'entraînement. Retour au menu principal.")
}
