package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []string
}

func initCharacter(name, class string, level, maxHP, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
}

func displayInfo(c Character) {
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nHP: %d/%d\nInventaire: %v\n",
		c.Name, c.Class, c.Level, c.CurrentHP, c.MaxHP, c.Inventory)
}

func accessInventory(c *Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide")
		return
	}
	for i, item := range c.Inventory {
		fmt.Printf("%d - %s\n", i+1, item)
	}
	fmt.Println("0 - Retour")
	var choice int
	fmt.Scanln(&choice)
	if choice > 0 && choice <= len(c.Inventory) {
		if c.Inventory[choice-1] == "Potion" {
			takePot(c, choice-1)
		}
	}
}

func takePot(c *Character, index int) {
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Printf("Vous avez bu une potion. HP: %d/%d\n", c.CurrentHP, c.MaxHP)
}

func addInventory(c *Character, item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("%s ajouté à l'inventaire\n", item)
}

func merchant(c *Character) {
	fmt.Println("Marchand :")
	fmt.Println("1 - Potion de vie (gratuit)")
	fmt.Println("0 - Retour")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		addInventory(c, "Potion")
	}
}

func menu(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1 - Afficher les informations")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Marchand")
		fmt.Println("0 - Quitter")
		fmt.Print("Choix: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			displayInfo(*c)
		case "2":
			accessInventory(c)
		case "3":
			merchant(c)
		case "0":
			return
		}
	}
}

func main() {
	c1 := initCharacter("Alan", "Elfe", 1, 100, 40, []string{"Potion", "Potion", "Potion"})
	menu(&c1)
}
