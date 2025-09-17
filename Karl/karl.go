package karl

import (
	"fmt"
	"projet_red/character"
	"projet_red/item"
)

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

<<<<<<< Updated upstream
func InventairePlein(j *character.Character) bool {
=======
func InventairePlein(j *Character) bool {
>>>>>>> Stashed changes
	total := len(j.Inventaire)
	return total >= 10 // Assuming max inventory size is 10
}
<<<<<<< Updated upstream
func Fabriquer(j *character.Character, it item.Item) {
	// For simplicity, just check gold and inventory size
=======

func Fabriquer(j *Character, it item) {
	recette, exists := recettes[it.name]
	if !exists {
		fmt.Println("Cet objet est indisponible chez le forgeron.")
		return
	}
	if j.Gold < recette.CoutOr {
		fmt.Println("Vous n'avez pas assez d'or pour fabriquer cet objet.")
		return
	}
	for ressource, qty := range recette.Ressources {
		if j.Ressources[ressource] < qty {
			fmt.Printf("Vous n'avez pas assez de %s pour fabriquer cet objet.\n", ressource)
			return
		}
	}
>>>>>>> Stashed changes
	if InventairePlein(j) {
		fmt.Println("Votre inventaire est plein. Vous ne pouvez pas fabriquer cet objet.")
		return
	}
	if j.Gold < it.Price {
		fmt.Println("Vous n'avez pas assez d'or pour fabriquer cet objet.")
		return
	}
	j.Gold -= it.Price
	// Convert item.Item to character.Item
	newItem := character.Item{Name: it.Name, Quantity: it.Quantity, Rarity: it.Rarity}
	j.Inventaire = append(j.Inventaire, newItem)
	fmt.Printf("Félicitations ! Vous avez fabriqué un(e) %s !\n", it.Name)
}
<<<<<<< Updated upstream
func MenuForgeron(j *character.Character) {
=======

func MenuForgeron(j *Character) {
>>>>>>> Stashed changes
	var choix int
	items := []item.Item{
		item.BasicSword, item.RareSword, item.EpicSword, item.LegendarySword, item.DemoniacSword, item.AngelicSword,
		item.RareArmor, item.EpicArmor, item.LegendaryArmor, item.DemoniacArmor, item.AngelicArmor,
		item.SmallBomb, item.BigBomb, item.Nuke, item.MagicStaff, item.CelestialBlade, item.BattleAxe, item.KnightsSword, item.SteelClaws, item.InfernalTrident,
	}
	for {
		fmt.Println("=====================================")
		fmt.Println("|        Bienvenue chez le          |")
		fmt.Println("|           Forgeron !              |")
		fmt.Println("=====================================")
		for i, it := range items {

			fmt.Printf("| %2d. %-20s (%3d Or) |\n", i+1, it.Name, it.Price)
		}
		fmt.Println("| 21. Quitter le forgeron           |")
		fmt.Println("=====================================")
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)
<<<<<<< Updated upstream
		if choix == 21 {
=======

		switch choix {
		case 1:
			Fabriquer(j, "Épée basique")
		case 2:
			Fabriquer(j, "Épée rare")
		case 3:
			Fabriquer(j, "Épée épique")
		case 4:
			Fabriquer(j, "Épée légendaire")
		case 5:
			Fabriquer(j, "Épée démoniaque")
		case 6:
			Fabriquer(j, "Épée angélique")
		case 7:
			Fabriquer(j, "Armure rare")
		case 8:
			Fabriquer(j, "Armure épique")
		case 9:
			Fabriquer(j, "Armure légendaire")
		case 10:
			Fabriquer(j, "Armure démoniaque")
		case 11:
			Fabriquer(j, "Armure angélique")
		case 12:
			Fabriquer(j, "Petite bombe")
		case 13:
			Fabriquer(j, "Grosse bombe")
		case 14:
			Fabriquer(j, "Bombe nucléaire")
		case 15:
			Fabriquer(j, "Baton magique")
		case 16:
			Fabriquer(j, "Lame céleste")
		case 17:
			Fabriquer(j, "Hache de guerre")
		case 18:
			Fabriquer(j, "Épée de chevalier")
		case 19:
			Fabriquer(j, "Griffes d'acier")
		case 20:
			Fabriquer(j, "Trident infernal")
		case 21:
>>>>>>> Stashed changes
			fmt.Println("Merci de votre visite ! À bientôt.")
			return
		}
		if choix > 0 && choix <= len(items) {
			Fabriquer(j, items[choix-1])
		} else {
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
func BonusEquipement(e Equipment) int {
	bonus := 0
	if e.Tete == "Casque épique" {
		bonus += 10
	}
	if e.Torse == "Armure épique" {
		bonus += 25
	}
	if e.Pieds == "Bottes épiques" {
		bonus += 15
	}
	return bonus
}

<<<<<<< Updated upstream
func RetirerDeLInventaire(j *character.Character, it character.Item) {
=======
func RetirerDeLInventaire(j *Character, it item) {
>>>>>>> Stashed changes
	for i, v := range j.Inventaire {
		if v.Name == it.Name {
			j.Inventaire = append(j.Inventaire[:i], j.Inventaire[i+1:]...)
			return
		}
	}
}

<<<<<<< Updated upstream
func AjouterALInventaire(j *character.Character, it item.Item) {
	if len(j.Inventaire) < 10 {
		newItem := character.Item{Name: it.Name, Quantity: it.Quantity, Rarity: it.Rarity}
		j.Inventaire = append(j.Inventaire, newItem)
=======
func AjouterALInventaire(j *Character, it item) {
	if len(j.Inventaire) < j.MaxInventaire {
		j.Inventaire = append(j.Inventaire, it)
>>>>>>> Stashed changes
	} else {
		fmt.Println("Votre inventaire est plein.")
	}
}

<<<<<<< Updated upstream
// func UpgradeInventorySlot(j *character.Character) {
//     // Inventory upgrade logic removed: NbAugmentationsInventaire not defined
// }
=======
func Equiper(j *Character, item string) {
	var emplacementPrecedent string
	var oldEquip Equipment

	switch item {
	case "Casque épique":
		emplacementPrecedent = j.Equipement.Tete
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Tete != "" {
			RetirerDeLInventaire(j, j.Equipement.Tete)
		}
		j.Equipement.Tete = item

	case "Armure épique":
		emplacementPrecedent = j.Equipement.Torse
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Torse != "" {
			RetirerDeLInventaire(j, j.Equipement.Torse)
		}
		j.Equipement.Torse = item

	case "Bottes épiques":
		emplacementPrecedent = j.Equipement.Pieds
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Pieds != "" {
			RetirerDeLInventaire(j, j.Equipement.Pieds)
		}
		j.Equipement.Pieds = item
	}

	if emplacementPrecedent != "" {
		AjouterALInventaire(j, emplacementPrecedent)
		j.MaxHP -= BonusEquipement(oldEquip)
	}
	RetirerDeLInventaire(j, item)
	j.MaxHP += BonusEquipement(Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds})

	fmt.Printf("Vous avez équipé %s !\n", item)
}

func UpgradeInventorySlot(j *Character) {
	if j.NbAugmentationsInventaire >= 3 {
		fmt.Println("Vous avez déjà atteint le nombre maximal d'augmentations d'inventaire (3).")
		return
	}
	if j.Or < 30 {
		fmt.Println("Vous n'avez pas assez d'or pour acheter une augmentation d'inventaire.")
		return
	}
	j.Or -= 30
	j.MaxInventaire += 10
	j.NbAugmentationsInventaire++
	fmt.Printf("Votre inventaire a été augmenté ! Capacité maximale : %d\n", j.MaxInventaire)
}
>>>>>>> Stashed changes

type Monster struct {
	Nom      string
	MaxHP    int
	HPactuel int
	Attaque  int
}

func InitGoblin() Monster {
	return Monster{
		Nom:      "Gobelin d'entrainement",
		MaxHP:    40,
		HPactuel: 40,
		Attaque:  5,
	}
}
func InitOgre() Monster {
	return Monster{
		Nom:      "Ogre",
		MaxHP:    60,
		HPactuel: 60,
		Attaque:  10,
	}
}
func InitLoupGarou() Monster {
	return Monster{
		Nom:      "Loup-Garou",
		MaxHP:    80,
		HPactuel: 80,
		Attaque:  15,
	}
}
func InitDragon() Monster {
	return Monster{
		Nom:      "Dragon",
		MaxHP:    150,
		HPactuel: 150,
		Attaque:  25,
	}
}

<<<<<<< Updated upstream
// GetMonsterIntro returns a fitting intro text for each monster
func GetMonsterIntro(monsterName string) string {
	switch monsterName {
	case "Gobelin d'entrainement":
		return "Le gobelin ricane et vous provoque !"
	case "Ogre":
		return "L'ogre grogne et tape du pied, prêt à vous écraser !"
	case "Loup-Garou":
		return "Le loup-garou pousse un hurlement terrifiant !"
	case "Dragon":
		return "Le dragon rugit, les flammes dansent dans sa gueule !"
	default:
		return "Votre adversaire vous fixe avec intensité..."
	}
}

func GoblinPattern(j *character.Character) {
=======
func GoblinPattern(j *Character) {
>>>>>>> Stashed changes
	goblin := InitGoblin()
	tour := 1

	for j.HPactuel > 0 && goblin.HPactuel > 0 {
		var degats int
		if tour%3 == 0 {
			degats = goblin.Attaque * 2
		} else {
			degats = goblin.Attaque
		}

		j.HPactuel -= degats
		if j.HPactuel < 0 {
			j.HPactuel = 0
		}

		fmt.Printf("%s inflige à %s %d de dégats\n", goblin.Nom, j.Nom, degats)
		fmt.Printf("%s : %d/%d PV\n", j.Nom, j.HPactuel, j.MaxHP)

		if j.HPactuel == 0 {
			fmt.Printf("%s est vaincu !\n", j.Nom)
			break
		}
		tour++
	}
}

<<<<<<< Updated upstream
func UtiliserObjet(j *character.Character, m *Monster, objet character.Item) {
	switch objet.Name {
=======
func UtiliserObjet(j *Character, m *Monster, objet string) {
	switch objet {
>>>>>>> Stashed changes
	case "Potion de soin":
		soin := 20
		j.HPactuel += soin
		if j.HPactuel > j.MaxHP {
			j.HPactuel = j.MaxHP
		}
<<<<<<< Updated upstream
		fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", objet.Name, soin)
=======
		fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", objet, soin)
>>>>>>> Stashed changes
		RetirerDeLInventaire(j, objet)
		fmt.Printf("%s : %d/%d PV\n", j.Nom, j.HPactuel, j.MaxHP)
	default:
		fmt.Printf("L'objet %s n'a pas d'effet utilisable en combat.\n", objet.Name)
	}
}

<<<<<<< Updated upstream
func CharacterTurn(j *character.Character, m *Monster) {
=======
func CharTurn(j *Character, m *Monster) {
>>>>>>> Stashed changes
	var choix int
	for {
		fmt.Println("=====================================")
		fmt.Println("|         Combat contre :", m.Nom)
		fmt.Println("=====================================")
		fmt.Println("| 1. Attaquer")
		fmt.Println("| 2. Inventaire")
		fmt.Println("=====================================")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			degats := 5
			m.HPactuel -= degats
			if m.HPactuel < 0 {
				m.HPactuel = 0
			}
			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s\n", j.Nom, degats, m.Nom)
			fmt.Printf("%s : %d/%d PV\n", m.Nom, m.HPactuel, m.MaxHP)
			return // Fin du tour du joueur
		case 2:
			if len(j.Inventaire) == 0 {
				fmt.Println("Votre inventaire est vide.")
				continue
			}
			fmt.Println("Inventaire :")
			for i, obj := range j.Inventaire {
				fmt.Printf("%d. %s\n", i+1, obj)
			}
			fmt.Print("Choisissez un objet à utiliser (0 pour annuler) : ")
			var choixObj int
			fmt.Scan(&choixObj)
			if choixObj == 0 {
				continue
			}
			if choixObj > 0 && choixObj <= len(j.Inventaire) {
				UtiliserObjet(j, m, j.Inventaire[choixObj-1])
				return // Fin du tour du joueur
			} else {
				fmt.Println("Choix invalide.")
				continue
			}
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
