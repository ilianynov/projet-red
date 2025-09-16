package character

import (
	"fmt"
	"projet_red/item"
)

// Affiche les infos du personnage
func DisplayInfo(c *Character) {
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nHP: %d/%d\nInventaire: %v\nCompétences: %v\nOr: %d\n",
		c.Nom, c.Classe, c.Niveau, c.HPactuel, c.MaxHP, c.Inventaire, c.Skill, c.Gold)
}

// Affiche l'inventaire du personnage
func AccessInventory(c *Character) {
	fmt.Println("Inventaire:")
	for i, it := range c.Inventaire {
		fmt.Printf("%d. %s x%d\n", i+1, it.Name, it.Quantity)
	}
}

type Character struct {
	Nom           string
	Classe        string
	Niveau        int
	MaxHP         int
	HPactuel      int
	Inventaire    []item.Item
	Skill         []string
	Gold          int // pièces d'or
	MaxInventaire int
	Equipement    struct {
		Tete  string
		Torse string
		Pieds string
	}
	Ressources                map[string]int
	NbAugmentationsInventaire int
	Or                        int
}

// Methods for item package compatibility
func (c *Character) GetInventaire() *[]item.Item {
	return &c.Inventaire
}

func (c *Character) GetGold() *int {
	return &c.Gold
}

// Items are now imported from item package

func SpendGold(c *Character, price int) bool {
	if c.Gold >= price {
		c.Gold -= price
		fmt.Println("Achat réussi !")
		return true
	} else {
		fmt.Println("Pas assez d'or !")
		return false
	}
}

func EarnGold(c *Character, amount int) {
	c.Gold += amount
	fmt.Printf("Vous avez gagné %d pièces d'or !\n", amount)
}

func AddSpellBookToInventory(c *Character, spellBookItem item.Item) {
	c.Inventaire = append(c.Inventaire, spellBookItem)
	fmt.Println("Livre de sort ajouté à l'inventaire :", spellBookItem.Name)
}

func ShowSpellBooksMenu(c *Character) {
	spellBookItems := map[string]string{
		"Livre de Sort : Boule de Feu":    "Boule de feu",
		"Livre de Sort : Éclair de Glace": "Éclair de Glace",
		"Livre de Sort : Soin Rapide":     "Soin Rapide",
	}
	found := false
	for i, it := range c.Inventaire {
		if _, ok := spellBookItems[it.Name]; ok {
			found = true
			fmt.Printf("%d. %s\n", i+1, it.Name)
		}
	}
	if !found {
		fmt.Println("Aucun livre de sort dans l'inventaire.")
		return
	}
	var choice int
	fmt.Print("Entrez le numéro du livre de sort à utiliser (0 pour annuler) : ")
	fmt.Scan(&choice)
	if choice > 0 && choice <= len(c.Inventaire) {
		selected := c.Inventaire[choice-1]
		if spell, ok := spellBookItems[selected.Name]; ok {
			UseSpellBook(c, selected, spell)
		} else {
			fmt.Println("Ce n'est pas un livre de sort valide.")
		}
	} else {
		fmt.Println("Annulé.")
	}
}

func InitCharacter(nom string, classe string, niveau int, maxHP int, hpActuel int, inventaire []item.Item) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		MaxHP:      maxHP,
		HPactuel:   hpActuel,
		Inventaire: inventaire,
		Skill:      []string{"Coup de poing"},
		Gold:       100,
	}
}

func SpellBook(c *Character, spell string) {
	for _, s := range c.Skill {
		if s == spell {
			fmt.Println("Ce sort est déjà appris.")
			return
		}
	}
	c.Skill = append(c.Skill, spell)
	fmt.Println("Vous avez appris le sort :", spell)
}

func UseSpellBook(c *Character, spellBookItem item.Item, spellName string) {
	var input string
	fmt.Printf("Voulez-vous apprendre le sort '%s' ? (o/n): ", spellName)
	fmt.Scan(&input)
	if input == "o" || input == "O" {
		SpellBook(c, spellName)
	} else {
		fmt.Println("Vous n'avez pas appris le sort.")
	}
}

func CharacterCreation() Character {
	var nom string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scan(&nom)

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Humain")
	fmt.Println("2. Loup-garou")
	fmt.Println("3. Hybride")
	fmt.Println("4. Nain")
	fmt.Println("5. Ange")
	fmt.Println("6. Démon")
	var classeChoix int
	fmt.Print("Votre choix : ")
	fmt.Scan(&classeChoix)

	classe := "Elfe"
	maxHP := 100
	hpActuel := 40
	inventaire := []item.Item{item.MagicStaff, item.HealthPotion}

	switch classeChoix {
	case 1:
		classe = "Humain"
		maxHP = 120
		inventaire = []item.Item{item.BasicSword, item.HealthPotion}
	case 2:
		classe = "Loup-garou"
		maxHP = 140
		inventaire = []item.Item{item.SteelClaws, item.HealthPotion}
	case 3:
		classe = "Hybride"
		maxHP = 110
		inventaire = []item.Item{item.CelestialBlade, item.HealthPotion}
	case 4:
		classe = "Nain"
		maxHP = 130
		inventaire = []item.Item{item.BattleAxe, item.HealthPotion}
	case 5:
		classe = "Ange"
		maxHP = 100
		inventaire = []item.Item{item.KnightsSword, item.HealthPotion}
	case 6:
		classe = "Démon"
		maxHP = 150
		inventaire = []item.Item{item.InfernalTrident, item.HealthPotion}
	}

	c1 := InitCharacter(nom, classe, 1, maxHP, hpActuel, inventaire)
	fmt.Println("Personnage créé :", c1)
	return c1
}
