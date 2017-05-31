package nicobot

//
//import (
//	"image"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestBot_IsPlaced(t *testing.T) {
//	var actual bool
//
//	b := Bot{}
//	actual = b.IsPlaced()
//	assert.False(t, actual)
//
//	b.point = &image.Point{1, 2}
//	actual = b.IsPlaced()
//	assert.True(t, actual)
//
//	b.point.X = 10
//	actual = b.IsPlaced()
//	assert.False(t, actual)
//
//	b.point.X = 0
//	b.point.Y = 10
//	actual = b.IsPlaced()
//	assert.False(t, actual)
//
//	b.point.X = 5
//	actual = b.IsPlaced()
//	assert.False(t, actual)
//}
//
//func TestBot_Move(t *testing.T) {
//	var err error
//
//	b := Bot{}
//	err = b.Move(Right)
//	assert.Equal(t, ErrorNotPlaced, err)
//
//	b.point = &image.Point{10, 10}
//	err = b.Move(Right)
//	assert.Equal(t, ErrorNotPlaced, err)
//
//	b.point.X = 0
//	b.point.Y = 0
//	b.Move(Left)
//	assert.Equal(t, ErrorFalling, err)
//	assert.Equal(t, 0, b.point.X)
//	assert.Equal(t, 0, b.point.Y)
//	assert.Equal(t, Left, b.direction)
//
//	b.Move(Right)
//	assert.NoError(t, err)
//	assert.Equal(t, 1, b.point.X)
//	assert.Equal(t, 0, b.point.Y)
//	assert.Equal(t, Right, b.direction)
//
//	err = b.Move(Down)
//	assert.Equal(t, ErrorFalling, err)
//	assert.Equal(t, 1, b.point.X)
//	assert.Equal(t, 0, b.point.Y)
//	assert.Equal(t, Down, b.direction)
//
//	err = b.Move(Up)
//	assert.NoError(t, err)
//	assert.Equal(t, 1, b.point.X)
//	assert.Equal(t, 1, b.point.Y)
//	assert.Equal(t, Up, b.direction)
//
//	err = b.Move(Down)
//	assert.NoError(t, err)
//	assert.Equal(t, 1, b.point.X)
//	assert.Equal(t, 0, b.point.Y)
//	assert.Equal(t, Down, b.direction)
//}
//
//func TestBot_Direction(t *testing.T) {
//	var actual string
//
//	b := Bot{}
//
//	actual = b.Facing()
//	assert.Equal(t, "N/A", actual)
//
//	b.direction = Down
//	actual = b.Facing()
//	assert.Equal(t, "SOUTH", actual)
//
//	b.direction = Up
//	actual = b.Facing()
//	assert.Equal(t, "NORTH", actual)
//
//	b.direction = Left
//	actual = b.Facing()
//	assert.Equal(t, "WEST", actual)
//
//	b.direction = Right
//	actual = b.Facing()
//	assert.Equal(t, "EAST", actual)
//
//}
