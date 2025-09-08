package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxHP      int
	HPactuel   int
	Inventaire []string
}

// Nouvelle fonction pour initialiser un personnage
func initCharacter(nom string, classe string, niveau int, maxHP int, hpActuel int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		MaxHP:      maxHP,
		HPactuel:   hpActuel,
		Inventaire: inventaire,
	}
}

func CharacterCreation() {
	char1 := Character{
		Nom:        "Elfe",
		Classe:     "Mage",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []string{},
	}
	fmt.Println("Personnage char1 :", char1)

	// Création d'un autre personnage c1 avec initCharacter
	c1 := initCharacter("VotreNom", "Elfe", 1, 100, 40, []string{"Potion", "Potion", "Potion"})
	fmt.Println("Personnage c1 :", c1)
}

func main() {
	CharacterCreation()
}
