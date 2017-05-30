package nicobot

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.False(t, actual)

	b.point.X = 0
	b.point.Y = 10
	actual = b.IsPlaced()
	assert.False(t, actual)

	b.point.X = 5
	actual = b.IsPlaced()
	assert.False(t, actual)
}

func TestBot_Move(t *testing.T) {
	var err error

	b := Bot{}
	err = b.Move(right)
	assert.Equal(t, ErrorNotPlaced, err)

	b.point = &image.Point{10, 10}
	err = b.Move(right)
	assert.Equal(t, ErrorNotPlaced, err)

	b.point.X = 0
	b.point.Y = 0
	err = b.Move(left)
	assert.Equal(t, ErrorFalling, err)
	assert.Equal(t, 0, b.point.X)
	assert.Equal(t, 0, b.point.Y)
	assert.Equal(t, left, b.direction)

	err = b.Move(right)
	assert.NoError(t, err)
	assert.Equal(t, 1, b.point.X)
	assert.Equal(t, 0, b.point.Y)
	assert.Equal(t, right, b.direction)

	err = b.Move(down)
	assert.Equal(t, ErrorFalling, err)
	assert.Equal(t, 1, b.point.X)
	assert.Equal(t, 0, b.point.Y)
	assert.Equal(t, down, b.direction)

	err = b.Move(up)
	assert.NoError(t, err)
	assert.Equal(t, 1, b.point.X)
	assert.Equal(t, 1, b.point.Y)
	assert.Equal(t, up, b.direction)

	err = b.Move(down)
	assert.NoError(t, err)
	assert.Equal(t, 1, b.point.X)
	assert.Equal(t, 0, b.point.Y)
	assert.Equal(t, down, b.direction)
}

func TestBot_Direction(t *testing.T) {
	var actual string

	b := Bot{}

	actual = b.Faces()
	assert.Equal(t, "N/A", actual)

	b.direction = down
	actual = b.Faces()
	assert.Equal(t, "SOUTH", actual)

	b.direction = up
	actual = b.Faces()
	assert.Equal(t, "NORTH", actual)

	b.direction = left
	actual = b.Faces()
	assert.Equal(t, "WEST", actual)

	b.direction = right
	actual = b.Faces()
	assert.Equal(t, "EAST", actual)

}
