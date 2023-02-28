package main

import (
	"github.com/aldernero/ulid-tui/pkg/tui"
	"os"
)

func main() {
	var text string
	if len(os.Args) > 1 {
		text = os.Args[1]
	}
	tui.StartTea(text)
}
