package lib

import (
	"log"

	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CliApp struct {
	CmdSets []CommandSet

	// TODO: Make sure this is useful.
	// TODO: Change the name if necessary
	StatusCh chan string

	App *tview.Application
	MainView *tview.Flex

	CmdSetList *tview.List
	CmdList *tview.List
	CmdContentText *tview.TextView


	// TODO:
}

func (cliApp* CliApp) Initialize(dbDir string) {

	cliApp.StatusCh = make(chan string, 1)

	// Load the commands recursively
	cliApp.CmdSets = loadCmds(dbDir)

	cliApp.App = tview.NewApplication()

	cliApp._setupView()

	// Make the main view the root of the app, and display it fullscreen
	cliApp.App.SetRoot(cliApp.MainView, true)

	// Focus on the panel with the list of command sets
	cliApp.App.SetFocus(cliApp.CmdSetList)

	for _, cmdSet := range cliApp.CmdSets {
		for _, cmd := range cmdSet.Commands {
			cmd.Print()
		}
	}

}

func (cliApp* CliApp) _setupView() {
	// Initialize the different elements in the app
	cliApp.CmdSetList = tview.NewList().
		ShowSecondaryText(false)
	cliApp.CmdSetList.SetBorder(true).
		SetTitle("Command sets")
	cliApp.CmdList = tview.NewList().
		ShowSecondaryText(false)
	cliApp.CmdList.SetBorder(true).
		SetTitle("Commands")
	cliApp.CmdContentText = tview.NewTextView().
		SetDynamicColors(true)

	// Place the elements in a flex layout
	cliApp.MainView = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(cliApp.CmdSetList, 0, 1, true).
		AddItem(cliApp.CmdList, 0, 1, false).
		AddItem(cliApp.CmdContentText, 0, 2, false)


	// TODO: Populate the column of commands sets
	// 	Select the first one
	// 	Populate the second column based on it
	// 	or
	// 	Leave the second and third columns empty

}

func (cliApp* CliApp) Run() {
	err := cliApp.App.Run()
	if err != nil {
		log.Fatal("Error in CLI app runtime: ", err.Error())
	}
}
