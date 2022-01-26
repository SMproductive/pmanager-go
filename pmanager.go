package main
/* TODO
* cloud implementation (google, nextcloud)
* keyboard shortcuts
*/
/* FIXME
* crashes when multi tapping save a couple of times
*/

import (
	"crypto/sha256"
	"crypto/rand"
	"encoding/json"
	"runtime"
	"io/ioutil"
	mrand "math/rand"
	"os"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SMproductive/pmanager-go/customWidget"
	"github.com/SMproductive/pmanager-go/customTheme"
)

var a fyne.App = app.New()
var windowLogin fyne.Window = a.NewWindow("PmanagerLogin")
var windowContent fyne.Window =  a.NewWindow("Pmanager")


var key [32]byte
var dataPath string

var data = map[string] []string{}
var dataID []string

var containerTitles = &fyne.Container{}

const website string = "https://github.com/SMproductive/pmanager-go"
const standardPath string = "/.pmanager"

func main() {
	runtime.GOMAXPROCS(10)
	a.Settings().SetTheme(customTheme.Nord{})
	go login()
	windowLogin.ShowAndRun()
}

func login() {
	lblDatabase := widget.NewLabel("Database: ")

	entryDatabase := widget.NewEntry()
	entryDatabase.OnSubmitted = func (path string) {
		dataPath = path
	}
	home, _ := os.UserHomeDir()
	folder := home + standardPath
	dataPath = home + standardPath + "/passwords"
	entryDatabase.Text = home

	logoPath := folder + "/logo.png"
	logoRes, err := fyne.LoadResourceFromPath(logoPath)
	if err != nil {
		logoRes = nil
	}
	logo := customWidget.NewIcon(logoRes, website)

	builtByPath := folder + "/builtBy.png"
	builtByRes, err := fyne.LoadResourceFromPath(builtByPath)
	if err != nil {
		builtByRes = nil
	}
	builtBy := customWidget.NewIcon(builtByRes, website)

	lblMasterPass := widget.NewLabel("Master Password:")
	entryMasterPass := widget.NewPasswordEntry()
	entryMasterPass.OnSubmitted = content /* shows decrypted database */

	gridLogin := layout.NewGridLayout(2)
	containerLogin := fyne.NewContainerWithLayout(gridLogin, lblDatabase, entryDatabase, lblMasterPass, entryMasterPass, logo, builtBy)

	windowLogin.Resize(fyne.NewSize(600, 400))
	windowLogin.CenterOnScreen()
	windowLogin.SetContent(containerLogin)
}

func content(password string) {
	/* Decryption of database */
	key = sha256.Sum256([]byte(password))
	cipherText, err := ioutil.ReadFile(dataPath)
	if err == nil {
		jsonData, err := dec(key[:], cipherText)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			return
		}
	}
	windowLogin.Hide()
	defer windowLogin.Close()

	/* data keys in slice are used as buffer when updating titles */
	buildDataID()

	containerTitles = container.NewGridWithRows(1)
	scrollTitles := container.NewHScroll(containerTitles)
	buildTitles()

	btnAddTitle := widget.NewButton("Add", addTitle)
	btnSave := widget.NewButton("Save", save)

	btnChangePassword := widget.NewButton("Change MPW", changeMasterPass)

	contentGrid := layout.NewGridLayout(2)
	contentContainer := fyne.NewContainerWithLayout(contentGrid)

	customWidget.SendTitle = make(chan string)
	go buildContent(customWidget.SendTitle, contentContainer)


	topLeftBox := container.NewHBox(btnAddTitle, btnSave, btnChangePassword)
	topSplit := container.NewHSplit(topLeftBox, scrollTitles)
	topSplit.Offset = 0
	mainSplit := container.NewVSplit(topSplit, contentContainer)
	mainSplit.Offset = 0.06


	windowContent.CenterOnScreen()
	windowContent.Resize(fyne.NewSize(1600, 900))
	windowContent.SetContent(mainSplit)
	go windowContent.Show()
}

