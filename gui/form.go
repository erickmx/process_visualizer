package gui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"

	"github.com/erickmx/process_visualizer/utils"
)

// MakeForm returns the form structure to create or edit a new form component
func MakeForm(formData *utils.FakeProcess, window *fyne.Window) (form *widget.Form) {
	pidStr := string(formData.Pid)
	form = &widget.Form{Items: []*widget.FormItem{
		MakeFormInput("Ingrese Nombre", "Nombre", &formData.Name, func(value string) {
			formData.Name = value
		}),
		MakeFormInput("Ingrese el id", "ID", &pidStr, func(value string) {
			intData, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				dialog.ShowError(err, *window)
				fmt.Println("error", err)
			}
			fmt.Println(intData)
			formData.Pid = int32(intData)
		}),
	}, OnSubmit: func() {
		fmt.Println(formData)
		formData.Name = ""
		formData.Pid = 0
	}}
	return form
}

// MakeFormInput creates a form input to manipulate it easily
func MakeFormInput(text string, placeHolder string, value *string, onChanged func(string)) *widget.FormItem {
	return &widget.FormItem{
		Text: text,
		Widget: &widget.Entry{
			PlaceHolder: placeHolder,
			Text:        *value,
			OnChanged:   onChanged,
		},
	}
}
