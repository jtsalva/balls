package player

import (
	"github.com/jtsalva/balls/game/constants"
	"math"
)

// Used when passing around general player info
type Player struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Used in game lobbies to update and keep track of positions
type ConnectedPlayer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	MousePos [2]int `json:"mousePos"`
	BallPos  [2]int `json:"ballPos"`
}

func (p *ConnectedPlayer) Move() {
	fBallPos := PosToFloat64(p.BallPos)
	fMousePos := PosToFloat64(p.MousePos)

	dx := fMousePos[0] - fBallPos[0]
	dy := fMousePos[1] - fBallPos[1]

	// Euclidean distance
	distance := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

	if distance > 0 {
		// Scale factor
		sf := float64(constants.VELOCITY) / distance

		if sf <= 1 {
			// Scale ball positions by the scale factor
			movementVector := [2]float64{dx * sf, dy * sf}

			p.BallPos = MovePosByVector(p.BallPos, VectorToInt(movementVector))
		}
	}
}
