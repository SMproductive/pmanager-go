package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/* test things */
var data = []string{"tilkjasdlklkasdljkaklsjdftle1 ", "title2 ", "title3 "}
/* test things end */

func main() {

	a := app.New()
	w := a.NewWindow("Pmanager")

	passwordEntryTest := widget.NewPasswordEntry() /* For testing */
	passwordEntryTest.Text = "your password"
	passwordEntryTest.Disable()
	btnAddTitle := widget.NewButton("Add", addTitle)
	btnChangeTitle := widget.NewButton("Change", changeTitle)
	btnRemoveTitle := widget.NewButton("Remove", removeTitle)
	//topSpacer := layout.NewSpacer()
	//topSpacer1 := layout.NewSpacer()
	titles := widget.NewList(numTitles, createTitle, updateTitle)

	topLeftBox := container.NewHBox(btnAddTitle, btnChangeTitle, btnRemoveTitle)

	topSplit := container.NewHSplit(topLeftBox, titles)
	mainSplit := container.NewVSplit(topSplit, passwordEntryTest)
	w.SetContent(mainSplit)

	w.ShowAndRun()
}

func addTitle() {

}

func changeTitle() {

}

func removeTitle() {

}

func numTitles() int {
	return len(data)
}

func createTitle() fyne.CanvasObject {
	return widget.NewEntry()
}

func updateTitle(id widget.ListItemID, obj fyne.CanvasObject) {
	obj.(*widget.Entry).Text = data[id]
	obj.(*widget.Entry).Disable()
}
