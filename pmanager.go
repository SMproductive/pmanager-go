package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SMproductive/pmanager-go/customWidget"
)

var a fyne.App = app.New()
var windowLogin fyne.Window = a.NewWindow("Pmanager")
var windowMain fyne.Window =  a.NewWindow("Pmanager")
/* test things */
var data = []string{"tilkjasdlklkasdljkaklsjdftle1 ", "title2 ", "title3 "}
/* test things end */

func main() {

	/* Login window */
	lblLogin := widget.NewLabel("Master Password:")
	entryLogin := widget.NewPasswordEntry()
	entryLogin.OnSubmitted = login
	gridLogin := layout.NewGridLayout(2)
	containerLogin := fyne.NewContainerWithLayout(gridLogin, lblLogin, entryLogin)
	windowLogin.SetContent(containerLogin)

	windowLogin.ShowAndRun()
}

func login(password string) {
	windowLogin.Close()
	/* testing */
	foo := customWidget.NewEntry()
	foo.Text = "foo"
	/* testing end */
	btnAddTitle := widget.NewButton("Add", addTitle)
	btnSave := widget.NewButton("Save", save)

	listTitles := widget.NewList(numTitles, createTitle, updateTitle)
	listTitles.OnSelected = selectedTitle

	contentGrid := layout.NewGridLayout(2)
	contentContainer := fyne.NewContainerWithLayout(contentGrid, foo)

	topLeftBox := container.NewHBox(btnAddTitle, btnSave)
	topSplit := container.NewHSplit(topLeftBox, listTitles)
	mainSplit := container.NewVSplit(topSplit, contentContainer)
	windowMain.SetContent(mainSplit)
	go windowMain.Show()
}

func addTitle() {

}

func save() {

}


func numTitles() int {
	return len(data)
}
func createTitle() fyne.CanvasObject {
	return widget.NewLabel("foo")
}
func updateTitle(id widget.ListItemID, obj fyne.CanvasObject) {
	obj.(*widget.Label).Alignment = 1 /* Aligned to center */
	obj.(*widget.Label).SetText(data[id])

}
func selectedTitle(id widget.ListItemID) {
	fmt.Println(id)
}
