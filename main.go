package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/erickmx/process_visualizer/gui"
	"github.com/erickmx/process_visualizer/utils"
)

var fakeProcecess []*utils.FakeProcess

func init() {
	if !utils.FileExists("./procecess.csv") {
		fmt.Println("Filling procecess information")
		procecess := utils.GetProcecess()
		okay := utils.SaveProcecess(procecess)
		if okay {
			fmt.Println("Procecess saved successfully")
		}
	}
	fakeProcecess = utils.ReadProcecess("./procecess.csv")
}

func main() {
	headings := []string{
		"ID",
		"Nombre",
		"Estado",
		"Usuario",
		"CPU",
		"Memoria",
		"Prioridad",
	}
	var rows [][]string
	for _, fakeProcess := range fakeProcecess {
		rows = append(rows, fakeProcess.ToArray())
	}

	a := app.New()
	w := a.NewWindow("Administrador de Procesos")

	w.Resize(fyne.NewSize(1200, 600))
	table := gui.MakeTable(headings, rows)
	scrollableTable := &widget.ScrollContainer{
		Content: table,
	}
	scrollableTable.Resize(fyne.NewSize(1200, 600))
	/*
		box := &widget.Box{
			Children: []fyne.CanvasObject{
				scrollableTable,
			},
		}
		box.Resize(fyne.NewSize(1200, 600))
	*/
	w.SetContent(scrollableTable)
	w.SetFullScreen(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}

/*
// Version 1 (functions)
package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	w.ShowAndRun()
}
*/

/*
package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/erickmx/process_visualizer/gui"
	"github.com/erickmx/process_visualizer/utils"
)

var fakeProcecess []*utils.FakeProcess

func init() {
	if !utils.FileExists("./procecess.csv") {
		fmt.Println("Filling procecess information")
		procecess := utils.GetProcecess()
		okay := utils.SaveProcecess(procecess)
		if okay {
			fmt.Println("Procecess saved successfully")
		}
	}
	fakeProcecess = utils.ReadProcecess("./procecess.csv")
}

func main() {
	headings := []string{
		"ID",
		"Nombre",
		"Estado",
		"Usuario",
		"CPU",
		"Memoria",
		"Prioridad",
	}
	var rows [][]string
	for _, fakeProcess := range fakeProcecess {
		rows = append(rows, fakeProcess.ToArray())
	}

	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(
		&widget.Box{Children: []fyne.CanvasObject{
			&widget.ScrollContainer{
				Content: gui.MakeTable(headings, rows),
			},
		}})
	w.Resize(fyne.NewSize(900, 900))
	w.ShowAndRun()
}

*/
