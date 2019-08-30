package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// MakeTable Creates a table widget to show the corresponding information
func MakeTable(headings []string, rows [][]string) *widget.Box {
	columns := RowsToColumns(headings, rows)

	objects := make([]fyne.CanvasObject, len(columns))
	for k, col := range columns {
		box := widget.NewVBox(widget.NewLabelWithStyle(headings[k], fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		for _, val := range col {
			box.Append(widget.NewLabel(val))
		}
		objects[k] = box
	}
	return widget.NewVBox(
		fyne.NewContainerWithLayout(layout.NewGridLayout(len(columns)), objects...),
	)
}

// RowsToColumns converts the rows to the coresponding columns
func RowsToColumns(headings []string, rows [][]string) [][]string {
	columns := make([][]string, len(headings))
	for _, row := range rows {
		for colK := range row {
			columns[colK] = append(columns[colK], row[colK])
		}
	}
	return columns
}
