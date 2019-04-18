package nicobot

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestRunCommand(t *testing.T) {
	bot := &Bot{}

	testFlow := []struct {
		Cmd    string
		Point  *image.Point
		Dir    string
		Error  error
		Report string
	}{
		{"MOVE", nil, "", ErrorNotPlaced, ""},
		{"REPORT", nil, "", nil, "Output: NO REPORT AVAILABLE"},
		{"PLACE 1,1,NORTH", &image.Point{1, 1}, North, nil, "Output: NO REPORT AVAILABLE"},
		{"REPORT", &image.Point{1, 1}, North, nil, "Output: 1,1,NORTH"},
		{"MOVE", &image.Point{1, 2}, North, nil, "Output: 1,1,NORTH"},
		{"RIGHT", &image.Point{1, 2}, East, nil, "Output: 1,1,NORTH"},
		{"RIGHT", &image.Point{1, 2}, South, nil, "Output: 1,1,NORTH"},
		{"LEFT", &image.Point{1, 2}, East, nil, "Output: 1,1,NORTH"},
		{"MOVE", &image.Point{2, 2}, East, nil, "Output: 1,1,NORTH"},
		{"MOVE", &image.Point{3, 2}, East, nil, "Output: 1,1,NORTH"},
		{"MOVE", &image.Point{4, 2}, East, nil, "Output: 1,1,NORTH"},
		{"MOVE", &image.Point{4, 2}, East, ErrorFalling, "Output: 1,1,NORTH"},
		{"REPORT", &image.Point{4, 2}, East, nil, "Output: 4,2,EAST"},
		{"LEFT", &image.Point{4, 2}, North, nil, "Output: 4,2,EAST"},
		{"MOVE", &image.Point{4, 3}, North, nil, "Output: 4,2,EAST"},
		{"MOVE", &image.Point{4, 4}, North, nil, "Output: 4,2,EAST"},
		{"REPORT", &image.Point{4, 4}, North, nil, "Output: 4,4,NORTH"},
		{"MOVE", &image.Point{4, 4}, North, ErrorFalling, "Output: 4,4,NORTH"},
		{"PLACE 5,5,south", &image.Point{5, 5}, South, nil, "Output: 4,4,NORTH"},
		{"MOVE", &image.Point{5, 5}, South, ErrorOffTable, "Output: 4,4,NORTH"},
		{"PLACE 5,5,5,5", &image.Point{5, 5}, South, errors.New("WRONG NUMBER OF ARGUMENTS. 3 EXPECTED"), "Output: 4,4,NORTH"},
		{"UP", &image.Point{5, 5}, South, errors.New("UNKNOWN COMMAND up"), "Output: 4,4,NORTH"},
		{"PLACE", &image.Point{5, 5}, South, errors.New("NOT ENOUGH ARGUMENTS"), "Output: 4,4,NORTH"},
		{"PLACE L,0,NORTH", &image.Point{5, 5}, South, errors.New("l IS NOT A VALID X POINT"), "Output: 4,4,NORTH"},
		{"PLACE 0,r,NORTH", &image.Point{5, 5}, South, errors.New("r IS NOT A VALID Y POINT"), "Output: 4,4,NORTH"},
		{"PLACE 0,0,H", &image.Point{5, 5}, South, errors.New("h IS NOT A VALID DIRECTION"), "Output: 4,4,NORTH"},
	}

	for i, tc := range testFlow {
		RunCommand(bot, tc.Cmd)

		if tc.Point == nil {
			assert.Nil(t, bot.point)
		} else {
			assert.Equal(t, tc.Point.X, bot.point.X)
			assert.Equal(t, tc.Point.Y, bot.point.Y)

		}

		assert.Equal(t, tc.Dir, bot.direction, "Wrong dir for test case at index %d", i)
		assert.Equal(t, tc.Report, bot.lastReport, "Wrong report for test case at index %d", i)

		if tc.Error != nil && bot.lastError != nil {
			assert.Equal(t, tc.Error.Error(), bot.lastError.Error())
		} else if tc.Error != bot.lastError {
			t.Fatalf("Expected and actual bot errors do not match for test flow at index %d", i)
		}
	}
}
