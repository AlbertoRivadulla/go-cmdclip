# Command clip CLI app


## TODO
- [/] Basic CLI app
    - [/] Configuration of the app
        - [x] Flag `--dbdir` to tell the app where the database is located (with default in `$HOME/cmdclip/database/`)
- [/] Interface of the app
    - [ ] ~~Interactive prompt with [manifoldco/promptui](https://github.com/manifoldco/promptui)~~
    - [/] Interactive prompt with [tview](https://github.com/rivo/tview)
    - [ ] Text UI with [Bubbletea](https://github.com/charmbracelet/bubbletea)
    - [ ] Try [Cobra](https://github.com/spf13/cobra)
- [ ] Interaction with the app
    - [ ] Move around with Vim commands
    - [ ] Fill the placeholders in the CLI, and copy the entire command to the clipboard
    - [ ] Fuzzy search ([sahilm/fuzzy](https://github.com/sahilm/fuzzy))
    - [ ] Display help at the bottom (this requires creating another flex in the vertical direction, that will contain the current one as the first element and a text view below it)
- [/] Manage database
    - [x] Support for JSON and YAML files
        - [x] JSON
        - [x] YAML
    - [x] Title and description fields for the different command sets
    - [/] Support for multiple files (hierarchy of commands in the CLI)
        - [x] Read commands from all valid files in the directory
        - [ ] Write the commands in a hierarchical data structure
    - [ ] Functionality to add/remove commands
    - [ ] Optional flags and placeholders
