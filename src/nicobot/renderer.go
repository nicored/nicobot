package nicobot

import (
	"bufio"
	"fmt"
	"os"

	"deco"
	"github.com/pkg/errors"
)

const (
	TableUnitsX = 5
	TableUnitsY = 5
)

// Renderer is a struct that make it possible to render the
// table and the bot
type Renderer struct {
	lines  int
	report bool
}

// unitLine is a plain line
var unitLine = "███████████"

// horiBot is a slice of strings where each
// item, in order, represent the bot facing EAST
var horiBot = []string{
	"███████    ",
	"  ███████  ",
	"    ███████",
	"  ███████  ",
	"███████    ",
}

// vertBot is a slice of strings where each
// item, in order, represent the bot facing NORTH
var vertBot = []string{
	"     █     ",
	"    ███    ",
	"  ███ ███  ",
	" ███   ███ ",
	"███     ███",
}

// Render renders the table, as well as the bot command prompt,
// the errors, the report and the state of the bot
func (r *Renderer) RenderWithPrompt(bot *Bot) {
	for {
		deco.Clear(r.lines)
		r.lines = 0

		r.renderError(bot.lastError)
		r.renderRows(bot)
		r.renderStatus(bot)
		r.renderCmdPrompt(bot)

		if r.report {
			r.renderReport(bot)
			r.report = false
		}
	}
}

// RenderWithCmd takes a cmd, parses it, and render the table, the bot,
// the errors, the report and the state of the bot
func (r *Renderer) RenderWithCmd(bot *Bot, cmd string) {
	r.report = ParseCommand(bot, cmd)

	deco.Clear(r.lines)
	r.lines = 0

	r.renderError(bot.lastError)
	r.renderRows(bot)
	r.renderStatus(bot)

	if r.report {
		r.renderReport(bot)
		r.report = false
	}
}

// renderBotLine prints a specific line for the bot in a single unit
func (r Renderer) renderBotLine(dir int, line int) {
	fmt.Print(getBotLine(dir, line))
}

// getBotLine returns a string containing the content of a specific
// line for the bot in a single unit
func getBotLine(dir int, line int) string {
	switch dir {
	case Right:
		return horiBot[line]
	case Up:
		return vertBot[line]
	case Down:
		return vertBot[len(vertBot)-(1+line)]
	case Left:
		reversed := ""
		for _, prune := range horiBot[line] {
			reversed = string(prune) + reversed
		}
		return reversed
	default:
		return fmt.Sprintf("%-11s", "")
	}
}

// renderRows prints out all rows representing the table to the console
func (r *Renderer) renderRows(bot *Bot) {
	for y := 0; y < TableUnitsY*5; y++ {
		if y%5 == 0 && y > 0 {
			r.renderRowSep()
		}

		r.renderColumns(bot, y)
		r.renderRowSep()
	}
}

// renderCmdPrompt display the a cmd prompt to allow
// the user to place and move the bot, as well as displaying the report
func (r *Renderer) renderCmdPrompt(bot *Bot) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Action: ")
	r.lines++

	action, err := reader.ReadString('\n')
	if err != nil {
		bot.lastError = errors.New("Oops, Something went wrong when reading your command")
		return
	}

	r.report = ParseCommand(bot, action)
}

// renderColumns prints all columns in a row to the console
func (r *Renderer) renderColumns(bot *Bot, y int) {
	for x := 0; x < TableUnitsX; x++ {
		if x > 0 {
			r.renderColSep()
		}

		if bot.point != nil && bot.point.X == x && 4-bot.point.Y == y/5 {
			deco.Color(deco.TextRed)
			r.renderBotLine(bot.direction, y%5)
			deco.Color(deco.DefaultColor)
		} else {
			r.renderUnitLine()
		}
	}
}

// renderStatus prints out information about the state of the bot
// in human readable fashion
func (r *Renderer) renderStatus(bot *Bot) {
	report := bot.String()

	fmt.Printf(deco.BgBlack+"%-30s%29s\n"+deco.DefaultColor, bot.lastCmd, report)

	r.lines++
}

// renderError prints out an error message with a red background
// If no msg is provided, a new line is still created as a placeholder
// but the background will remain as the default one
func (r *Renderer) renderError(err error) {
	r.lines++

	if err == nil {
		fmt.Println()
		return
	}

	fmt.Printf(deco.BgRed+"%-59s\n"+deco.DefaultColor, err)
}

// renderReport prints a new line containing report's information
// in the following format: "Output: X,Y,DIRECTION
func (r *Renderer) renderReport(bot *Bot) {
	var output string

	if bot.point == nil {
		output = "Output: NO REPORT AVAILABLE"
	} else {
		output = fmt.Sprintf("Output: %d,%d,%s", bot.point.X, bot.point.Y, bot.Facing())
	}

	fmt.Printf("%-59s", output)
	r.lines++
}

// renderColSep prints a character separator between units on the same line
func (r Renderer) renderColSep() {
	fmt.Print(" ")
}

// renderRowSep prints a line separator between units
func (r *Renderer) renderRowSep() {
	fmt.Println()
	r.lines++
}

// renderUnitLine prints a plain line in a single unit to the console
func (r Renderer) renderUnitLine() {
	fmt.Print(unitLine)
}
