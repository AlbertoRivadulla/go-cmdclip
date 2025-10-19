package lib

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadCmds(dbDir string) {
	dbDirPath, _ := filepath.Abs(dbDir)

	if _, err := os.Stat(dbDirPath); os.IsNotExist(err) {
		log.Fatal("Error: ", err.Error())
	}
	fmt.Printf("Loading commands from %s...\n", dbDirPath)

	// TODO: Read the commands in the directory

	commands, err := _loadCmdsFromDir(dbDirPath)
	if err != nil {
		log.Fatal("Could not load commands: ", err.Error())
	}
	if len(commands) == 0 {
		log.Fatal("No commands found in the given path")
	}

	fmt.Printf("\n\nCommands found:\n\n")

	for _, cmd := range commands {
		cmd.print()
	}
}

func _loadCmdsFromDir(directory string) ([]Command, error) {
	var commands []Command

	// filepath.WalkDir traverses the directory and subdirectories within it, so there is no need to call this function
	// recursively
	err := filepath.WalkDir(directory, func(path string, dirInfo fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dirInfo.IsDir() {
			return nil
		}

		extension := filepath.Ext(path)
		if extension == ".yaml" || extension == ".yml" || extension == ".json" {
			newCmds, err := _loadCmdsFromFile(path)
			if err != nil {
				return fmt.Errorf("Failed to load commands in %s. Error: %s", path, err.Error())
			}
			if len(newCmds) > 0 {
				commands = append(commands, newCmds...)
			}
		}

		return nil
	})

	return commands, err
}

func _loadCmdsFromFile(filePath string) ([]Command, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var commands []Command
	switch extension := filepath.Ext(filePath); extension {
	case ".json":
		err = json.Unmarshal(data, &commands)
	case ".yaml":
		err = yaml.Unmarshal(data, &commands)
	default:
		return nil, nil
	}

	return commands, err
}
