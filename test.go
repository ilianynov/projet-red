package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// --- Structures principales ---

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Exp        int
	MaxHP      int
	HP         int
	MaxMana    int
	Mana       int
	Inventaire []string
	Or         int
	Equipement Equipment
	CapInv     int
}

type Equipment struct {
	Nom   string
	Bonus int // Bonus PV max
}

type Monster struct {
	Nom     string
	HP      int
	Att     int
	Pattern []int // Pattern d'attaque
}

// --- Fonctions utilitaires ---

func input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func randomClasse() string {
	classes := []string{"Guerrier", "Mage", "Voleur", "Archer", "Paladin", "Nain", "Ange", "Démon", "Hybride", "Loup-garou"}
	rand.Seed(time.Now().UnixNano())
	return classes[rand.Intn(len(classes))]
}

// --- Création du personnage ---

func initCharacter() Character {
	nom := input("Entrez le nom de votre personnage : ")
	classe := randomClasse()
	fmt.Println("Votre classe aléatoire :", classe)
	maxHP := 100
	maxMana := 50
	if classe == "Guerrier" {
		maxHP = 120
	} else if classe == "Mage" {
		maxMana = 100
	}
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     1,
		Exp:        0,
		MaxHP:      maxHP,
		HP:         maxHP,
		MaxMana:    maxMana,
		Mana:       maxMana,
		Inventaire: []string{"Potion de soin", "Potion de mana"},
		Or:         50,
		Equipement: Equipment{"Aucun", 0},
		CapInv:     5,
	}
}

// --- Affichage des infos ---

func displayInfo(c Character) {
	fmt.Println("----- Infos du personnage -----")
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau, "Exp :", c.Exp)
	fmt.Println("PV :", c.HP, "/", c.MaxHP)
	fmt.Println("Mana :", c.Mana, "/", c.MaxMana)
	fmt.Println("Or :", c.Or)
	fmt.Println("Equipement :", c.Equipement.Nom, "(Bonus PV :", c.Equipement.Bonus, ")")
	fmt.Println("Inventaire :", c.Inventaire, "(Capacité :", len(c.Inventaire), "/", c.CapInv, ")")
	fmt.Println("-------------------------------")
}

// --- Gestion de l'inventaire ---

func accessInventory(c *Character) {
	fmt.Println("Inventaire :", c.Inventaire)
	fmt.Println("1. Utiliser une potion de soin")
	fmt.Println("2. Utiliser une potion de mana")
	fmt.Println("3. Retirer un objet")
	fmt.Println("4. Retour")
	choix := input("Choix : ")
	switch choix {
	case "1":
		takePot(c)
	case "2":
		takeManaPot(c)
	case "3":
		removeInventory(c)
	}
}

func addInventory(c *Character, item string) bool {
	if len(c.Inventaire) < c.CapInv {
		c.Inventaire = append(c.Inventaire, item)
		return true
	}
	fmt.Println("Inventaire plein !")
	return false
}

func removeInventory(c *Character) {
	fmt.Println("Quel objet retirer ?")
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	choix := input("Numéro : ")
	idx, err := strconv.Atoi(choix)
	if err == nil && idx > 0 && idx <= len(c.Inventaire) {
		c.Inventaire = append(c.Inventaire[:idx-1], c.Inventaire[idx:]...)
		fmt.Println("Objet retiré.")
	} else {
		fmt.Println("Choix invalide.")
	}
}

// --- Potions ---

