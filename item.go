package main

import "fmt"

func BuyItem(c *Character, it item) {
	if len(c.Inventaire) >= 10 {
		fmt.Println("Votre inventaire est plein.")
		return
	}
	if c.Gold >= it.price {
		c.Gold -= it.price
		c.Inventaire = append(c.Inventaire, it)
		fmt.Printf("Vous avez acheté %s pour %d pièces d'or.\n", it.name, it.price)
	} else {
		fmt.Println("Pas assez d'or pour acheter cet objet.")
	}
}

func ShopMenu(c *Character) {

	allItems := []item{
		BasicSword, HealthPotion, PoisonPotion, SmallBomb, BigBomb, Nuke,
		RareSword, RareArmor, EpicSword, EpicArmor, InventoryUpgrade,
		Uranium, BattleAxe, CelestialBlade, KnightsSword, SteelClaws, InfernalTrident,
	}

	var items []item
	for _, it := range allItems {
		if it.rarity >= 3 && it.rarity <= 5 {
			items = append(items, it)
		}
	}
	for {
		fmt.Println("\n================ BOUTIQUE =================")
		fmt.Printf("Or actuel : %d\n", c.Gold)
		fmt.Println("-------------------------------------------")
		for i, it := range items {
			fmt.Printf("%2d. %-25s %4d or  [Rareté: %d]\n", i+1, it.name, it.price, it.rarity)
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
			fmt.Printf("\nVous allez acheter : %s pour %d or. Confirmer ? (o/n) : ", itemChoisi.name, itemChoisi.price)
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

var SpellBookFireball = item{name: "Livre de Sort : Boule de Feu", quantity: 1, rarity: 2, price: 25}
var SpellBookIce = item{name: "Livre de Sort : Éclair de Glace", quantity: 1, rarity: 2, price: 25}
var SpellBookHeal = item{name: "Livre de Sort : Soin Rapide", quantity: 1, rarity: 2, price: 30}

func CanAddItem(inventaire []item) bool {
	return len(inventaire) < 10
}

func AddItemToInventory(char *Character, newItem item) {
	if CanAddItem(char.Inventaire) {
		char.Inventaire = append(char.Inventaire, newItem)
	} else {
		fmt.Println("Inventaire plein ! Impossible d'ajouter un nouvel item.")
	}
}

type item struct {
	name     string
	quantity int
	rarity   int
	price    int
}

var Gold = item{name: "Or", quantity: 1, rarity: 0, price: 1}
var MagicStaff = item{name: "Baton Magique", quantity: 1, rarity: 0, price: 5}
var CelestialBlade = item{name: "Lame céleste", quantity: 1, rarity: 0, price: 5}
var BattleAxe = item{name: "Hache de guerre", quantity: 1, rarity: 0, price: 5}
var KnightsSword = item{name: "Épée de chevalier", quantity: 1, rarity: 0, price: 5}
var SteelClaws = item{name: "Griffes d'acier", quantity: 1, rarity: 0, price: 5}
var InfernalTrident = item{name: "Trident infernal", quantity: 1, rarity: 0, price: 5}

var HealthPotion = item{name: "Potion de santé", quantity: 1, rarity: 1, price: 10}
var BasicSword = item{name: "Épée basique", quantity: 1, rarity: 1, price: 20}

var PoisonPotion = item{name: "Potion empoisonnée", quantity: 1, rarity: 1, price: 15}
var Wood = item{name: "Bois", quantity: 1, rarity: 1, price: 2}

var SmallBomb = item{name: "Petite bombe", quantity: 1, rarity: 2, price: 50}
var Iron = item{name: "Fer", quantity: 1, rarity: 2, price: 5}
var Cuir = item{name: "Cuir", quantity: 1, rarity: 2, price: 5}

var RareSword = item{name: "Épée rare", quantity: 1, rarity: 3, price: 100}
var RareArmor = item{name: "Armure rare", quantity: 1, rarity: 3, price: 100}

var EpicSword = item{name: "Épée épique", quantity: 1, rarity: 4, price: 200}
var EpicArmor = item{name: "Armure épique", quantity: 1, rarity: 4, price: 200}
var BigBomb = item{name: "Grosse bombe", quantity: 1, rarity: 4, price: 200}

var InventoryUpgrade = item{name: "Augmentation capacité inventaire", quantity: 1, rarity: 5, price: 400}
var Uranium = item{name: "Uranium", quantity: 1, rarity: 5, price: 400}

var LegendarySword = item{name: "Épée légendaire", quantity: 1, rarity: 6, price: 750}
var LegendaryArmor = item{name: "Armure légendaire", quantity: 1, rarity: 6, price: 750}
var DemonBLood = item{name: "Sang de démon", quantity: 1, rarity: 6, price: 750}
var AngelFeather = item{name: "Plume d'ange", quantity: 1, rarity: 6, price: 750}
var Obsidian = item{name: "Obsidienne", quantity: 1, rarity: 6, price: 750}

var DemoniacSword = item{name: "Épée démoniaque", quantity: 1, rarity: 7, price: 1000}
var DemoniacArmor = item{name: "Armure démoniaque", quantity: 1, rarity: 7, price: 1000}
var AngelicSword = item{name: "Épée angélique", quantity: 1, rarity: 7, price: 1000}
var AngelicArmor = item{name: "Armure angélique", quantity: 1, rarity: 7, price: 1000}

var Nuke = item{name: "Bombe nucléaire", quantity: 1, rarity: 8, price: 2500}
