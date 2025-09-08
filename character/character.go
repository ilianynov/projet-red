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
		Nom:        "Elfe",
		Classe:     "Mage",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []string{},
	}
	fmt.Println(char1)
}

func main() {
	CharacterCreation()
}
