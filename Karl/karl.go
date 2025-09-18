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
	return total >= 10 
}
func Fabriquer(j *character.Character, it item.Item) {
	
	if InventairePlein(j) {
		fmt.Println("Votre inventaire est plein. Vous ne pouvez pas fabriquer cet objet.")
		return
	}
	if j.Gold < it.Price {
		fmt.Println("Vous n'avez pas assez d'or pour fabriquer cet objet.")
		return
	}
	j.Gold -= it.Price
	
	newItem := character.Item{Name: it.Name, Quantity: it.Quantity, Rarity: it.Rarity}
	j.Inventaire = append(j.Inventaire, newItem)
	fmt.Printf("Félicitations ! Vous avez fabriqué un(e) %s !\n", it.Name)
}
func MenuForgeron(j *character.Character) {
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
		fmt.Println("| 21. Quittez le forgeron           |")
		fmt.Println("=====================================")
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)
		if choix == 21 {
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

func RetirerDeLInventaire(j *character.Character, it character.Item) {
	for i, v := range j.Inventaire {
		if v.Name == it.Name {
			j.Inventaire = append(j.Inventaire[:i], j.Inventaire[i+1:]...)
			return
		}
	}
}

func AjouterALInventaire(j *character.Character, it item.Item) {
	if len(j.Inventaire) < 10 {
		newItem := character.Item{Name: it.Name, Quantity: it.Quantity, Rarity: it.Rarity}
		j.Inventaire = append(j.Inventaire, newItem)
	} else {
		fmt.Println("Votre inventaire est plein.")
	}
}


type Monster struct {
	Nom      string
	MaxHP    int
	HPactuel int
	Attaque  int
}

func InitGoblin() Monster {
	return Monster{
		Nom:      "Gobelins d'entrainement",
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

func UtiliserObjet(j *character.Character, m *Monster, objet character.Item) {
	switch objet.Name {
	case "Potion de soin":
		soin := 20
		j.HPactuel += soin
		if j.HPactuel > j.MaxHP {
			j.HPactuel = j.MaxHP
		}
		fmt.Printf("Vous utilisez %s et récupérez %d PV !\n", objet.Name, soin)
		RetirerDeLInventaire(j, objet)
		fmt.Printf("%s : %d/%d PV\n", j.Nom, j.HPactuel, j.MaxHP)
	default:
		fmt.Printf("L'objet %s n'a pas d'effet utilisable en combat.\n", objet.Name)
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
