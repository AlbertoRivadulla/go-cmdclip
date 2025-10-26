package lib

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
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

	// Initialize the clipboard handler
	err := clipboard.Init()
	if err != nil {
		log.Fatal("Error initializing the clipboard handler: %s", err.Error())
		  panic(err)
	}

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
	cliApp.CmdContentText.SetBorder(true).
		SetBorderPadding(5, 5, 2, 2)

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

	cliApp.setupInputHandling()
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
	cliApp.CmdContentText.SetText(fmt.Sprintf("[::i]%s[::-]\n\n\n\n%s", command.Description, command.Command))

	// TODO: Improve the formatting of this
}

func (cliApp* CliApp) setupInputHandling() {
	
	// Global navigation
	cliApp.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' || event.Key() == tcell.KeyCtrlC {
			cliApp.App.Stop()
			return nil
		}
		return event
	})

	// Navigation inside the lists
	handleListNavigation := func(event *tcell.EventKey, list *tview.List, currIdx *int) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			index := list.GetCurrentItem() + 1
			if index < 0 {
				index = list.GetItemCount() - 1
			}
			if index == list.GetItemCount() {
				index = 0
			}
			list.SetCurrentItem(index)
			*currIdx = index
		case 'k':
			index := list.GetCurrentItem() - 1
			if index < 0 {
				index = list.GetItemCount() - 1
			}
			if index == list.GetItemCount() {
				index = 0
			}
			list.SetCurrentItem(index)
			*currIdx = index
		case 'l': // Move to the right
			if list == cliApp.CmdSetList {
				cliApp.App.SetFocus(cliApp.CmdList)
				return nil
			} else if list == cliApp.CmdList {
				cliApp.App.SetFocus(cliApp.CmdContentText)
				return nil
			}
		case 'h': // Move to the left
			if list == cliApp.CmdList {
				cliApp.App.SetFocus(cliApp.CmdSetList)
				return nil
			}
		case 'y':
			if list == cliApp.CmdList {
				cliApp.copyCurrentCmdToClipboard()
			}
		}

		if event.Key() == tcell.KeyEsc {
			if list == cliApp.CmdSetList {
				cliApp.App.Stop()
				return nil
			} else if list == cliApp.CmdList {
				cliApp.App.SetFocus(cliApp.CmdSetList)
				return nil
			}
		}

		return event
	}
	cliApp.CmdSetList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return handleListNavigation(event, cliApp.CmdSetList, &cliApp.CurrentCmdSetIdx)
	})
	cliApp.CmdList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return handleListNavigation(event, cliApp.CmdList, &cliApp.CurrentCmdIdx)
	})

	// Input in the command content TextView
	cliApp.CmdContentText.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'h':
			cliApp.App.SetFocus(cliApp.CmdList)
		case 'y':
			cliApp.copyCurrentCmdToClipboard()
		}

		if event.Key() == tcell.KeyEsc {
			cliApp.App.SetFocus(cliApp.CmdList)
		}

		if event.Key() == tcell.KeyEnter {
			cliApp.copyCurrentCmdToClipboard()
		}

		return event
	})
}

func (cliApp* CliApp) copyCurrentCmdToClipboard() {
	// TODO: copy the selected command to the clipboard and close the app
	// This should only work if the command has no placeholder fields
	// Otherwise show a message in the status line
	cmdText := cliApp.CmdSets[cliApp.CurrentCmdSetIdx].Commands[cliApp.CurrentCmdIdx].Command
	clipboard.Write(clipboard.FmtText, []byte(cmdText))
	cliApp.App.Stop()
}

func (cliApp* CliApp) Run() {
	err := cliApp.App.Run()
	if err != nil {
		log.Fatal("Error in CLI app runtime %s: ", err.Error())
	}
}
