var SpellBookFireball = item{name: "Livre de Sort : Boule de Feu", quantity: 1, rarity: 2, price: 25}
var SpellBookIce = item{name: "Livre de Sort : Éclair de Glace", quantity: 1, rarity: 2, price: 25}
var SpellBookHeal = item{name: "Livre de Sort : Soin Rapide", quantity: 1, rarity: 2, price: 30}
var SpellBookFireball = item{name: "Livre de Sort : Boule de Feu", quantity: 1, rarity: 2, price: 25}

package main
var HealthPotion = item{name: "Potion de santé", quantity: 1, rarity: 1, price: 8}
var SpellBookFireball = item{name: "Livre de Sort : Boule de Feu", quantity: 1, rarity: 2, price: 25}



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

<<<<<<< Updated upstream
// Rareté 0
var Gold = item{name: "Or", quantity: 1, rarity: 0, price: 1}
var MagicStaff = item{name: "Baton Magique", quantity: 1, rarity: 0, price: 5}
var CelestialBlade = item{name: "Lame céleste", quantity: 1, rarity: 0, price: 5}
var BattleAxe = item{name: "Hache de guerre", quantity: 1, rarity: 0, price: 5}
var KnightsSword = item{name: "Épée de chevalier", quantity: 1, rarity: 0, price: 5}
var SteelClaws = item{name: "Griffes d'acier", quantity: 1, rarity: 0, price: 5}
var InfernalTrident = item{name: "Trident infernal", quantity: 1, rarity: 0, price: 5}

// Rareté 1
var BasicSword = item{name: "Épée basique", quantity: 1, rarity: 1, price: 20}
var HealthPotion = item{name: "Potion de santé", quantity: 1, rarity: 1, price: 10}
var PoisonPotion = item{name: "Potion empoisonnée", quantity: 1, rarity: 1, price: 15}
var Wood = item{name: "Bois", quantity: 1, rarity: 1, price: 2}

// Rareté 2
var SmallBomb = item{name: "Petite bombe", quantity: 1, rarity: 2, price: 50}
var Iron = item{name: "Fer", quantity: 1, rarity: 2, price: 5}
var Cuir = item{name: "Cuir", quantity: 1, rarity: 2, price: 5}

// Rareté 3
var RareSword = item{name: "Épée rare", quantity: 1, rarity: 3, price: 100}
var RareArmor = item{name: "Armure rare", quantity: 1, rarity: 3, price: 100}

// Rareté 4
var EpicSword = item{name: "Épée épique", quantity: 1, rarity: 4, price: 200}
var EpicArmor = item{name: "Armure épique", quantity: 1, rarity: 4, price: 200}
var BigBomb = item{name: "Grosse bombe", quantity: 1, rarity: 4, price: 200}

// Rareté 5
var InventoryUpgrade = item{name: "Augmentation capacité inventaire", quantity: 1, rarity: 5, price: 400}
var Uranium = item{name: "Uranium", quantity: 1, rarity: 5, price: 400}

// Rareté 6
var LegendarySword = item{name: "Épée légendaire", quantity: 1, rarity: 6, price: 750}
var LegendaryArmor = item{name: "Armure légendaire", quantity: 1, rarity: 6, price: 750}
var DemonBLood = item{name: "Sang de démon", quantity: 1, rarity: 6, price: 750}
var AngelFeather = item{name: "Plume d'ange", quantity: 1, rarity: 6, price: 750}
var Obsidian = item{name: "Obsidienne", quantity: 1, rarity: 6, price: 750}

// Rareté 7
var DemoniacSword = item{name: "Épée démoniaque", quantity: 1, rarity: 7, price: 1000}
var DemoniacArmor = item{name: "Armure démoniaque", quantity: 1, rarity: 7, price: 1000}
var AngelicSword = item{name: "Épée angélique", quantity: 1, rarity: 7, price: 1000}
var AngelicArmor = item{name: "Armure angélique", quantity: 1, rarity: 7, price: 1000}

// Rareté 8
var Nuke = item{name: "Bombe nucléaire", quantity: 1, rarity: 8, price: 2500}
=======
var Gold = item{name: "Or", quantity: 1, rarity: 0}
var BasicSword = item{name: "Épée basique", quantity: 1, rarity: 1}
var HealthPotion = item{name: "Potion de santé", quantity: 1, rarity: 1}
var PoisonPotion = item{name: "Potion empoisonnée", quantity: 1, rarity: 1}
var InventoryUpgrade = item{name: "Augmentation capacité inventaire", quantity: 1, rarity: 5}
var RareSword = item{name: "Épée rare", quantity: 1, rarity: 3}
var RareArmor = item{name: "Armure rare", quantity: 1, rarity: 3}
var EpicSword = item{name: "Épée épique", quantity: 1, rarity: 4}
var EpicArmor = item{name: "Armure épique", quantity: 1, rarity: 4}
var LegendarySword = item{name: "Épée légendaire", quantity: 1, rarity: 6}
var LegendaryArmor = item{name: "Armure légendaire", quantity: 1, rarity: 6}
var DemoniacSword = item{name: "Épée démoniaque", quantity: 1, rarity: 7}
var DemoniacArmor = item{name: "Armure démoniaque", quantity: 1, rarity: 7}
var AngelicSword = item{name: "Épée angélique", quantity: 1, rarity: 7}
var AngelicArmor = item{name: "Armure angélique", quantity: 1, rarity: 7}
var SmallBomb = item{name: "Petite bombe", quantity: 1, rarity: 2}
var BigBomb = item{name: "Grosse bombe", quantity: 1, rarity: 4}
var Nuke = item{name: "Bombe nucléaire", quantity: 1, rarity: 8}
var Wood = item{name: "Bois", quantity: 1, rarity: 1}
var Iron = item{name: "Fer", quantity: 1, rarity: 2}
var Cuir = item{name: "Cuir", quantity: 1, rarity: 2}
var Uranium = item{name: "Uranium", quantity: 1, rarity: 5}
var DemonBLood = item{name: "Sang de démon", quantity: 1, rarity: 6}
var AngelFeather = item{name: "Plume d'ange", quantity: 1, rarity: 6}
var Obsidian = item{name: "Obsidienne", quantity: 1, rarity: 6}
var MagicStaff = item{name:"Baton Magique" , quantity: 1, rarity: 0}
var CelestialBlade = item{name: "Lame céleste" , quantity: 1, rarity: 0}
var BattleAxe = item{name: "Hache de guerre" , quantity: 1, rarity 0}
var KnightsSword = item{name: "Épée de chevalier" , quantity: 1, rarity 0}
var SteelClaws = item{name: "Griffes d'acier" , quantity: 1, rarity 0}
var InfernalTrident = item{name: "Trident infernal" , quantity 1, rarity 0}









