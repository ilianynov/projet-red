package main

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
func initOgre() Monster {
	return Monster{
		Nom:      "Ogre",
		MaxHP:    60,
		HPactuel: 60,
		Attaque:  10,
	}
}
func initLoupGarou() Monster {
	return Monster{
		Nom:      "Loup-Garou",
		MaxHP:    80,
		HPactuel: 80,
		Attaque:  15,
	}
}
func initDragon() Monster {
	return Monster{
		Nom:      "Dragon",
		MaxHP:    150,
		HPactuel: 150,
		Attaque:  25,
	}
}
