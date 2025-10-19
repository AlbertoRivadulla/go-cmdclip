package main

import (
	"cmdclip/lib"

	"flag"
)

func main() {
	dbDirPath := flag.String("dbdir", "$HOME/cmdclip/database", "the directory of the database")

	flag.Parse()

	lib.LoadCmds(*dbDirPath)
}

