package character

import "fmt"

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

type Item struct {
	Name     string
	Quantity int
	Rarity   int
}

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	MaxHP      int
	HPactuel   int
	Inventaire []Item
	Skill      []string
	Gold       int // pièces d'or
}

var MagicStaff = Item{Name: "Baton Magique", Quantity: 1, Rarity: 0}
var CelestialBlade = Item{Name: "Lame céleste", Quantity: 1, Rarity: 0}
var BattleAxe = Item{Name: "Hache de guerre", Quantity: 1, Rarity: 0}
var KnightsSword = Item{Name: "Épée de chevalier", Quantity: 1, Rarity: 0}
var SteelClaws = Item{Name: "Griffes d'acier", Quantity: 1, Rarity: 0}
var InfernalTrident = Item{Name: "Trident infernal", Quantity: 1, Rarity: 0}

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

func AddSpellBookToInventory(c *Character, spellBookItem Item) {
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

func InitCharacter(nom string, classe string, niveau int, maxHP int, hpActuel int, inventaire []Item) Character {
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

func UseSpellBook(c *Character, spellBookItem Item, spellName string) {
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
	char1 := Character{
		Nom:        "Elyndra",
		Classe:     "Mage",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []Item{MagicStaff},
	}
	fmt.Println("Personnage char1 :", char1)

	char2 := Character{
		Nom:        "Borin",
		Classe:     "Guerrier",
		Niveau:     1,
		MaxHP:      120,
		HPactuel:   120,
		Inventaire: []Item{BattleAxe},
	}
	char3 := Character{
		Nom:        "Arthur",
		Classe:     "Humain",
		Niveau:     1,
		MaxHP:      100,
		HPactuel:   100,
		Inventaire: []Item{KnightsSword},
	}
	char4 := Character{
		Nom:        "Lysandre",
		Classe:     "Hybride",
		Niveau:     1,
		MaxHP:      90,
		HPactuel:   90,
		Inventaire: []Item{SteelClaws},
	}
	char5 := Character{
		Nom:        "Séraphine",
		Classe:     "Ange",
		Niveau:     1,
		MaxHP:      75,
		HPactuel:   75,
		Inventaire: []Item{CelestialBlade},
	}
	char6 := Character{
		Nom:        "Azazel",
		Classe:     "Démon",
		Niveau:     1,
		MaxHP:      120,
		HPactuel:   120,
		Inventaire: []Item{InfernalTrident},
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

	HealthPotion := Item{Name: "Potion de santé", Quantity: 1, Rarity: 0}
	c1 := InitCharacter("VotreNom", "Elfe", 1, 100, 40, []Item{MagicStaff, HealthPotion})
	fmt.Println("Personnage c1 :", c1)
	return c1
}
