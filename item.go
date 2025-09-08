package main

type item struct {
	name     string
	quantity int
	rarity   int
	price    int
}

var Gold = item{name: "Or", quantity: 1, rarity: 0, price: 0}
var BasicSword = item{name: "Épée basique", quantity: 1, rarity: 1}
var HealthPotion = item{name: "Potion de santé", quantity: 1, rarity: 1}
var PoisonPotion = item{name: "Potion empoisonnée", quantity: 1, rarity: 1}
var InventoryUpgrade = item{name: "Augmentation capacité inventaire", quantity: 1, rarity: 5}
var RareSword = item{name: "Épée rare", quantity: 1, rarity: 3}
var RareArmor = item{name: "Armure rare", quantity: 1, rarity: 3}
var EpicSword = item{name: "Épée épique", quantity: 1, rarity: 4}
var EpicArmor = item{name: "Armure épique", quantity: 1, rarity: 4}
var LegendarySword = item{name: "Épée légendaire", quantity: 1, rarity: 5}
var LegendaryArmor = item{name: "Armure légendaire", quantity: 1, rarity: 5}
var DemoniacSword = item{name: "Épée démoniaque", quantity: 1, rarity: 7}
var DemoniacArmor = item{name: "Armure démoniaque", quantity: 1, rarity: 7}
var AngelicSword = item{name: "Épée angélique", quantity: 1, rarity: 7}
var AngelicArmor = item{name: "Armure angélique", quantity: 1, rarity: 7}
var SmallBomb = item{name: "Petite bombe", quantity: 1, rarity: 2}
var BigBomb = item{name: "Grosse bombe", quantity: 1, rarity: 4}
var Nuke = item{name: "Bombe nucléaire", quantity: 1, rarity: 8}
