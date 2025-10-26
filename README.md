# Command clip CLI app


## TODO
- [x] Basic CLI app
    - [x] Configuration of the app
        - [x] Flag `--dbdir` to tell the app where the database is located (with default in `$HOME/cmdclip/database/`)
- [x] Interface of the app
    - [x] Interactive prompt with [tview](https://github.com/rivo/tview)
    - [x] Add a status line at the bottom
    - [x] Display help at the bottom
- [/] Interaction with the app
    - [x] Move around with Vim key bindings (`hjkl`) 
    - [/] Copy the selected command to the clipboard
        - [x] Copy it with `y`, when focused on the list of commands or the command description
        - [x] Copy it with `Enter` when focused on the command description
        - [ ] If the command has placeholder values, show an error in the status box
    - [ ] Run the command directly, pressing `Enter` when focused on the command description component
    - [ ] Fill the placeholders in the CLI, and copy the entire command to the clipboard
    - [ ] Fuzzy search with `/` ([sahilm/fuzzy](https://github.com/sahilm/fuzzy))
- [/] Manage database
    - [x] Support for JSON and YAML files
        - [x] JSON
        - [x] YAML
    - [x] Title and description fields for the different command sets
    - [/] Support for multiple files (hierarchy of commands in the CLI)
        - [x] Read commands from all valid files in the directory
        - [ ] Write the commands in a hierarchical data structure
    - [ ] Optional flags and placeholders
- [ ] Future improvements
    - [ ] Functionality to add/remove commands from the CLI
    - [ ] Switch the interface library to [Bubbletea](https://github.com/charmbracelet/bubbletea)
