package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/nicored/nicobot/deco"
	"github.com/nicored/nicobot/nicobot"
)

func main() {
	bot := &nicobot.Bot{}
	renderer := &nicobot.Renderer{}

	args := os.Args[1:]
	if len(args) == 0 {
		renderer.RenderWithPrompt(bot)
		return
	}

	// If arguments were found, then we try to process the file
	processFromFile(args, bot, renderer)
}

func processFromFile(args []string, bot *nicobot.Bot, renderer *nicobot.Renderer) {
	file, err := os.Open(args[0])
	if err != nil {
		deco.Color(deco.BgRed)
		fmt.Println("\033[31m" + err.Error() + "\033[0m")
		deco.Color(deco.DefaultColor)

		os.Exit(1)
	}

	buff := bufio.NewReader(file)

	ticker := time.Tick(500 * time.Millisecond)
	for {
		cmd, err := buff.ReadString('\n')
		if err == io.EOF {
			renderer.RenderWithCmd(bot, cmd)
			break
		} else if err != nil {
			fmt.Println("\033[31m" + err.Error() + "\033[0m")
			os.Exit(1)
		}

		renderer.RenderWithCmd(bot, cmd)
		<-ticker
	}

	os.Exit(0)
}
