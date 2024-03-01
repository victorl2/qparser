package parser

type QuakeGames struct {
	GameDetails map[string]*Game `json:"games"`
}

type Game struct {
	TotalKills  int                 `json:"total_kills"`
	Players     []string            `json:"players"`
	PlayerSet   map[string]struct{} `json:"-"`
	Kills       map[string]int      `json:"kills"`
	KillByMeans map[string]int      `json:"kills_by_means"`
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

func NewGame() *Game {
	return &Game{
		Kills:       make(map[string]int),
		KillByMeans: make(map[string]int),
		Players:     make([]string, 0),
		PlayerSet:   make(map[string]struct{}),
	}
}

func (g *Game) AddPlayer(player string) {
	if _, exists := g.PlayerSet[player]; !exists {
		g.Players = append(g.Players, player)
		g.PlayerSet[player] = struct{}{}
	}
}

func (g *Game) AddKill(killAction *Killing) {
	if _, exists := g.PlayerSet[killAction.Killer]; !exists && killAction.Killer != "<world>" {
		g.Players = append(g.Players, killAction.Killer)
		g.PlayerSet[killAction.Killer] = struct{}{}
	}

	if _, exists := g.PlayerSet[killAction.Killed]; !exists {
		g.Players = append(g.Players, killAction.Killed)
		g.PlayerSet[killAction.Killed] = struct{}{}
	}

	if killAction.Killer != "<world>" {
		g.Kills[killAction.Killer]++
	} else {
		g.Kills[killAction.Killed]--
	}

	g.KillByMeans[killAction.Weapon]++
	g.TotalKills++
}
