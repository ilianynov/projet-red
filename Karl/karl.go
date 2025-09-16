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

func InventairePlein(j *character.Character) bool {
	total := len(j.Inventaire)
	return total >= j.MaxInventaire
}

func Fabriquer(j *character.Character, it item.Item) {
	recette, exists := recettes[it.Name]
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
	if InventairePlein(j) {
		fmt.Println("Votre inventaire est plein. Vous ne pouvez pas fabriquer cet objet.")
		return
	}
	j.Gold -= recette.CoutOr
	for ressource, qty := range recette.Ressources {
		j.Ressources[ressource] -= qty
	}
	j.Inventaire = append(j.Inventaire, it)
	fmt.Printf("Félicitations ! Vous avez fabriqué un(e) %s !\n", it.Name)
}

func MenuForgeron(j *character.Character) {
	var choix int
	for {
		fmt.Println("=====================================")
		fmt.Println("|        Bienvenue chez le          |")
		fmt.Println("|           Forgeron !              |")
		fmt.Println("=====================================")
		fmt.Println("| 1. Épée basique (10 Or)           |")
		fmt.Println("| 2. Épée rare (20 Or)              |")
		fmt.Println("| 3. Épée épique (25 Or)            |")
		fmt.Println("| 4. Épée légendaire (40 Or)        |")
		fmt.Println("| 5. Épée démoniaque (60 Or)        |")
		fmt.Println("| 6. Épée angélique (60 Or)         |")
		fmt.Println("| 7. Armure rare (30 Or)            |")
		fmt.Println("| 8. Armure épique (50 Or)          |")
		fmt.Println("| 9. Armure légendaire (70 Or)      |")
		fmt.Println("| 10. Armure démoniaque (90 Or)     |")
		fmt.Println("| 11. Armure angélique (90 Or)      |")
		fmt.Println("| 12. Petite bombe (30 Or)          |")
		fmt.Println("| 13. Grosse bombe (50 Or)          |")
		fmt.Println("| 14. Bombe nucléaire (100 Or)      |")
		fmt.Println("| 15. Baton magique (75 Or)         |")
		fmt.Println("| 16. Lame céleste (100 Or)         |")
		fmt.Println("| 17. Hache de guerre (75 Or)       |")
		fmt.Println("| 18. Épée de chevalier (80 Or)     |")
		fmt.Println("| 19. Griffes d'acier (70 Or)       |")
		fmt.Println("| 20. Trident infernal (90 Or)      |")
		fmt.Println("| 21. Quitter le forgeron           |")
		fmt.Println("=====================================")
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			Fabriquer(j, item.BasicSword)
		case 2:
			Fabriquer(j, item.RareSword)
		case 3:
			Fabriquer(j, item.EpicSword)
		case 4:
			Fabriquer(j, item.LegendarySword)
		case 5:
			Fabriquer(j, item.DemoniacSword)
		case 6:
			Fabriquer(j, item.AngelicSword)
		case 7:
			Fabriquer(j, item.RareArmor)
		case 8:
			Fabriquer(j, item.EpicArmor)
		case 9:
			Fabriquer(j, item.LegendaryArmor)
		case 10:
			Fabriquer(j, item.DemoniacArmor)
		case 11:
			Fabriquer(j, item.AngelicArmor)
		case 12:
			Fabriquer(j, item.SmallBomb)
		case 13:
			Fabriquer(j, item.BigBomb)
		case 14:
			Fabriquer(j, item.Nuke)
		case 15:
			Fabriquer(j, item.MagicStaff)
		case 16:
			Fabriquer(j, item.CelestialBlade)
		case 17:
			Fabriquer(j, item.BattleAxe)
		case 18:
			Fabriquer(j, item.KnightsSword)
		case 19:
			Fabriquer(j, item.SteelClaws)
		case 20:
			Fabriquer(j, item.InfernalTrident)
		case 21:
			fmt.Println("Merci de votre visite ! À bientôt.")
			return
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}

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

func RetirerDeLInventaire(j *character.Character, it item.Item) {
	for i, v := range j.Inventaire {
		if v.Name == it.Name {
			j.Inventaire = append(j.Inventaire[:i], j.Inventaire[i+1:]...)
			return
		}
	}
}

func AjouterALInventaire(j *character.Character, it item.Item) {
	if len(j.Inventaire) < j.MaxInventaire {
		j.Inventaire = append(j.Inventaire, it)
	} else {
		fmt.Println("Votre inventaire est plein.")
	}
}

func Equiper(j *character.Character, item string) {
	var emplacementPrecedent string
	var oldEquip Equipment

	switch item {
	case "Casque épique":
		emplacementPrecedent = j.Equipement.Tete
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Tete != "" {
			// RetirerDeLInventaire(j, j.Equipement.Tete) // Needs item.Item, not string
		}
		j.Equipement.Tete = item

	case "Armure épique":
		emplacementPrecedent = j.Equipement.Torse
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Torse != "" {
			// RetirerDeLInventaire(j, j.Equipement.Torse) // Needs item.Item, not string
		}
		j.Equipement.Torse = item

	case "Bottes épiques":
		emplacementPrecedent = j.Equipement.Pieds
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Pieds != "" {
			// RetirerDeLInventaire(j, j.Equipement.Pieds) // Needs item.Item, not string
		}
		j.Equipement.Pieds = item
	}

	if emplacementPrecedent != "" {
		// AjouterALInventaire(j, emplacementPrecedent) // Needs item.Item, not string
		j.MaxHP -= BonusEquipement(oldEquip)
	}
	// RetirerDeLInventaire(j, item) // Needs item.Item, not string
	j.MaxHP += BonusEquipement(Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds})

	fmt.Printf("Vous avez équipé %s !\n", item)
}

func UpgradeInventorySlot(j *character.Character) {
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

func UtiliserObjet(j *character.Character, m *Monster, objet string) {
	switch objet {
	case "Potion de soin":
		soin := 20
		j.HPactuel += soin
		if j.HPactuel > j.MaxHP {
			j.HPactuel = j.MaxHP
		}
		fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", objet, soin)
		RetirerDeLInventaire(j, item.HealthPotion)
		fmt.Printf("%s : %d/%d PV\n", j.Nom, j.HPactuel, j.MaxHP)
	default:
		fmt.Printf("L'objet %s n'a pas d'effet utilisable en combat.\n", objet)
	}
}

func CharacterTurn(j *character.Character, m *Monster) {
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
				fmt.Printf("%d. %s\n", i+1, obj.Name)
			}
			fmt.Print("Choisissez un objet à utiliser (0 pour annuler) : ")
			var choixObj int
			fmt.Scan(&choixObj)
			if choixObj == 0 {
				continue
			}
			if choixObj > 0 && choixObj <= len(j.Inventaire) {
				UtiliserObjet(j, m, j.Inventaire[choixObj-1].Name)
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
