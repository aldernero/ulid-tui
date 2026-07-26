package main

import (
	"os"

	"github.com/aldernero/ulid-tui/pkg/tui"
)

func main() {
	var text string
	if len(os.Args) > 1 {
		text = os.Args[1]
	}
	tui.StartTea(text)
}