func addTitle() {
	str := "new"
	if data[str] == nil {
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
}

func save() {
	/* clean data */
	for k := range data {
		offset := 0
		for i := range data[k]{
			if data[k][i] == "" {
				offset++
				if i+1 == len(data[k]) {
					break
				}
				data[k][i] = data[k][i+1]
			}
		}
		data[k] = data[k][:len(data[k])-offset]
	}
	delete(data, "")

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	cipherText, err := enc(key[:], jsonData)
	if err != nil {
		panic(err)
	}
	dir, _ := os.UserHomeDir()
	dir +=  "/.pmanager"
	os.Mkdir(dir, 0660)
	ioutil.WriteFile(dataPath, cipherText, 0660)
	buildDataID()
	buildTitles()
}

func changeMasterPass() {
	win := a.NewWindow("PmanagerMasterPass")

	logoPath, _ := os.UserHomeDir()
	logoPath += standardPath + "/logo.png"
	logoRes, err := fyne.LoadResourceFromPath(logoPath)
	if err != nil {
		logoRes = nil
	}
	logo := customWidget.NewIcon(logoRes, website)

	builtByPath, _ := os.UserHomeDir()
	builtByPath += standardPath + "/builtBy.png"
	builtByRes, err := fyne.LoadResourceFromPath(builtByPath)
	if err != nil {
		builtByRes = nil
	}
	builtBy := customWidget.NewIcon(builtByRes, website)

	lblMasterPass := widget.NewLabel("New Master Password:")
	entryMasterPass := widget.NewPasswordEntry()
	entryMasterPass.OnSubmitted = func(text string) {
		key = sha256.Sum256([]byte(text))
		win.Close()
	}

	gridLogin := layout.NewGridLayout(2)
	containerLogin := fyne.NewContainerWithLayout(gridLogin, lblMasterPass, entryMasterPass, logo, builtBy)

	win.Resize(fyne.NewSize(600, 400))
	win.CenterOnScreen()
	win.SetContent(containerLogin)
	go win.Show()
}

func buildTitles() {
	/* remove all */
	for i := len(containerTitles.Objects); i > 0; i-- {
		containerTitles.Remove(containerTitles.Objects[i-1])
	}

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
}

func buildContent(chosenTitle <-chan string, con *fyne.Container) {
	for /* true */{
		title, ok := <-chosenTitle
		if ok { /* remove old widgets */
			for i := len(con.Objects); i > 0; i-- {
				con.Remove(con.Objects[i-1])
			}
			/* build new widgets */
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
			/* add functionality */
			btnAdd := widget.NewButton("Add", func() {
				data[title] = append(data[title], "new", "new")
				customWidget.SendTitle <- title
			})
			btnGen := widget.NewButton("Generate", func() {
				generateRandom(title)
				customWidget.SendTitle <- title
			})
			con.Add(btnAdd)
			con.Add(btnGen)
			con.Refresh()
		}
	}
}

func generateRandom(title string) {
	titleConfig := "Random Generator"
	word := ""
	/* make default ascii list */
	if data[titleConfig] == nil {
		data[titleConfig] = append(data[titleConfig], "Length", "21", "Chars", "")
		for i := 0x21; i < 0x7f; i++ {
			data[titleConfig][3] += string(i)
		}
		ent := customWidget.NewTitleEntry()
		ent.ID = titleConfig
		ent.Text = titleConfig
		ent.OnSubmitted = func(string) {
			ent.Submitted(data, dataID)
		}
		containerTitles.Add(ent)
		containerTitles.Refresh()
	}
	length, err := strconv.Atoi(data[titleConfig][1])
	/* if no number set default value */
	if err != nil {
		length = 21
		data[titleConfig][1] = strconv.Itoa(length)
	}

	/* make nice random string */
	seedSlice := make([]byte, 8)
	rand.Read(seedSlice)
	var seed int64
	for i, v := range seedSlice {
		seed += int64(v << i*8)
	}
	mrand.Seed(int64(seed))

	skip := make([]byte, 1)
	for i := 0; i < length; i++ {
		rand.Read(skip)
		for k := byte(0); k < skip[0]; k++ {
			mrand.Int()
		}
		word += string(data[titleConfig][3][mrand.Int()%len(data[titleConfig][3])])
	}
	data[title] = append(data[title], "New random", word)
}

/* builds the slice from ground up new */
func buildDataID() {
	dataID = make([]string, 0, 10)
	for k := range data {
		dataID = append(dataID, k)
	}
}
