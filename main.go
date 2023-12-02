package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/chodyo/advent-2023/internal/commands/trebuchet"
)

type Commands struct {
	// Day 1
	Trebuchet trebuchet.Command `command:"trebuchet" description:"Advent 2023 Day 1"`
}

// Commands are executed in `ParseArgs`
// https://pkg.go.dev/github.com/jessevdk/go-flags?utm_source=godoc#hdr-Commands
func main() {
	// Disable logs timestamp
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	commands := Commands{}
	parser := flags.NewParser(&commands, flags.Default)

	if _, err := parser.ParseArgs(os.Args[1:]); err != nil {
		if flags.WroteHelp(err) {
			return
		}
		os.Exit(1)
	}
}
