package item

import (
	"fmt"
	"projet_red/character"
)

func BuyItem(c *character.Character, it Item) {
	if len(c.Inventaire) >= 10 {
		fmt.Println("Votre inventaire est plein.")
		return
	}
	if c.Gold >= it.Price {
		c.Gold -= it.Price
		converted := character.Item{Name: it.Name, Quantity: it.Quantity, Rarity: it.Rarity}
		c.Inventaire = append(c.Inventaire, converted)
		fmt.Printf("Vous avez acheté %s pour %d pièces d'or.\n", it.Name, it.Price)
	} else {
		fmt.Println("Pas assez d'or pour acheter cet objet.")
	}
}

func ShopMenu(c *character.Character) {

	allItems := []Item{
		BasicSword, HealthPotion, PoisonPotion, SmallBomb, BigBomb, Nuke,
		RareSword, RareArmor, EpicSword, EpicArmor, InventoryUpgrade,
		Uranium, BattleAxe, CelestialBlade, KnightsSword, SteelClaws, InfernalTrident,
	}

	var items []Item
	for _, it := range allItems {
		if it.Rarity >= 3 && it.Rarity <= 5 {
			items = append(items, it)
		}
	}
	for {
		fmt.Println("\n================ BOUTIQUE =================")
		fmt.Printf("Or actuel : %d\n", c.Gold)
		fmt.Println("-------------------------------------------")
		for i, it := range items {
			fmt.Printf("%2d. %-25s %4d or  [Rareté: %d]\n", i+1, it.Name, it.Price, it.Rarity)
		}
		fmt.Println("-------------------------------------------")
		fmt.Println(" 0. Quitter")
		fmt.Print("Votre choix : ")
		var choix int
		fmt.Scan(&choix)
		if choix == 0 {
			fmt.Println("Merci de votre visite !")
			break
		}
		if choix > 0 && choix <= len(items) {
			itemChoisi := items[choix-1]
			fmt.Printf("\nVous allez acheter : %s pour %d or. Confirmer ? (o/n) : ", itemChoisi.Name, itemChoisi.Price)
			var confirm string
			fmt.Scan(&confirm)
			if confirm == "o" || confirm == "O" {
				BuyItem(c, itemChoisi)
			} else {
				fmt.Println("Achat annulé.")
			}
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}

var SpellBookFireball = Item{Name: "Livre de Sort : Boule de Feu", Quantity: 1, Rarity: 2, Price: 25}
var SpellBookIce = Item{Name: "Livre de Sort : Éclair de Glace", Quantity: 1, Rarity: 2, Price: 25}
var SpellBookHeal = Item{Name: "Livre de Sort : Soin Rapide", Quantity: 1, Rarity: 2, Price: 30}

func CanAddItem(inventaire []character.Item) bool {
	return len(inventaire) < 10
}

func AddItemToInventory(char *character.Character, newItem Item) {
	if CanAddItem(char.Inventaire) {
		// Convert item.Item to character.Item for compatibility
		converted := character.Item{Name: newItem.Name, Quantity: newItem.Quantity, Rarity: newItem.Rarity}
		char.Inventaire = append(char.Inventaire, converted)
	} else {
		fmt.Println("Inventaire plein ! Impossible d'ajouter un nouvel item.")
	}
}

type Item struct {
	Name     string
	Quantity int
	Rarity   int
	Price    int
}

// Items disponibles chez le forgeron
var BlacksmithItems = []Item{
	{Name: "Épée de chevalier", Quantity: 1, Rarity: 1, Price: 50},
	{Name: "Hache de guerre", Quantity: 1, Rarity: 1, Price: 60},
	{Name: "Lame céleste", Quantity: 1, Rarity: 2, Price: 120},
	{Name: "Baton Magique", Quantity: 1, Rarity: 1, Price: 80},
	{Name: "Griffes d'acier", Quantity: 1, Rarity: 1, Price: 70},
	{Name: "Trident infernal", Quantity: 1, Rarity: 2, Price: 150},
}

// Items disponibles chez le marchand (potions et poisons)
var ShopItems = []Item{
	{Name: "Potion de santé", Quantity: 1, Rarity: 0, Price: 10},
	{Name: "Potion de mana", Quantity: 1, Rarity: 0, Price: 12},
	{Name: "Petit poison", Quantity: 1, Rarity: 0, Price: 15},
	{Name: "Grand poison", Quantity: 1, Rarity: 1, Price: 30},
}

var Gold = Item{Name: "Or", Quantity: 1, Rarity: 0, Price: 1}
var MagicStaff = Item{Name: "Baton Magique", Quantity: 1, Rarity: 0, Price: 5}
var CelestialBlade = Item{Name: "Lame céleste", Quantity: 1, Rarity: 0, Price: 5}
var BattleAxe = Item{Name: "Hache de guerre", Quantity: 1, Rarity: 0, Price: 5}
var KnightsSword = Item{Name: "Épée de chevalier", Quantity: 1, Rarity: 0, Price: 5}
var SteelClaws = Item{Name: "Griffes d'acier", Quantity: 1, Rarity: 0, Price: 5}
var InfernalTrident = Item{Name: "Trident infernal", Quantity: 1, Rarity: 0, Price: 5}

var HealthPotion = Item{Name: "Potion de santé", Quantity: 1, Rarity: 1, Price: 10}
var BasicSword = Item{Name: "Épée basique", Quantity: 1, Rarity: 1, Price: 20}

var PoisonPotion = Item{Name: "Potion empoisonnée", Quantity: 1, Rarity: 1, Price: 15}
var Wood = Item{Name: "Bois", Quantity: 1, Rarity: 1, Price: 2}

var SmallBomb = Item{Name: "Petite bombe", Quantity: 1, Rarity: 2, Price: 50}
var Iron = Item{Name: "Fer", Quantity: 1, Rarity: 2, Price: 5}
var Cuir = Item{Name: "Cuir", Quantity: 1, Rarity: 2, Price: 5}

var RareSword = Item{Name: "Épée rare", Quantity: 1, Rarity: 3, Price: 100}
var RareArmor = Item{Name: "Armure rare", Quantity: 1, Rarity: 3, Price: 100}

var EpicSword = Item{Name: "Épée épique", Quantity: 1, Rarity: 4, Price: 200}
var EpicArmor = Item{Name: "Armure épique", Quantity: 1, Rarity: 4, Price: 200}
var BigBomb = Item{Name: "Grosse bombe", Quantity: 1, Rarity: 4, Price: 200}

var InventoryUpgrade = Item{Name: "Augmentation capacité inventaire", Quantity: 1, Rarity: 5, Price: 400}
var Uranium = Item{Name: "Uranium", Quantity: 1, Rarity: 5, Price: 400}

var LegendarySword = Item{Name: "Épée légendaire", Quantity: 1, Rarity: 6, Price: 750}
var LegendaryArmor = Item{Name: "Armure légendaire", Quantity: 1, Rarity: 6, Price: 750}
var DemonBLood = Item{Name: "Sang de démon", Quantity: 1, Rarity: 6, Price: 750}
var AngelFeather = Item{Name: "Plume d'ange", Quantity: 1, Rarity: 6, Price: 750}
var Obsidian = Item{Name: "Obsidienne", Quantity: 1, Rarity: 6, Price: 750}

var DemoniacSword = Item{Name: "Épée démoniaque", Quantity: 1, Rarity: 7, Price: 1000}
var DemoniacArmor = Item{Name: "Armure démoniaque", Quantity: 1, Rarity: 7, Price: 1000}
var AngelicSword = Item{Name: "Épée angélique", Quantity: 1, Rarity: 7, Price: 1000}
var AngelicArmor = Item{Name: "Armure angélique", Quantity: 1, Rarity: 7, Price: 1000}

var Nuke = Item{Name: "Bombe nucléaire", Quantity: 1, Rarity: 8, Price: 2500}
