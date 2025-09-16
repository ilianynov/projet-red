package karl

type Recette struct {
	CoutOr     int
	Ressources map[string]int
}

var recettes = map[string]Recette{
	"Épée basique": {
		CoutOr: 10,
		Ressources: map[string]int{
			"Bois": 4,
			"Fer":  2,
		},
	},
	"Épée rare": {
		CoutOr: 20,
		Ressources: map[string]int{
			"Bois": 7,
			"Fer":  4,
		},
	},
	"Épée épique": {
		CoutOr: 25,
		Ressources: map[string]int{
			"Bois":    10,
			"Fer":     6,
			"Uranium": 1,
		},
	},
	"Épée légendaire": {
		CoutOr: 40,
		Ressources: map[string]int{
			"Bois":       12,
			"Fer":        8,
			"Uranium":    2,
			"Obsidienne": 1,
		},
	},
	"Épée démoniaque": {
		CoutOr: 60,
		Ressources: map[string]int{
			"Bois":          15,
			"Fer":           10,
			"Uranium":       3,
			"Obsidienne":    2,
			"Sang de démon": 1,
		},
	},
	"Épée angélique": {
		CoutOr: 60,
		Ressources: map[string]int{
			"Bois":         15,
			"Fer":          10,
			"Uranium":      3,
			"Obsidienne":   2,
			"Plume d'ange": 1,
		},
	},
	"Armure rare": {
		CoutOr: 20,
		Ressources: map[string]int{
			"Cuir": 5,
			"Bois": 4,
			"Fer":  3,
		},
	},
	"Armure épique": {
		CoutOr: 30,
		Ressources: map[string]int{
			"Cuir":    8,
			"Bois":    8,
			"Fer":     5,
			"Uranium": 1,
		},
	},
	"Armure légendaire": {
		CoutOr: 50,
		Ressources: map[string]int{
			"Cuir":       10,
			"Bois":       10,
			"Fer":        7,
			"Uranium":    2,
			"Obsidienne": 1,
		},
	},
	"Armure démoniaque": {
		CoutOr: 70,
		Ressources: map[string]int{
			"Cuir":          12,
			"Bois":          12,
			"Fer":           8,
			"Uranium":       3,
			"Obsidienne":    2,
			"Sang de démon": 1,
		},
	},
	"Armure angélique": {
		CoutOr: 70,
		Ressources: map[string]int{
			"Cuir":         12,
			"Bois":         12,
			"Fer":          8,
			"Uranium":      3,
			"Obsidienne":   2,
			"Plume d'ange": 1,
		},
	},
	"Petite bombe": {
		CoutOr: 30,
		Ressources: map[string]int{
			"Fer":     3,
			"Uranium": 1,
		},
	},
	"Grosse bombe": {
		CoutOr: 50,
		Ressources: map[string]int{
			"Fer":     6,
			"Uranium": 2,
		},
	},
	"Bombe nucléaire": {
		CoutOr: 100,
		Ressources: map[string]int{
			"Fer":           12,
			"Uranium":       5,
			"Obsidienne":    2,
			"Sang de démon": 1,
			"Plume d'ange":  1,
		},
	},
}
