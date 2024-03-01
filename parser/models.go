package parser

type QuakeGames struct {
	GameDetails map[string]*Game `json:"games"` // Use pointers to Games for easy updates
}

type Game struct {
	TotalKills  int            `json:"total_kills"`
	Players     []string       `json:"players"`
	Kills       map[string]int `json:"kills"`
	KillByMeans map[string]int `json:"kills_by_means"`
}

func NewGroupQuakeGames() *QuakeGames {
	return &QuakeGames{
		GameDetails: make(map[string]*Game),
	}
}

// NewGame initializes a new Game instance
func NewGame() *Game {
	return &Game{
		Kills: make(map[string]int),
	}
}

// AddKill increments the kill count for a player and total kills
func (g *Game) AddKill(killerName string, killedName string, weapon string) {
	if _, exists := g.Kills[killerName]; !exists {
		g.Players = append(g.Players, killerName)
	}

	if _, exists := g.Kills[killedName]; !exists {
		g.Players = append(g.Players, killedName)
	}

	if killerName != "<world>" {
		g.Kills[killerName]++
	} else {
		g.Kills[killedName]--
	}

	g.KillByMeans[weapon]++
	g.TotalKills++
}
