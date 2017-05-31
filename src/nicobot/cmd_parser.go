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

func ParseCommand(bot *Bot, action string) bool {
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
		if isMulti, ok := allowedCmds[cmd]; ok && isMulti == false {
			if _, ok = directions[cmd]; ok {
				bot.direction = cmd
			} else if cmd == Move {
				bot.Move(bot.direction)
			} else if cmd == Report {
				return true
			} else {
				bot.lastError = fmt.Errorf("UNKNOWN COMMAND %s", action)
				return false
			}
		}
	} else if err != nil {
		bot.lastError = err
		return false
	}

	if isMulti, ok := allowedCmds[cmd]; ok && isMulti == true {
		if cmd != Place {
			bot.lastError = fmt.Errorf("UNKNOWN COMMAND %s", action)
			return false
		}

		args, err := cmdReader.ReadString('\n')
		if err != nil && err != io.EOF {
			bot.lastError = err
			return false
		}

		parsePlaceArguments(bot, args)
		return false
	}

	return false
}

// parsePlaceArguments parses all arguments for PLACE cmd and places
// the bot accordingly if correct parameters were provided
func parsePlaceArguments(bot *Bot, argStr string) {
	// Split them with comma delimiter
	args := strings.Split(argStr, ",")

	// We must have 3 arguments; X,Y,DIRECTION
	if len(args) != 3 {
		bot.lastError = errors.New("WRONG NUMBER OF ARGUMENTS")
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
		bot.lastError = fmt.Errorf("%s IS NOT A VALID X POINT", args[1])
		return
	}

	// direction must exist
	dir := strings.TrimSpace(args[2])
	if _, ok := dirFacing[dir]; !ok {
		bot.lastError = fmt.Errorf("%s IS NOT A VALID DIRECTION", dir)
		return
	}

	switch dir {
	case West:
		dir = Left
	case East:
		dir = Right
	case North:
		dir = Up
	case South:
		dir = Down
	}

	bot.Place(dir, x, y)
}

func sanitizeCmd(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	return strings.ToLower(cmd)
}
