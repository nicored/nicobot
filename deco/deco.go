package deco

import "fmt"

const (
	DefaultColor = "\033[0m"
	TextRed      = "\033[31m"
	BgBlack      = "\033[107;40m"
	BgRed        = "\033[101m"
)

func Color(color string) {
	fmt.Printf(color)
}

// clear clears lines number printed rows from the console
func Clear(lines int) {
	for l := 0; l < lines; l++ {
		// cursor up
		fmt.Print("\033[0A")

		// clear line
		fmt.Print("\033[2K\r")
	}
}
