package main
/* TODO
* data encryption
* cloud implementation (google, nextcloud)
* keyboard shortcuts
*/

import (
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SMproductive/pmanager-go/customWidget"
	"github.com/SMproductive/pmanager-go/customTheme"
)

var a fyne.App = app.New()
var windowLogin fyne.Window = a.NewWindow("PmanagerLogin")
var windowMain fyne.Window =  a.NewWindow("Pmanager")
var dataPath binding.String = binding.NewString()
var data = map[string] []string{}
var dataID []string
var containerTitles = &fyne.Container{}

func main() {
	a.Settings().SetTheme(customTheme.Nord{})
	/* Login window */
	logo := canvas.NewImageFromFile("/home/max/projects/logo/icon.png")
	logo.FillMode = canvas.ImageFillContain
	poweredBy := canvas.NewImageFromFile("/home/max/projects/logo/poweredBy.png")
	poweredBy.FillMode = canvas.ImageFillContain

	lblDatabase := widget.NewLabel("Database: ")
	entryDatabase := widget.NewEntryWithData(dataPath)
	home, _ := os.UserHomeDir()
	home += "/.pmanager/passwords"
	dataPath.Set(home)

	lblMasterPass := widget.NewLabel("Master Password:")
	entryMasterPass := widget.NewPasswordEntry()
	entryMasterPass.OnSubmitted = loggedIn
	gridLogin := layout.NewGridLayout(2)
	containerLogin := fyne.NewContainerWithLayout(gridLogin, lblDatabase, entryDatabase, lblMasterPass, entryMasterPass, logo, poweredBy)

	windowLogin.Resize(fyne.NewSize(600, 400))
	windowLogin.CenterOnScreen()
	windowLogin.SetContent(containerLogin)
	windowLogin.ShowAndRun()
}

func loggedIn(password string) {
	windowLogin.Close()

	/* data keys in slice are used as buffer when updating titles */
	for k := range data {
		dataID = append(dataID, k)
	}

	containerTitles = container.NewGridWithRows(1)
	scrollTitles := container.NewHScroll(containerTitles)
	/* Build all titles */
	for i := len(data); i > 0; i-- {
		ent := customWidget.NewTitleEntry()
		ent.OnSubmitted = func(string) {
			ent.Submitted(data, dataID)
		}
		ent.Text = dataID[i-1]
		ent.ID = dataID[i-1]
		containerTitles.Add(ent)
		ent.Refresh()
	}


	btnAddTitle := widget.NewButton("Add", addTitle)
	btnSave := widget.NewButton("Save", save)

	contentGrid := layout.NewGridLayout(2)
	contentContainer := fyne.NewContainerWithLayout(contentGrid)

	customWidget.SendTitle = make(chan string)
	go buildContent(customWidget.SendTitle, contentContainer)


	topLeftBox := container.NewHBox(btnAddTitle, btnSave)
	topSplit := container.NewHSplit(topLeftBox, scrollTitles)
	topSplit.Offset = 0
	mainSplit := container.NewVSplit(topSplit, contentContainer)
	mainSplit.Offset = 0.06


	windowMain.CenterOnScreen()
	windowMain.Resize(fyne.NewSize(1600, 900))
	windowMain.SetContent(mainSplit)
	go windowMain.Show()
}

func addTitle() {
	str := "new"
	data[str] = append(data[str] , str, str)
	dataID = append(dataID, str)
	ent := customWidget.NewTitleEntry()
	ent.OnSubmitted = func(string) {
		ent.Submitted(data, dataID)
	}
	ent.Text = str
	ent.ID = str
	containerTitles.Add(ent)
	containerTitles.Refresh()

}

func save() {

}

func buildContent(chosenTitle <-chan string, con *fyne.Container) {
	for /* true */{
		title, ok := <-chosenTitle
		if ok { /* remove old widgets */
			for i := len(con.Objects); i > 0; i-- {
				con.Remove(con.Objects[i-1])
			}
			for i, v := range data[title] {
				ent := customWidget.NewContentEntry()
				ent.ID = &data[title][i]
				ent.Text = v
					ent.Password = i % 2 == 1
				ent.OnSubmitted = func(string) {
					ent.Submitted()
				}
				con.Add(ent)
			}
			btnAdd := widget.NewButton("Add", func() {
				data[title] = append(data[title], "new", "new")
				customWidget.SendTitle <- title
			})
			con.Add(btnAdd)
			con.Refresh()
		}
	}
}
