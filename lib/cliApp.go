package lib

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CliApp struct {
	CmdSets []CommandSet

	// TODO: Make sure this is useful.
	// TODO: Change the name if necessary
	StatusCh chan string

	App *tview.Application
	MainView *tview.Grid

	CmdSetListTable *tview.Table
	CmdListTable *tview.Table
	CmdContentTable *tview.Table


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
	cliApp.App.SetFocus(cliApp.CmdSetListTable)

	for _, cmdSet := range cliApp.CmdSets {
		for _, cmd := range cmdSet.Commands {
			cmd.Print()
		}
	}

}

func (cliApp* CliApp) _setupView() {
	// Initialize the grid layout
	cliApp.MainView = tview.NewGrid().
		// SetBorders(true).
		// SetColumns(40, 40, -1). // Two columns with size 40 chars, the third column occupies the remaining space
		SetColumns(-1, -1, -2). // Proportional size. The third row will be twice as large as the other two
		SetRows(0).
		SetGap(1, 1)
	

	// Initialize the different tables in the app
	cliApp.CmdSetListTable = tview.NewTable()
	cliApp.CmdListTable = tview.NewTable()
	cliApp.CmdContentTable = tview.NewTable()
	// cliApp.CmdContentTable = tview.NewTable().
	// 	SetBorders(true)

	// Wrap the first two columns in a frame, and add padding and text in their headers
	cmdSetListFrame := tview.NewFrame(cliApp.CmdSetListTable).
		SetBorders(0, 0, 0, 0, 1, 1).
		AddText("[::b]Command sets[::-]", true, tview.AlignCenter, tcell.ColorGreen)
	cmdSetListFrame.SetBorder(true).
		SetBorderColor(tcell.ColorWhite)
	cmdListFrame := tview.NewFrame(cliApp.CmdListTable).
		SetBorders(0, 0, 0, 0, 1, 1).
		AddText("[::b]Commands[::-]", true, tview.AlignCenter, tcell.ColorGreen)
	cmdListFrame.SetBorder(true).
		SetBorderColor(tcell.ColorWhite)


	// // Wrap each in a Frame
	// commandsFrame := tview.NewFrame(commandsTable).
	// 	SetBorders(0, 0, 0, 0, 1, 1). // internal padding (top,bottom,left,right,hPad,vPad)
	// 	AddText("Commands", true, tview.AlignCenter, tcell.ColorGreen)
	//
	// detailsFrame := tview.NewFrame(detailsTable).
	// 	SetBorders(0, 0, 0, 0, 1, 1).
	// 	AddText("Details", true, tview.AlignCenter, tcell.ColorGreen)
	//



	// cliApp.MainView.AddItem(cliApp.CmdSetListTable, 0, 0, 1, 1, 1, 1, true)
	// cliApp.MainView.AddItem(cliApp.CmdListTable, 0, 1, 1, 1, 1, 1, false)
	cliApp.MainView.AddItem(cmdSetListFrame, 0, 0, 1, 1, 1, 1, false)
	cliApp.MainView.AddItem(cmdListFrame, 0, 1, 1, 1, 1, 1, false)
	cliApp.MainView.AddItem(cliApp.CmdContentTable, 0, 2, 1, 1, 1, 1, false)







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
