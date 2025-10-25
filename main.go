package main

import (
	"cmdclip/lib"

	"flag"
	"fmt"
	// "io"
	// "log"
)

func main() {
	dbDirPath := flag.String("dbdir", "$HOME/cmdclip/database", "the directory of the database")

	// log.SetOutput(io.Discard)

	flag.Parse()

	var commandSets []lib.CommandSet
	commandSets = lib.LoadCmds(*dbDirPath)
	for _, cmdSet := range commandSets {
		for _, cmd := range cmdSet.Commands {
			cmd.Print()
		}
	}

	// TODO: Initialize the CLI app


	// TODO: Pass the commands to the app


	// TODO: Start app loop
}

