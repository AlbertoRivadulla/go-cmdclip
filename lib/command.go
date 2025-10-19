package lib

import (
	"fmt"
)

type Command struct {
	Name string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Command string `json:"command" yaml:"command"`
}

func (command* Command) print() {
	fmt.Printf("Name: %s\nDescription: %s\nCommand: %s\n\n", 
		command.Name, 
		command.Description, 
		command.Command,
	)
}
