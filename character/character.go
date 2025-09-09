package main

// No imports needed for this file currently
import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxHP      int
	HPactuel   int
	Inventaire []string
}

func CharacterCreation() {
	char1 := Character{
		Nom:        "Elyndra",
		Classe:     "Mage",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []string{},
	}
	char2 := Character{
		Nom:        "Borin",
		Classe:     "Guerrier",
		Niveau:     1,
		MaxHP:      120,
		HPactuel:   120,
		Inventaire: []string{},
	}
	char3 := Character{
		Nom:        "Arthur",
		Classe:     "Humain",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []string{},
	}
	char4 := Character{
		Nom:        "Lysandre",
		Classe:     "Hybride",
		Niveau:     1,
		MaxHP:      90,
		HPactuel:   90,
		Inventaire: []string{},
	}
	char5 := Character{
		Nom:        "Séraphine",
		Classe:     "Ange",
		Niveau:     1,
		MaxHP:      75,
		HPactuel:   75,
		Inventaire: []string{},
	}
	char6 := Character{
		Nom:        "Azazel",
		Classe:     "Démon",
		Niveau:     1,
		MaxHP:      120,
		HPactuel:   120,
		Inventaire: []string{},
	}

	fmt.Println(char1)
	fmt.Println(char2)
	fmt.Println(char3)
	fmt.Println(char4)
	fmt.Println(char5)
	fmt.Println(char6)
	fmt.Println("Personnages créés avec succès !")
	fmt.Println("Nom:", char1.Nom, "Classe:", char1.Classe, "Niveau:", char1.Niveau, "MaxHP:", char1.MaxHP, "HPactuel:", char1.HPactuel)
	fmt.Println("Nom:", char2.Nom, "Classe:", char2.Classe, "Niveau:", char2.Niveau, "MaxHP:", char2.MaxHP, "HPactuel:", char2.HPactuel)

}
