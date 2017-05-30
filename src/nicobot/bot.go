package nicobot

import (
	"errors"
	"image"
)

const (
	_    = iota
	left // starts at 1
	down
	up
	right
)

var (
	ErrorFalling   = errors.New("Robot is going to fall")
	ErrorNotPlaced = errors.New("Robot needs to be placed first (eg. PLACE 1,2,NORTH)")
)

type Bot struct {
	direction int // from the left, down, up and right constants
	point     *image.Point
}

// Move updates the Bot's coordinates by incrementing x or y
// based on the direction. The function returns an error if the move
// cannot be completed; ErrorNotPlaced or ErrorFalling
func (b *Bot) Move(dir int) error {
	if b.IsPlaced() == false {
		return ErrorNotPlaced
	}

	// No matter if the bot can move or not, we want to the direction
	b.direction = dir

	if dir == right || dir == left {
		return b.moveHorizontally(dir)
	}

	return b.moveVertically(dir)
}

// Faces returns the direction the bot is based on the previous move.
// Possible outputs are: N/A, EAST, WEST, SOUTH, NORTH
func (b Bot) Faces() string {
	switch b.direction {
	case right:
		return "EAST"
	case left:
		return "WEST"
	case down:
		return "SOUTH"
	case up:
		return "NORTH"
	default:
		return "N/A"
	}
}

// IsMovable returns true if the bot can move
// The function will return false if the bot is not placed, or if it
// is outside of the table's boundaries. True otherwise
func (b Bot) IsPlaced() bool {
	return b.point != nil && isOnTable(b.point.X, b.point.Y)
}

func (b *Bot) moveHorizontally(dir int) error {
	move := 1
	if dir == left {
		move = -1
	}

	if isOnTable(b.point.X+move, 0) == false {
		return ErrorFalling
	}

	b.point.X += move
	return nil
}

func (b *Bot) moveVertically(dir int) error {
	move := 1
	if dir == down {
		move = -1
	}

	if isOnTable(0, b.point.Y+move) == false {
		return ErrorFalling
	}

	b.point.Y += move
	return nil
}

// isOnTable returns false if the given x and y are outside of the table's boundaries
func isOnTable(x, y int) bool {
	return x >= 0 && x < TableUnitsX && y >= 0 && y < TableUnitsY
}
