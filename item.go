package main

type item struct {
	name     string
	quantity int
	rarity   int
}

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
