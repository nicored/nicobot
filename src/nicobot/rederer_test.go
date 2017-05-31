package nicobot

import (
	"testing"

	"bufio"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
)

var botDirTestCases = []struct {
	Dir      string
	Expected string
}{
	{
		Dir: East,
		Expected: "███████    " +
			"  ███████  " +
			"    ███████" +
			"  ███████  " +
			"███████    ",
	},
	{
		Dir: West,
		Expected: "    ███████" +
			"  ███████  " +
			"███████    " +
			"  ███████  " +
			"    ███████",
	},
	{
		Dir: North,
		Expected: "     █     " +
			"    ███    " +
			"  ███ ███  " +
			" ███   ███ " +
			"███     ███",
	},
	{
		Dir: South,
		Expected: "███     ███" +
			" ███   ███ " +
			"  ███ ███  " +
			"    ███    " +
			"     █     ",
	},
	{
		Dir: NonApplicable,
		Expected: "           " +
			"           " +
			"           " +
			"           " +
			"           ",
	},
}

func TestGetBotLine(t *testing.T) {
	for _, testCase := range botDirTestCases {
		actual := ""
		for i := 0; i < 5; i++ {
			actual += getBotLine(testCase.Dir, i)
		}
		assert.Equal(t, testCase.Expected, actual)
	}
}

func TestRenderer_RenderWithCmd(t *testing.T) {
	headerLines := 1
	bottomLines := 1
	sepRows := ConsoleUnitsPerTableUnitsY - 1
	expectedLines := TableUnitsY*ConsoleUnitsPerTableUnitsY + headerLines + bottomLines + sepRows

	bot := &Bot{}
	renderer := &Renderer{}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderer.RenderWithCmd(bot, "PLACE 0,4,NORTH")
	bufReader := bufio.NewReader(r)

	// back to normal state
	os.Stdout = old // restoring the real stdout
	w.Close()

	lines := 0
	for {
		line, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if lines == 0 {
			assert.Equal(t, "\n", line)
			lines++
			continue
		} else if lines > headerLines+TableUnitsY*ConsoleUnitsPerTableUnitsY+sepRows {
			lines++
			continue
		}

		lines++
	}

	assert.Equal(t, expectedLines, lines)

}
