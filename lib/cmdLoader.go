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

func loadCmds(dbDir string) ([]CommandSet, error) {
	dbDirPath, _ := filepath.Abs(dbDir)

	if _, err := os.Stat(dbDirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: %s", err.Error())
	}

	commandSets, err := _loadCmdsFromDir(dbDirPath)
	if err != nil {
		return nil, err
	}
	if len(commandSets) == 0 {
		return nil, fmt.Errorf("No commands found in the path: %s", dbDir)
	}

	return commandSets, nil
}

func _loadCmdsFromDir(directory string) ([]CommandSet, error) {
	var commandSets []CommandSet

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
			newCmdSet, err := _loadCmdsFromFile(path)
			if err != nil || newCmdSet == nil {
				log.Printf("Could not load commands in %s. Error: %s", path, err.Error())
				return nil
			}
			if len(newCmdSet.Commands) > 0 {
				commandSets = append(commandSets, *newCmdSet)
			}
		}

		return nil
	})

	return commandSets, err
}

func _loadCmdsFromFile(filePath string) (*CommandSet, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var commandSet CommandSet
	switch extension := filepath.Ext(filePath); extension {
	case ".json":
		err = json.Unmarshal(data, &commandSet)
	case ".yaml":
		err = yaml.Unmarshal(data, &commandSet)
	default:
		return nil, nil
	}

	return &commandSet, err
}
