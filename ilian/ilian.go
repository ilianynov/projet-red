package main

import (
	"math/rand"
	"time"
)

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxHP      int
	HPactuel   int
	Inventaire []string
}

func randomClasse() string {
	classes := []string{"Humain", "Loupgarou", "Hybride", "Nain", "Ange", "Démon"}
	rand.Seed(time.Now().UnixNano())
	return classes[rand.Intn(len(classes))]
}

func main() {
	hero := Character{
		Nom:        "Ynov",
		Classe:     randomClasse(),
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []string{"Livre"},
	}
}
