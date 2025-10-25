package lib

import (
	"fmt"
)

type Command struct {
	Name string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Command string `json:"command" yaml:"command"`
}

type CommandSet struct {
	Title string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Commands []Command `json:"commands" yaml:"commands"`
}

func (command* Command) print() {
	fmt.Printf("Name: %s\nDescription: %s\nCommand: %s\n\n", 
		command.Name, 
		command.Description, 
		command.Command,
	)
}
