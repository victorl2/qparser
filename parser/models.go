package parser

import "fmt"

type QuakeGames struct {
	GameDetails map[string]*Game `json:"games"` // Use pointers to Games for easy updates
}

type Game struct {
	TotalKills  int            `json:"total_kills"`
	Players     []string       `json:"players"`
	Kills       map[string]int `json:"kills"`
	KillByMeans map[string]int `json:"kills_by_means"`
}

type Killing struct {
	Killer string
	Killed string
	Weapon string
}

func NewGroupQuakeGames() *QuakeGames {
	return &QuakeGames{
		GameDetails: make(map[string]*Game),
	}
}

func (group *QuakeGames) AddGame(gameID string, game *Game) {
	group.GameDetails[gameID] = game
}

// NewGame initializes a new Game instance
func NewGame() *Game {
	return &Game{
		Kills: make(map[string]int),
	}
}

// AddKill increments the kill count for a player and total kills
func (g *Game) AddKill(killAction *Killing) {
	if _, exists := g.Kills[killAction.Killer]; !exists {
		g.Players = append(g.Players, killAction.Killer)
	}

	if _, exists := g.Kills[killAction.Killed]; !exists {
		g.Players = append(g.Players, killAction.Killed)
	}

	if killAction.Killer != "<world>" {
		g.Kills[killAction.Killer]++
	} else {
		g.Kills[killAction.Killed]--
	}

	// if the weapon is not already in the map, add it
	if _, exists := g.KillByMeans[killAction.Weapon]; !exists {
		fmt.Println("Adding weapon to map")
		fmt.Println(killAction.Weapon)
		g.KillByMeans[killAction.Weapon] = 1
	} else {
		g.KillByMeans[killAction.Weapon] += 1
	}

	g.TotalKills++
}
