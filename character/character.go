package main

// No imports needed for this file currently
import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxHP      int
	HPactuel   int
	Inventaire []item
}

func CharacterCreation() {
	char1 := Character{
		Nom:        "Elyndra",
		Classe:     "Mage",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []item{MagicStaff},
	}
	char2 := Character{
		Nom:        "Borin",
		Classe:     "Guerrier",
		Niveau:     1,
		MaxHP:      120,
		HPactuel:   120,
		Inventaire: []item{BattleAxe},
	}
	char3 := Character{
		Nom:        "Arthur",
		Classe:     "Humain",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []item{KnightsSword},
	}
	char4 := Character{
		Nom:        "Lysandre",
		Classe:     "Hybride",
		Niveau:     1,
		MaxHP:      90,
		HPactuel:   90,
		Inventaire: []item{SteelClaws},
	}
	char5 := Character{
		Nom:        "Séraphine",
		Classe:     "Ange",
		Niveau:     1,
		MaxHP:      75,
		HPactuel:   75,
		Inventaire: []item{CelestialBlade},
	}
	char6 := Character{
		Nom:        "Azazel",
		Classe:     "Démon",
		Niveau:     1,
		MaxHP:      120,
		HPactuel:   120,
		Inventaire: []item{InfernalTrident},
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

type item struct {
	name     string
	quantity int
	rarity   int

}

var MagicStaff = item{name:"Baton Magique" , quantity: 1, rarity: 0}
var CelestialBlade = item{name: "Lame céleste" , quantity: 1, rarity: 0}
var BattleAxe = item{name: "Hache de guerre" , quantity: 1, rarity 0}
var KnightsSword = item{name: "Épée de chevalier" , quantity: 1, rarity 0}
var SteelClaws = item{name: "Griffes d'acier" , quantity: 1, rarity 0}
var InfernalTrident = item{name: "Trident infernal" , quantity 1, rarity 0}


