package matheo

import (
	"fmt"
	karl "projet_red/Karl"
	"projet_red/character"
	"time"
)

func PoisonPot(c *character.Character) {
	for i := 0; i < 3; i++ {
		c.HPactuel -= 10
		if c.HPactuel < 0 {
			c.HPactuel = 0
		}
		fmt.Printf("HP: %d/%d\n", c.HPactuel, c.MaxHP)
		time.Sleep(1 * time.Second)
	}
}

func IsDead(c *character.Character) bool {
	if c.HPactuel <= 0 {
		ShowDeathArt()
		fmt.Println("je vous croyais plus fort... Résurrection ?")
		c.HPactuel = c.MaxHP / 2
		return true
	}
	return false
}

func MonsterAttack(m *karl.Monster, c *character.Character) {
	damage := m.Attaque
	c.HPactuel -= damage
	if c.HPactuel < 0 {
		c.HPactuel = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts.\n", m.Nom, c.Nom, damage)
	fmt.Printf("%s : %d/%d PV\n", c.Nom, c.HPactuel, c.MaxHP)
}

func ShowDeathArt() {
	fmt.Println(
		"      ______\n" +
			"   .-        -.\n" +
			"  /            \\\n" +
			" |,  .-.  .-.  ,|\n" +
			" | )(_o/  \\o_)( |\n" +
			" |/     /\\     \\\n" +
			" (_     ^^     _)\n" +
			"  \\__|IIIIII|__/\n" +
			"   | \\IIIIII/ |\n" +
			"   \\          /\n" +
			"    `--------`\n",
	)
}
