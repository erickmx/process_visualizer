package gui

import (
		"fyne.io/fyne/widget"
)

// MakeScrollableTable use the MakeTable function but wraps the element into a ScrollContainer to avoid the overflow
func MakeScrollableTable(headings []string, rows[][]string) *widget.ScrollContainer {
	table := MakeTable(headings, rows)
	return &widget.ScrollContainer{
		Content: table,
	}
}
