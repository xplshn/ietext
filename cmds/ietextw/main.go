package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rivo/tview"
)

func main() {
	var placeholder string
	flag.StringVar(&placeholder, "p", "Enter text here...", "Set placeholder text")
	flag.Parse()

	app := tview.NewApplication()

	textArea := tview.NewTextArea().
		SetWrap(false).
		SetPlaceholder(placeholder)
	textArea.SetTitle("Text Area").SetBorder(true)
	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)
	pages := tview.NewPages()

	// Load file content if provided
	if flag.NArg() > 0 {
		filePath := flag.Arg(0)
		content, err := ioutil.ReadFile(filePath)
		if err == nil {
			textArea.SetText(string(content), false)
		} else {
			fmt.Println("Error reading file:", err)
		}
	} else {
		// Read from STDIN
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// Data is being piped to stdin
			inputBytes, _ := ioutil.ReadAll(os.Stdin)
			textArea.SetText(strings.TrimSpace(string(inputBytes)), false)
		}
	}

	mainView := tview.NewGrid().
		SetRows(0, 1).
		AddItem(textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(position, 1, 1, 1, 1, 0, 0, false)
	pages.AddPage("main", mainView, true, true)

	// Create an event function
	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}
	// Handle events
	textArea.SetChangedFunc(updateInfos)
	updateInfos()

	app.SetRoot(pages, true).EnableMouse(true)

	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output the content of the text area
	fmt.Println(textArea.GetText())
}
