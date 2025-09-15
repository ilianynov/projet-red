package main

import (
	"fmt"
	"projet_red/character"
)

func goblinPattern(j *Character) {
	goblin := initGoblin()
	tour := 1

	for j.HPactuel > 0 && goblin.HPactuel > 0 {
		var degats int
		if tour%3 == 0 {
			degats = goblin.Attaque * 2
		} else {
			degats = goblin.Attaque
		}

		j.HPactuel -= degats
		if j.HPactuel < 0 {
			j.HPactuel = 0
		}

		fmt.Printf("%s inflige à %s %d de dégats\n", goblin.Nom, j.Nom, degats)
		fmt.Printf("%s : %d/%d PV\n", j.Nom, j.HPactuel, j.MaxHP)

		if j.HPactuel == 0 {
			fmt.Printf("%s est vaincu !\n", j.Nom)
			break
		}
		tour++
	}
}

func charTurn(j *Character, m *Monster) {
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
				fmt.Printf("%d. %s\n", i+1, obj)
			}
			fmt.Print("Choisissez un objet à utiliser (0 pour annuler) : ")
			var choixObj int
			fmt.Scan(&choixObj)
			if choixObj == 0 {
				continue
			}
			if choixObj > 0 && choixObj <= len(j.Inventaire) {
				utiliserObjet(j, m, j.Inventaire[choixObj-1])
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

func trainingFight(c *character.Character, m *Monster) {
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
