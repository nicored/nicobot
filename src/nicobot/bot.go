package nicobot

import (
	"errors"
	"fmt"
	"image"
)

const (
	// Commands
	Left   = "left"
	Right  = "right"
	Move   = "move"
	Report = "report"
	Place  = "place"

	// Direction
	NonApplicable = "N/a"
	West          = "west"
	East          = "east"
	North         = "north"
	South         = "south"

	// Angle for each direction
	DegNorth = 0
	DegEast  = 90
	DegSouth = 180
	DegWest  = 270

	// Static error messages
	ErrMsgFalling   = "BOT REFUSES TO FALL AND BREAK"
	ErrMsgNotPlaced = "BOT CANNOT MOVE IF NOT PLACED (eg. PLACE 1,2,NORTH)"
	ErrMsgOffTable  = "BOT IS OFF THE TABLE, PLACE BOT IF YOU WANT TO PLAY"

	// Static Status messages
	StatusNotPlaced = "BOT IS NOT PLACED"
	StatusOffTable  = "BOT IS OFF THE TABLE"
)

var (
	// ErrorFalling is thrown when the next move would make the bot fall
	ErrorFalling = errors.New(ErrMsgFalling)

	// ErrorNotPlaced is thrown when a move is requested but the bot has not been placed
	ErrorNotPlaced = errors.New(ErrMsgNotPlaced)

	// ErrorOffTable is thrown when a move is requested but the bot is off the time
	ErrorOffTable = errors.New(ErrMsgOffTable)

	// allowedCmds is a map listing all possible commands.
	// The value tells us if the command requires arguments to run
	allowedCmds = map[string]int{
		Left:   0,
		Right:  0,
		Report: 0,
		Move:   0,
		Place:  3,
	}

	// dirFacing is map that holds the facing directions
	// combined with their angle
	dirFacing = map[string]int{
		North: DegNorth,
		East:  DegEast,
		South: DegSouth,
		West:  DegWest,
	}
)

// Bot is a struct that holds information about the
// state of the bot (error, report, last cmd executed)
type Bot struct {
	direction  string       // from the West, East, North, South
	point      *image.Point // is the current point of the bot on or out of the table
	lastError  error        // is the last reported error if any
	lastCmd    string       // is the last command ran by the user, or read from a file
	lastReport string       // is the output of the last report that was run by the user
}

// Move updates the Bot's coordinates by incrementing x or y
// based on the direction. The function returns an error if the move
// cannot be completed; ErrorNotPlaced or ErrorFalling
func (b *Bot) Move() {
	b.lastError = nil

	if b.IsPlaced() == false {
		b.lastError = ErrorNotPlaced
		return
	} else if !isOnTable(b.point.X, b.point.Y) {
		b.lastError = ErrorOffTable
		return
	}

	if b.direction == East || b.direction == West {
		b.lastError = b.moveLeftRight(b.direction)
		return
	}

	b.lastError = b.moveUpDown(b.direction)
}

// Turn sets the new direction the bot is facing
// based on the turning direction that is provided. left / right
// The bot spins by 90 or -90 degrees (right and left respectively)
func (b *Bot) Turn(turningDir string) {
	turningAngle := 90
	if turningDir == Left {
		turningAngle = -90
	}

	angle := (dirFacing[b.direction] + turningAngle + 360) % 360
	for dir, dirAngle := range dirFacing {
		if dirAngle != angle {
			continue
		}

		b.direction = dir
	}
}

// Place places the bot at a given point (in or out of the table)
// and facing a given direction (west, east, north, south)
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

	return fmt.Sprintf("DIR: %s | X: %d | Y: %d", b.direction, b.point.X, b.point.Y)
}

// IsMovable returns true if the bot can move
// The function will return false if the bot is not placed, or if it
// is outside of the table's boundaries. True otherwise
func (b Bot) IsPlaced() bool {
	return b.point != nil
}

// moveLeftRight moves the robot by 1 unit left or right on the table
// and prevents the bot from falling if the next move takes it out
// of the table's bounds
func (b *Bot) moveLeftRight(dir string) error {
	move := 1
	if dir == West {
		move = -1
	}

	if isOnTable(b.point.X+move, 0) == false {
		return ErrorFalling
	}

	b.point.X += move
	return nil
}

// moveUpDown moves the robot by 1 unit up or down on the table
// and prevents the bot from falling if the next move takes it out
// of the table's bounds
func (b *Bot) moveUpDown(dir string) error {
	move := 1
	if dir == South {
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
