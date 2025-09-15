package main

import "fmt"

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

func inventairePlein(j *Character) bool {
	total := len(j.Inventaire)
	return total >= j.MaxInventaire
}

func fabriquer(j *Character, it item) {
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
	if inventairePlein(j) {
		fmt.Println("Votre inventaire est plein. Vous ne pouvez pas fabriquer cet objet.")
		return
	}
	j.Gold -= recette.CoutOr
	for ressource, qty := range recette.Ressources {
		j.Ressources[ressource] -= qty
	}
	j.Inventaire = append(j.Inventaire, it)
	fmt.Printf("Félicitations ! Vous avez fabriqué un(e) %s !\n", it.name)
}

func menuForgeron(j *Character) {
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
		fmt.Println("| 15. Quitter                       |")
		fmt.Println("=====================================")
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fabriquer(j, "Épée basique")
		case 2:
			fabriquer(j, "Épée rare")
		case 3:
			fabriquer(j, "Épée épique")
		case 4:
			fabriquer(j, "Épée légendaire")
		case 5:
			fabriquer(j, "Épée démoniaque")
		case 6:
			fabriquer(j, "Épée angélique")
		case 7:
			fabriquer(j, "Armure rare")
		case 8:
			fabriquer(j, "Armure épique")
		case 9:
			fabriquer(j, "Armure légendaire")
		case 10:
			fabriquer(j, "Armure démoniaque")
		case 11:
			fabriquer(j, "Armure angélique")
		case 12:
			fabriquer(j, "Petite bombe")
		case 13:
			fabriquer(j, "Grosse bombe")
		case 14:
			fabriquer(j, "Bombe nucléaire")
		case 15:
			fmt.Println("Merci de votre visite ! À bientôt.")
			return
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}

func bonusEquipement(e Equipment) int {
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

func retirerDeLInventaire(j *Character, it item) {
	for i, v := range j.Inventaire {
		if v.name == it.name {
			j.Inventaire = append(j.Inventaire[:i], j.Inventaire[i+1:]...)
			return
		}
	}
}

func ajouterALInventaire(j *Character, it item) {
	if len(j.Inventaire) < j.MaxInventaire {
		j.Inventaire = append(j.Inventaire, it)
	} else {
		fmt.Println("Votre inventaire est plein.")
	}
}

func equiper(j *Character, item string) {
	var emplacementPrecedent string
	var oldEquip Equipment

	switch item {
	case "Casque épique":
		emplacementPrecedent = j.Equipement.Tete
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Tete != "" {
			retirerDeLInventaire(j, j.Equipement.Tete)
		}
		j.Equipement.Tete = item

	case "Armure épique":
		emplacementPrecedent = j.Equipement.Torse
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Torse != "" {
			retirerDeLInventaire(j, j.Equipement.Torse)
		}
		j.Equipement.Torse = item

	case "Bottes épiques":
		emplacementPrecedent = j.Equipement.Pieds
		oldEquip = Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds}
		if j.Equipement.Pieds != "" {
			retirerDeLInventaire(j, j.Equipement.Pieds)
		}
		j.Equipement.Pieds = item
	}

	if emplacementPrecedent != "" {
		ajouterALInventaire(j, emplacementPrecedent)
		j.MaxHP -= bonusEquipement(oldEquip)
	}
	retirerDeLInventaire(j, item)
	j.MaxHP += bonusEquipement(Equipment{Tete: j.Equipement.Tete, Torse: j.Equipement.Torse, Pieds: j.Equipement.Pieds})

	fmt.Printf("Vous avez équipé %s !\n", item)
}

func upgradeInventorySlot(j *Character) {
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

func initGoblin() Monster {
	return Monster{
		Nom:      "Gobelin d'entrainement",
		MaxHP:    40,
		HPactuel: 40,
		Attaque:  5,
	}
}

func goblinPattern(j *Character) {
	goblin := initGoblin()
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

func utiliserObjet(j *Character, m *Monster, objet string) {
	switch objet {
	case "Potion de soin":
		soin := 20
		j.HPactuel += soin
		if j.HPactuel > j.MaxHP {
			j.HPactuel = j.MaxHP
		}
		fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", objet, soin)
		retirerDeLInventaire(j, objet)
		fmt.Printf("%s : %d/%d PV\n", j.Nom, j.HPactuel, j.MaxHP)
	default:
		fmt.Printf("L'objet %s n'a pas d'effet utilisable en combat.\n", objet)
	}
}

func characterTurn(j *Character, m *Monster) {
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
			// Attaque basique
			degats := 5
			m.HPactuel -= degats
			if m.HPactuel < 0 {
				m.HPactuel = 0
			}
			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s\n", j.Nom, degats, m.Nom)
			fmt.Printf("%s : %d/%d PV\n", m.Nom, m.HPactuel, m.MaxHP)
			return // Fin du tour du joueur
		case 2:
			// Inventaire
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
				utiliserObjet(j, m, j.Inventaire[choixObj-1])
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
