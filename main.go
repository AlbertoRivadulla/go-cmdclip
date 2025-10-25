package main

import (
	"cmdclip/lib"

	"flag"
	// "io"
	// "log"
)

func main() {
	dbDirPath := flag.String("dbdir", "$HOME/cmdclip/database", "the directory of the database")

	// log.SetOutput(io.Discard)

	flag.Parse()

	lib.LoadCmds(*dbDirPath)
}

