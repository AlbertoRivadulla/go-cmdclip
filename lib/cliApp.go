package lib

import (
	"fmt"
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
	SecondColFlex *tview.Flex

	CmdSetList *tview.List
	CmdSetDescrText *tview.TextView
	CmdList *tview.List
	CmdContentText *tview.TextView

	CurrentCmdSetIdx int
	CurrentCmdIdx int


	// TODO:
}

func (cliApp* CliApp) Initialize(dbDir string) {
	// Initialize variables
	cliApp.StatusCh = make(chan string, 1)
	cliApp.CurrentCmdSetIdx = 0
	cliApp.CurrentCmdIdx = 0

	// Load the commands recursively
	cliApp.CmdSets = loadCmds(dbDir)

	cliApp.App = tview.NewApplication()

	cliApp.setupView()

	// Make the main view the root of the app, and display it fullscreen
	cliApp.App.SetRoot(cliApp.MainView, true)

	// Focus on the panel with the list of command sets
	cliApp.App.SetFocus(cliApp.CmdSetList)
}

func (cliApp* CliApp) setupView() {
	// Initialize the different elements in the app
	cliApp.CmdSetList = tview.NewList().
		ShowSecondaryText(false)
	cliApp.CmdSetList.SetBorder(true).
		SetTitle("Command sets")
	cliApp.CmdSetDescrText = tview.NewTextView().
		SetDynamicColors(true)
	cliApp.CmdSetDescrText.SetBorder(true).
		SetBorderPadding(1, 1, 2, 2)
	cliApp.CmdList = tview.NewList().
		ShowSecondaryText(false)
	cliApp.CmdList.SetBorder(true).
		SetTitle("Commands")
	cliApp.CmdContentText = tview.NewTextView().
		SetDynamicColors(true)

	// Create the second column
	cliApp.SecondColFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(cliApp.CmdSetDescrText, 0, 1, false).
		AddItem(cliApp.CmdList, 0, 3, false)

	// Place the elements in a flex layout
	cliApp.MainView = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(cliApp.CmdSetList, 0, 1, true).
		AddItem(cliApp.SecondColFlex, 0, 1, false).
		AddItem(cliApp.CmdContentText, 0, 2, false)
	
	cliApp.setupCmdSets()
}

func (cliApp* CliApp) setupCmdSets() {
	cliApp.CmdSetList.Clear()
	for _, cmdSet := range cliApp.CmdSets {
		cliApp.CmdSetList.AddItem(cmdSet.Title, "", 0, func() {
			// The function to run when a command set is selected
			cliApp.App.SetFocus(cliApp.CmdList)
		})
	}

	cliApp.CmdSetList.SetChangedFunc(cliApp.setupCmds)

	cliApp.setupCmds(cliApp.CurrentCmdSetIdx, "", "", 0)
}

func (cliApp* CliApp) setupCmds(index int, mainText string, secondaryText string, shortcut rune) {
	cliApp.CurrentCmdSetIdx = index
	cliApp.CmdSetDescrText.SetText(fmt.Sprintf("[::i]%s[::-]", cliApp.CmdSets[index].Description))

	cliApp.CmdList.Clear()
	for _, cmd := range cliApp.CmdSets[index].Commands {
		cliApp.CmdList.AddItem(cmd.Name, "", 0, func() {
			// The function to run when a command is selected
			cliApp.App.SetFocus(cliApp.CmdContentText)
		})
	}

	cliApp.CmdList.SetChangedFunc(cliApp.setupCmdContent)

	cliApp.setupCmdContent(cliApp.CurrentCmdIdx, "", "", 0)
}

func (cliApp* CliApp) setupCmdContent(index int, mainText string, secondaryText string, shortcut rune) {
	cliApp.CurrentCmdIdx = index
	command := cliApp.CmdSets[cliApp.CurrentCmdSetIdx].Commands[index]
	cliApp.CmdContentText.SetText(fmt.Sprintf("%s\n\n%s", command.Description, command.Command))

	// TODO: Improve the formatting of this
}

func (cliApp* CliApp) Run() {
	err := cliApp.App.Run()
	if err != nil {
		log.Fatal("Error in CLI app runtime: ", err.Error())
	}
}
