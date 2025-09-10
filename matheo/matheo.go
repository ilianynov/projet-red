package main

import (
	"fmt"
	"time"
)

func poisonPot(c *Character) {
	for i := 0; i < 3; i++ {
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("HP: %d/%d\n", c.HP, c.MaxHP)
		time.Sleep(1 * time.Second)
	}
}

func isDead(c *Character) bool {
	if c.HP <= 0 {
		fmt.Println("je vous croyais plus fort... Résurrection ?")
		c.HP = c.MaxHP / 2
		return true
	}
	return false
}

func inventorylimit(c*character)
for inventorylimit 
