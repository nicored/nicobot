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
		switch cmd {
		case "up":
			bot.Move(Up)
		case "left":
			bot.Move(Left)
		case "down":
			bot.Move(Down)
		case "right":
			bot.Move(Right)
		case "move":
			bot.Move(bot.direction)
		case "report":
			return true
		default:
			bot.lastError = fmt.Errorf("UNKNOWN COMMAND %s", action)
		}
		return false
	} else if err != nil {
		bot.lastError = err
		return false
	}

	// If we have arguments with the command, then we are expecting "PLACE"
	if cmd == "place" {
		// We read until the end of the line
		args, err := cmdReader.ReadString('\n')
		if err != nil && err != io.EOF {
			bot.lastError = err
			return false
		}

		// And we finally parse all arguments
		parsePlaceArguments(bot, args)
		return false
	}

	bot.lastError = fmt.Errorf("UNKNOWN COMMAND %s", action)
	return false
}

// parsePlaceArguments parses all arguments for PLACE cmd and places
// the bot accordingly if correct parameters were provided
func parsePlaceArguments(bot *Bot, argStr string) {
	// Split them with comma delimeter
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

	//
	var dir int
	dirStr := strings.TrimSpace(args[2])
	switch dirStr {
	case "north":
		dir = Up
	case "west":
		dir = Left
	case "south":
		dir = Down
	case "east":
		dir = Right
	default:
		bot.lastError = fmt.Errorf("%s IS NOT A CORRECT DIRECTION (EAST, WEST, NORTH, SOUTH)", dirStr)
		return
	}

	bot.Place(dir, x, y)
}

func sanitizeCmd(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	return strings.ToLower(cmd)
}
