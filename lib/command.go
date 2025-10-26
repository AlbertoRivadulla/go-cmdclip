package lib

import (
	"fmt"
)

type Placeholder struct {
	Name string
	BeginIdx int
}

type Command struct {
	Name string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Command string `json:"command" yaml:"command"`
	Placeholders []Placeholder
}

type CommandSet struct {
	Title string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Commands []Command `json:"commands" yaml:"commands"`
}

func (command* Command) Print() {
	fmt.Printf("Name: %s\nDescription: %s\nCommand: %s\n\n", 
		command.Name, 
		command.Description, 
		command.Command,
	)
}