func takePot(c *Character) {
	for i, item := range c.Inventaire {
		if item == "Potion de soin" {
			heal := 30
			c.HP += heal
			if c.HP > c.MaxHP {
				c.HP = c.MaxHP
			}
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Println("Vous récupérez", heal, "PV !")
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de soin.")
}

func takeManaPot(c *Character) {
	for i, item := range c.Inventaire {
		if item == "Potion de mana" {
			regen := 20
			c.Mana += regen
			if c.Mana > c.MaxMana {
				c.Mana = c.MaxMana
			}
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Println("Vous récupérez", regen, "Mana !")
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de mana.")
}

func poisonPot(c *Character) {
	for i, item := range c.Inventaire {
		if item == "Potion empoisonnée" {
			c.HP -= 20
			if c.HP < 0 {
				c.HP = 0
			}
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			fmt.Println("Aïe ! C'était une potion empoisonnée !")
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion empoisonnée.")
}

// --- Système de mort/résurrection ---

func isDead(c *Character) bool {
	if c.HP <= 0 {
		fmt.Println("Vous êtes mort ! Résurrection...")
		c.HP = c.MaxHP / 2
		c.Mana = c.MaxMana / 2
		return true
	}
	return false
}

// --- Gestion des sorts ---

func spellBook(c *Character) {
	fmt.Println("Sorts disponibles :")
	fmt.Println("1. Boule de feu (10 mana, 25 dégâts)")
	fmt.Println("2. Soin magique (15 mana, +30 PV)")
	fmt.Println("3. Retour")
	choix := input("Choix : ")
	switch choix {
	case "1":
		if c.Mana >= 10 {
			c.Mana -= 10
			fmt.Println("Vous lancez une boule de feu ! (utilisable en combat)")
		} else {
			fmt.Println("Pas assez de mana.")
		}
	case "2":
		if c.Mana >= 15 {
			c.Mana -= 15
			c.HP += 30
			if c.HP > c.MaxHP {
				c.HP = c.MaxHP
			}
			fmt.Println("Vous vous soignez magiquement !")
		} else {
			fmt.Println("Pas assez de mana.")
		}
	}
}

// --- Marchand ---

func merchant(c *Character) {
	fmt.Println("Bienvenue chez le marchand !")
	fmt.Println("1. Potion de soin (15 or)")
	fmt.Println("2. Potion de mana (10 or)")
	fmt.Println("3. Potion empoisonnée (5 or)")
	fmt.Println("4. Augmenter capacité inventaire (+1, 30 or)")
	fmt.Println("5. Quitter")
	choix := input("Choix : ")
	switch choix {
	case "1":
		if c.Or >= 15 && addInventory(c, "Potion de soin") {
			c.Or -= 15
			fmt.Println("Achat réussi !")
		}
	case "2":
		if c.Or >= 10 && addInventory(c, "Potion de mana") {
			c.Or -= 10
			fmt.Println("Achat réussi !")
		}
	case "3":
		if c.Or >= 5 && addInventory(c, "Potion empoisonnée") {
			c.Or -= 5
			fmt.Println("Achat réussi !")
		}
	case "4":
		if c.Or >= 30 {
			c.Or -= 30
			c.CapInv++
			fmt.Println("Capacité d'inventaire augmentée !")
		}
	}
}

// --- Forgeron ---

func blacksmith(c *Character) {
	fmt.Println("Bienvenue chez le forgeron !")
	fmt.Println("1. Fabriquer une épée (+10 PV max, 40 or)")
	fmt.Println("2. Fabriquer une armure (+20 PV max, 70 or)")
	fmt.Println("3. Quitter")
	choix := input("Choix : ")
	switch choix {
	case "1":
		if c.Or >= 40 {
			c.Or -= 40
			c.Equipement = Equipment{"Épée", 10}
			c.MaxHP += 10
			fmt.Println("Vous avez fabriqué une épée !")
		}
	case "2":
		if c.Or >= 70 {
			c.Or -= 70
			c.Equipement = Equipment{"Armure", 20}
			c.MaxHP += 20
			fmt.Println("Vous avez fabriqué une armure !")
		}
	}
}

// --- Combat d'entraînement ---

func trainingFight(c *Character) {
	gob := Monster{"Gobelin d'entraînement", 60, 10, []int{10, 15, 5}}
	turn := 0
	fmt.Println("Un gobelin d'entraînement apparaît !")
	for gob.HP > 0 && c.HP > 0 {
		fmt.Printf("\nVotre PV : %d/%d | Gobelin : %d PV\n", c.HP, c.MaxHP, gob.HP)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser inventaire")
		fmt.Println("3. Lancer un sort")
		choix := input("Action : ")
		switch choix {
		case "1":
			dmg := 15 + rand.Intn(6)
			gob.HP -= dmg
			fmt.Println("Vous frappez le gobelin pour", dmg, "dégâts !")
		case "2":
			accessInventory(c)
		case "3":
			if c.Mana >= 10 {
				gob.HP -= 25
				c.Mana -= 10
				fmt.Println("Boule de feu ! 25 dégâts.")
			} else {
				fmt.Println("Pas assez de mana.")
			}
		}
		if gob.HP > 0 {
			att := gob.Pattern[turn%len(gob.Pattern)]
			c.HP -= att
			fmt.Println("Le gobelin attaque pour", att, "dégâts !")
			turn++
		}
		if isDead(c) {
			fmt.Println("Vous avez été vaincu par le gobelin.")
			return
		}
	}
	if gob.HP <= 0 {
		fmt.Println("Vous avez vaincu le gobelin ! +20 or, +10 exp")
		c.Or += 20
		c.Exp += 10
		if c.Exp >= 20 {
			c.Niveau++
			c.Exp = 0
			c.MaxHP += 10
			c.HP = c.MaxHP
			fmt.Println("Niveau supérieur ! PV max augmenté.")
		}
	}
}

// --- Menu principal ---

func menu() {
	char := initCharacter()
	for {
		fmt.Println("\n--- Menu principal ---")
		fmt.Println("1. Afficher infos")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Entraînement (combat)")
		fmt.Println("6. Sorts")
		fmt.Println("7. Quitter")
		choix := input("Votre choix : ")
		switch choix {
		case "1":
			displayInfo(char)
		case "2":
			accessInventory(&char)
		case "3":
			merchant(&char)
		case "4":
			blacksmith(&char)
		case "5":
			trainingFight(&char)
		case "6":
			spellBook(&char)
		case "7":
			fmt.Println("Merci d'avoir joué !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
