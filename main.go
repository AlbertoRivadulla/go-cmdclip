package main

import (
	"cmdclip/lib"

	"flag"
	// "fmt"
	// "io"
	// "log"

	// "github.com/rivo/tview"
)

func main() {
	dbDirPath := flag.String("dbdir", "$HOME/cmdclip/database", "the directory of the database")
	flag.Parse()

	// log.SetOutput(io.Discard)

	var cliApp lib.CliApp
	cliApp.Initialize(*dbDirPath)

	cliApp.Run()
}
