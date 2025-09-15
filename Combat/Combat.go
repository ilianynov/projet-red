package main

import "fmt"

// Supposé :
// func charTurn(c *Character, m *Monster) {}
// func goblinPattern(c *Character, m *Monster) {}
// type Character, type Monster déjà définis

func trainingFight(c *Character, m *Monster) {
	tour := 1
	for c.HP > 0 && m.HPactuel > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		charTurn(c, m)
		if m.HPactuel <= 0 {
			fmt.Printf("Vous avez vaincu %s !\n", m.Nom)
			break
		}
		goblinPattern(c, m)
		if c.HP <= 0 {
			fmt.Println("Vous avez été vaincu !")
			break
		}
		tour++
	}
	fmt.Println("Fin du combat d'entraînement. Retour au menu principal.")
}
