package chess

import (
	"fmt"
	"math"
)

type Position struct {
	X int
	Y int
}

func parsePosition(pos string) (Position, error) {
	if len(pos) != 2 {
		return Position{}, fmt.Errorf("invalid position: %s", pos)
	}
	x := pos[0]
	y := pos[1]
	if x < 'a' || x > 'h' || y < '1' || y > '8' {
		return Position{}, fmt.Errorf("invalid position: %s", pos)
	}
	return Position{int(x - 'a'), int(y - '1')}, nil
}

func CanKnightAttack(white, black string) (bool, error) {
	whitePos, err := parsePosition(white)
	if err != nil {
		return false, err
	}
	blackPos, err := parsePosition(black)
	if err != nil {
		return false, err
	}
	if whitePos.X == blackPos.X && whitePos.Y == blackPos.Y {
		return false, fmt.Errorf("figures on one cell: %s", white)
	}

	numWhiteCell := whitePos.X + whitePos.Y*8
	numBlackCell := blackPos.X + blackPos.Y*8
	def := math.Abs(float64(numWhiteCell - numBlackCell))

	if def == 6 || def == 10 || def == 15 || def == 17 {
		return true, nil
	}

	return false, nil
}
