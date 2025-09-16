package main

import (
	"fmt"
	"time"
)

func PoisonPot(c *Character) {
	for i := 0; i < 3; i++ {
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("HP: %d/%d\n", c.HP, c.MaxHP)
		time.Sleep(1 * time.Second)
	}
}

func IsDead(c *Character) bool {
	if c.HP <= 0 {
		fmt.Println("      ______")
		fmt.Println("   .-        -.")
		fmt.Println("  /            \\")
		fmt.Println(" |,  .-.  .-.  ,|")
		fmt.Println(" | )(_o/  \\o_)( |")
		fmt.Println(" |/     /\\     \\|")
		fmt.Println(" (_     ^^     _)")
		fmt.Println("  \\__|IIIIII|__/")
		fmt.Println("   | \\IIIIII/ |")
		fmt.Println("   \\          /")
		fmt.Println("    `--------`")
		fmt.Println("je vous croyais plus fort... Résurrection ?")
		c.HP = c.MaxHP / 2
		return true
	}
	return false
}
