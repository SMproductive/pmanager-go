package main

import (
	//"fmt"
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SMproductive/pmanager-go/customWidgets"
)

var a fyne.App = app.New()
var windowLogin fyne.Window = a.NewWindow("Pmanager")
var windowMain fyne.Window =  a.NewWindow("Pmanager")
var dataBase binding.String = binding.NewString()
/* test things */
var data = []string{"one", "two", "three"}
var bar []binding.String = make([]binding.String, 3)
/* test things end */

func main() {
	/* Login window */
	lblDatabase := widget.NewLabel("Database: ")
	entryDatabase := widget.NewEntryWithData(dataBase)
	home, _ := os.UserHomeDir()
	home += "/.pmanager/passwords"
	dataBase.Set(home)
	lblMasterPass := widget.NewLabel("Master Password:")
	entryMasterPass := widget.NewPasswordEntry()
	entryMasterPass.OnSubmitted = login
	gridLogin := layout.NewGridLayout(2)
	containerLogin := fyne.NewContainerWithLayout(gridLogin, lblDatabase, entryDatabase, lblMasterPass, entryMasterPass)
	windowLogin.SetContent(containerLogin)
	windowLogin.ShowAndRun()
}

func login(password string) {
	windowLogin.Close()
	/* testing */
	for i := range bar {
		bar[i] = binding.NewString()
		bar[i].Set(data[i])
	}
	/*foo := customWidget.NewEntry()
	foo1 := customWidget.NewEntry()
	foo2 := customWidget.NewEntry()*/
	/* testing end */
	btnAddTitle := widget.NewButton("Add", addTitle)
	btnSave := widget.NewButton("Save", save)

	listTitles := widget.NewList(numTitles, createTitle, updateTitle)
	listTitles.OnSelected = selectedTitle

	contentGrid := layout.NewGridLayout(2)
	contentContainer := fyne.NewContainerWithLayout(contentGrid)

	/* testing */
	customWidget.SendTitle = make(chan string)
	go buildContent(customWidget.SendTitle, contentContainer)
	/* testing end */

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
	return customWidget.NewTitleEntry()
}
func updateTitle(id widget.ListItemID, obj fyne.CanvasObject) {
	/*obj.(*customWidget.Label).Alignment = 1 /* Aligned to center */
	if !obj.(*customWidget.TitleEntry).IsBound {
		obj.(*customWidget.TitleEntry).BindStr(bar[id])
	}
}
func selectedTitle(id widget.ListItemID) {
}

func buildContent(chosenTitle <-chan string, con *fyne.Container) {
	/* remove old widgets */
	for {
		title, ok := <-chosenTitle
		if ok {
			for _, v := range con.Objects {
				con.Remove(v)//con.Objects[i])
			}
		}
		/* build wanted content */
		switch title {
		case "one":
			lab := customWidget.NewContentEntry()
			lab.SetText("hello")
			con.Add(lab)
		case "two":
			lab := customWidget.NewContentEntry()
			lab.SetText("hello2")
			con.Add(lab)
		case "three":
			lab := customWidget.NewContentEntry()
			lab.SetText("hello3")
			con.Add(lab)
		}
		con.Refresh()
	}
}
