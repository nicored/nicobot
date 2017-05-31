package nicobot

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetBotLine(t *testing.T) {
	testCases := []struct {
		Dir      string
		Expected string
	}{
		{
			Dir: Right,
			Expected: "███████    " +
				"  ███████  " +
				"    ███████" +
				"  ███████  " +
				"███████    ",
		},
		{
			Dir: Left,
			Expected: "    ███████" +
				"  ███████  " +
				"███████    " +
				"  ███████  " +
				"    ███████",
		},
		{
			Dir: Up,
			Expected: "     █     " +
				"    ███    " +
				"  ███ ███  " +
				" ███   ███ " +
				"███     ███",
		},
		{
			Dir: Down,
			Expected: "███     ███" +
				" ███   ███ " +
				"  ███ ███  " +
				"    ███    " +
				"     █     ",
		},
		{
			Dir: 0,
			Expected: "           " +
				"           " +
				"           " +
				"           " +
				"           ",
		},
	}

	for _, testCase := range testCases {
		actual := ""
		for i := 0; i < 5; i++ {
			actual += getBotLine(testCase.Dir, i)
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}
