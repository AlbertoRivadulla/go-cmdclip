# Command clip CLI app


## TODO
- [ ] Basic CLI app
    - [ ] Configuration of the app
        - [ ] Flag to tell the app where the database is located (with default in `$HOME/cmdclip/database/`)
        - [ ] 
- [ ] Interface of the app
    - [ ] First version: basic CLI app
    - [ ] Interactive prompt with [manifoldco/promptui](https://github.com/manifoldco/promptui)
    - [ ] Text UI with [bubbletea](https://github.com/charmbracelet/bubbletea)
- [ ] Interaction with the app
    - [ ] Move around with Vim commands
    - [ ] Fill the placeholders in the CLI, and copy the entire command to the clipboard
    - [ ] Fuzzy search ([sahilm/fuzzy](https://github.com/sahilm/fuzzy))
- [ ] Manage database
    - [ ] Support for JSON and YAML files
    - [ ] Support for multiple files (hierarchy of commands in the CLI)
    - [ ] Functionality to add/remove commands
    - [ ] Optional flags and placeholders
