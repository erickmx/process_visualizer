package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/erickmx/process_visualizer/gui"
	"github.com/erickmx/process_visualizer/parallel"
	"github.com/erickmx/process_visualizer/utils"
)

var fakeProcecess []*utils.FakeProcess

func init() {
	os.Setenv("FYNE_THEME", "light")
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
	formData := &utils.FakeProcess{}
	w := initGui(&a, headings, rows, formData)
	go runProcecess(fakeProcecess)
	time.Sleep(1 * time.Second)
	for _, r := range fakeProcecess {
		if r.Status == "R" {
			fmt.Println(r)
		}
	}
	(*w).ShowAndRun()
}

func initGui(a *fyne.App, headings []string, rows [][]string, formData *utils.FakeProcess) *fyne.Window {
	w := (*a).NewWindow("Administrador de Procesos")
	w.Resize(fyne.NewSize(1200, 600))
	scrollableTable := gui.MakeScrollableTable(headings, rows)
	scrollableTable.Resize(fyne.NewSize(1200, 400))
	form := gui.MakeForm(formData, &w)
	form.Resize(fyne.NewSize(1200, 200))
	form.OnSubmit = func() {
		fmt.Println("Hello")
	}
	vBox := widget.NewVBox(form, scrollableTable)
	vBox.Resize(fyne.NewSize(1200, 600))
	w.SetContent(vBox)
	w.SetFullScreen(true)
	w.CenterOnScreen()
	return &w
}

func runProcecess(proceces []*utils.FakeProcess) {
	fmt.Println("Run procecess")
	var runningID int32
	procecesLen := int32(len(proceces))
	runningID = proceces[rand.Int31n(procecesLen)].Pid
	for _, fakeProcess := range proceces {
		go parallel.ParallelizeProcess(fakeProcess, runningID)
	}
}
