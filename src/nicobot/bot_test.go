package nicobot

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot_Turn(t *testing.T) {
	bot := &Bot{
		direction: North,
	}

	bot.Turn(Right)
	assert.Equal(t, East, bot.direction)

	bot.Turn(Right)
	assert.Equal(t, South, bot.direction)

	bot.Turn(Right)
	assert.Equal(t, West, bot.direction)

	bot.Turn(Right)
	assert.Equal(t, North, bot.direction)

	bot.Turn(Left)
	assert.Equal(t, West, bot.direction)

	bot.Turn(Left)
	assert.Equal(t, South, bot.direction)
}

func TestBot_IsPlaced(t *testing.T) {
	var actual bool

	b := Bot{}
	actual = b.IsPlaced()
	assert.False(t, actual)

	b.point = &image.Point{1, 2}
	actual = b.IsPlaced()
	assert.True(t, actual)

	b.point.X = 10
	actual = b.IsPlaced()
	assert.True(t, actual)
}

func TestBot_Move(t *testing.T) {
	b := &Bot{}

	b.Move()
	assert.Error(t, ErrorNotPlaced, b.lastError)

	b.Place(East, 0, 0)
	b.Move()
	assert.Equal(t, 1, b.point.X)
	assert.Equal(t, 0, b.point.Y)

	b.Place(West, 10, 10)
	b.Move()
	assert.Equal(t, ErrorOffTable, b.lastError)
	assert.Equal(t, 10, b.point.X)
	assert.Equal(t, 10, b.point.Y)

	b.Place(West, 3, 1)
	b.Move()
	assert.Equal(t, 2, b.point.X)
	assert.Equal(t, 1, b.point.Y)

	b.Place(East, 4, 2)
	b.Move()
	assert.Equal(t, ErrorFalling, b.lastError)
	assert.Equal(t, 4, b.point.X)
	assert.Equal(t, 2, b.point.Y)

	b.Place(North, 1, 2)
	b.Move()
	assert.Equal(t, 1, b.point.X)
	assert.Equal(t, 3, b.point.Y)

	b.Place(North, 2, 4)
	b.Move()
	assert.Equal(t, ErrorFalling, b.lastError)
	assert.Equal(t, 2, b.point.X)
	assert.Equal(t, 4, b.point.Y)

	b.Place(South, 1, 1)
	b.Move()
	assert.Equal(t, 1, b.point.X)
	assert.Equal(t, 0, b.point.Y)
}

func TestBot_String(t *testing.T) {
	b := &Bot{}

	status := b.String()
	assert.Equal(t, StatusNotPlaced, status)

	b.Place(West, 10, 10)
	status = b.String()
	assert.Equal(t, StatusOffTable, status)

	b.Place(West, 1, 3)
	status = b.String()
	assert.Equal(t, "DIR: west | X: 1 | Y: 3", status)
}
