package nicobot

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ParseCommand parses the command entered by the user to place or move
// the bot, get the report
func RunCommand(bot *Bot, action string) {
	// Last error is always the result of current cmd
	bot.lastError = nil
	bot.lastCmd = strings.TrimSpace(action)

	// Running cmds are case insensitive
	action = strings.ToLower(bot.lastCmd)

	// We create a reader to go through the cmd and arguments
	r := bytes.NewReader([]byte(action))
	cmdReader := bufio.NewReader(r)

	// The action is expected to be the first word
	cmd, err := cmdReader.ReadString(' ')
	cmd = sanitizeCmd(cmd)

	// If EOF on the first read, we expect the cmd to be one of the following
	if err == io.EOF {
		nReqArgs, ok := allowedCmds[cmd]
		if ok {
			if nReqArgs > 0 {
				bot.lastError = errors.New("NOT ENOUGH ARGUMENTS")
				return
			}

			runSingleCmd(bot, cmd)
			return
		}

		bot.lastError = fmt.Errorf("UNKNOWN COMMAND %s", cmd)
		return
	} else if err != nil {
		bot.lastError = err
		return
	}

	if isMulti, ok := allowedCmds[cmd]; ok && isMulti > 0 {
		runCmdWithArgs(bot, cmd, cmdReader)
	}
}

// runSingleCmd performs commands that do not require arguments
func runSingleCmd(bot *Bot, cmd string) {
	if cmd == Left || cmd == Right {
		bot.Turn(cmd)
	} else if cmd == Move {
		bot.Move()
	} else if cmd == Report {
		if bot.IsPlaced() == false {
			bot.lastReport = "Output: NO REPORT AVAILABLE"
		} else {
			bot.lastReport = fmt.Sprintf("Output: %d,%d,%s", bot.point.X, bot.point.Y, strings.ToUpper(bot.direction))
		}
	}
}

// runCmdWithArgs parses arguments from a command
func runCmdWithArgs(bot *Bot, cmd string, cmdReader *bufio.Reader) {
	argsStr, err := cmdReader.ReadString('\n')
	if err != nil && err != io.EOF {
		bot.lastError = err
		return
	}

	// Split args with comma delimiter
	args := strings.Split(argsStr, ",")
	if cmd == Place {
		runPlace(bot, args)
		return
	}

	bot.lastError = fmt.Errorf("UNKNOWN COMMAND %s", cmd)
}

// parsePlaceArguments parses all arguments for PLACE cmd and places
// the bot accordingly if correct parameters were provided
func runPlace(bot *Bot, args []string) {
	// We must have 3 arguments; X,Y,DIRECTION
	if len(args) != allowedCmds[Place] {
		bot.lastError = fmt.Errorf("WRONG NUMBER OF ARGUMENTS. %d EXPECTED", allowedCmds[Place])
		return
	}

	// x must be an integer
	x, err := strconv.Atoi(strings.TrimSpace(args[0]))
	if err != nil {
		bot.lastError = fmt.Errorf("%s IS NOT A VALID X POINT", args[0])
		return
	}

	// y must be an integer
	y, err := strconv.Atoi(strings.TrimSpace(args[1]))
	if err != nil {
		bot.lastError = fmt.Errorf("%s IS NOT A VALID Y POINT", args[1])
		return
	}

	// direction must exist
	facing := strings.TrimSpace(args[2])
	_, ok := dirFacing[facing]
	if ok == false {
		bot.lastError = fmt.Errorf("%s IS NOT A VALID DIRECTION", facing)
		return
	}

	bot.Place(facing, x, y)
}

// sanitizeCmd trim spaces and transforms the command to a lowercase string
func sanitizeCmd(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	return strings.ToLower(cmd)
}
