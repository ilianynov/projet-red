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
	fmt.Println("   _____")
	fmt.Println("  /     \\")
	fmt.Println(" | () () |")
	fmt.Println(" |  ^_^  |")
	fmt.Println(" | \___/ |")
	fmt.Println("    .-''''''-.")
	fmt.Println("  .'  _    _  '.")
	fmt.Println(" /   (o)  (o)   \")
	fmt.Println("|      /\       |")
	fmt.Println("|     ____      |")
	fmt.Println(" \   (____)   /")
	fmt.Println("  '.        .'")
	fmt.Println("    '-....-'")
		return false
}
