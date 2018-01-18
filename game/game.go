package game

import (
	"github.com/jtsalva/balls/player"
	// "log"
	// "math"
)

type GameInstance struct {
	Players []player.ConnectedPlayer
	Quit    bool
}

var Instance GameInstance

func (g *GameInstance) UpdatePlayers() {
	// Move Players
	for i, _ := range g.Players {
		g.Players[i].Move()
	}
}

func (g *GameInstance) ConnectPlayer(playerId int, playerName string, mousePos [2]int) {
	if !g.PlayerExists(playerId) {
		g.Players = append(g.Players, player.ConnectedPlayer{
			Id:       playerId,
			Name:     playerName,
			MousePos: mousePos,
			BallPos:  [2]int{0, 0}})
	}
}

func (g *GameInstance) SyncPlayerMouse(playerId int, mousePos [2]int) {
	for i, _ := range g.Players {
		if g.Players[i].Id == playerId {
			g.Players[i].MousePos = mousePos
			break
		}
	}
}

func (g *GameInstance) PlayerExists(playerId int) bool {
	for i, _ := range g.Players {
		if g.Players[i].Id == playerId {
			return true
		}
	}
	return false
}

func init() {
	Instance.Quit = false
}
