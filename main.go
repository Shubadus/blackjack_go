package main

import (
	"os"

	"github.com/Shubadus/blackjack_go/internal/gui"
)

func run() int {
	g := gui.New()
	if err := g.Start(); err != nil {
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
