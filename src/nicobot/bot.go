package nicobot

import (
	"errors"
	"fmt"
	"image"
)

const (
	Left  = "left"
	Down  = "down"
	Up    = "up"
	Right = "right"

	Move   = "move"
	Report = "report"
	Place  = "place"

	NonApplicable = "N/a"
	West          = "west"
	East          = "east"
	North         = "north"
	South         = "south"

	ErrMsgFalling   = "BOT REFUSES TO FALL AND BREAK"
	ErrMsgNotPlaced = "BOT CANNOT MOVE IF NOT PLACED (eg. PLACE 1,2,NORTH)"
	ErrMsgOffTable  = "BOT IS OFF THE TABLE, PLACE BOT IF YOU WANT TO PLAY"

	StatusNotPlaced = "BOT IS NOT PLACED"
	StatusOffTable  = "BOT IS OFF THE TABLE"
)

var (
	ErrorFalling   = errors.New(ErrMsgFalling)
	ErrorNotPlaced = errors.New(ErrMsgNotPlaced)

	allowedCmds = map[string]bool{
		Left:   false,
		Right:  false,
		Down:   false,
		Up:     false,
		Report: false,
		Move:   false,
		Place:  true,
	}

	directions = map[string]bool{
		Left:  true,
		Right: true,
		Down:  true,
		Up:    true,
	}

	dirFacing = map[string]bool{
		West:  true,
		East:  true,
		South: true,
		North: true,
	}
)

type Bot struct {
	direction string // from the West, East, North, South
	point     *image.Point
	lastError error
	lastCmd   string
}

// Move updates the Bot's coordinates by incrementing x or y
// based on the direction. The function returns an error if the move
// cannot be completed; ErrorNotPlaced or ErrorFalling
func (b *Bot) Move(dir string) {
	// No matter if the bot can move or not, we want to the direction
	b.direction = dir

	if b.IsPlaced() == false {
		b.lastError = ErrorNotPlaced
		return
	} else if !isOnTable(b.point.X, b.point.Y) {
		b.lastError = errors.New(ErrMsgOffTable)
		return
	}

	if dir == Right || dir == Left {
		b.lastError = b.moveHorizontally(dir)
		return
	}

	b.lastError = b.moveVertically(dir)
}

func (b *Bot) Place(dir string, x int, y int) {
	b.point = &image.Point{x, y}
	b.direction = dir
}

// String returns the direction and coordinates of the bot
// in a readable string format
func (b Bot) String() string {
	if b.point == nil {
		return StatusNotPlaced
	} else if !isOnTable(b.point.X, b.point.Y) {
		return StatusOffTable
	}

	return fmt.Sprintf("DIR: %s | X: %d | Y: %d", b.Facing(), b.point.X, b.point.Y)
}

// Facing returns the direction the bot is based on the previous move.
// Possible outputs are: N/A, EAST, WEST, SOUTH, NORTH
func (b Bot) Facing() string {
	switch b.direction {
	case Right:
		return East
	case Left:
		return West
	case Down:
		return South
	case Up:
		return North
	default:
		return NonApplicable
	}
}

// IsMovable returns true if the bot can move
// The function will return false if the bot is not placed, or if it
// is outside of the table's boundaries. True otherwise
func (b Bot) IsPlaced() bool {
	return b.point != nil
}

func (b *Bot) moveHorizontally(dir string) error {
	move := 1
	if dir == Left {
		move = -1
	}

	if isOnTable(b.point.X+move, 0) == false {
		return ErrorFalling
	}

	b.point.X += move
	return nil
}

func (b *Bot) moveVertically(dir string) error {
	move := 1
	if dir == Down {
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
