package character

import (
	"fmt"
	"projet_red/item"
)

// Predefined characters
var Matheo = Character{
	Nom:           "yzeldine",
	Classe:        "Elfe",
	Niveau:        1,
	MaxHP:         100,
	HPactuel:      100,
	Inventaire:    []item.Item{{Name: "Baton Magique", Quantity: 1, Rarity: 0}, {Name: "Potion de santé", Quantity: 1, Rarity: 1}},
	Skill:         []string{"Coup de poing"},
	Gold:          100,
	MaxInventaire: 10,
	Equipement: struct {
		Tete  string
		Torse string
		Pieds string
	}{},
	Ressources:                map[string]int{},
	NbAugmentationsInventaire: 0,
	Or:                        0,
}

var Ilian = Character{
	Nom:           "Ilian",
	Classe:        "Humain",
	Niveau:        1,
	MaxHP:         120,
	HPactuel:      120,
	Inventaire:    []item.Item{{Name: "Épée basique", Quantity: 1, Rarity: 0}},
	Skill:         []string{"Coup d'épée"},
	Gold:          80,
	MaxInventaire: 10,
	Equipement: struct {
		Tete  string
		Torse string
		Pieds string
	}{},
	Ressources:                map[string]int{},
	NbAugmentationsInventaire: 0,
	Or:                        0,
}

var Karl = Character{
	Nom:           "Karl",
	Classe:        "Nain",
	Niveau:        1,
	MaxHP:         150,
	HPactuel:      150,
	Inventaire:    []item.Item{{Name: "Hache", Quantity: 1, Rarity: 1}},
	Skill:         []string{"Coup de hache"},
	Gold:          60,
	MaxInventaire: 10,
	Equipement: struct {
		Tete  string
		Torse string
		Pieds string
	}{},
	Ressources:                map[string]int{},
	NbAugmentationsInventaire: 0,
	Or:                        0,
}

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
	c1 := InitCharacter("VotreNom", "Elfe", 1, 100, 40, []item.Item{item.MagicStaff, item.HealthPotion})
	fmt.Println("Personnage c1 :", c1)
	return c1
}
