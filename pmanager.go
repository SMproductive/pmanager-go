package main
/* TODO
* cloud implementation (google, nextcloud)
*/

import (
	"fmt"
	"crypto/sha256"
	"crypto/rand"
	"encoding/json"
	"runtime"
	"io/ioutil"
	mrand "math/rand"
	"os"
	"strconv"
	"sort"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SMproductive/pmanager-go/customWidget"
	"github.com/SMproductive/pmanager-go/customTheme"
)

const titleConfig string = "Random Generator"

var a fyne.App

var key [32]byte
var dataPath string
var data map[string] []string
var dataID []string

const website string = "https://github.com/SMproductive/pmanager-go"

func main() {
	runtime.GOMAXPROCS(10)
	a = app.New()
	a.Settings().SetTheme(customTheme.Nord{})

	win := a.NewWindow("pmanager")
	win.CenterOnScreen()
	login(win)
	win.ShowAndRun()
}

func login(win fyne.Window) {
	home, _ := os.UserHomeDir()
	dataPath = home + "/.pmanager/passwords"

	lblDatabase := widget.NewLabel("Database: ")

	entryDatabase := customWidget.NewLoginEntry()
	entryDatabase.Password = false
	entryDatabase.Text = dataPath

	logo := customWidget.NewIcon(resourceLogoPng, website)
	builtBy := customWidget.NewIcon(resourceBuiltByPng, website)

	lblMasterPass := widget.NewLabel("Master Password:")
	entryMasterPass := customWidget.NewLoginEntry()
	entryMasterPass.OnSubmitted = func(password string) {
		dataPath = entryDatabase.Text
		if decrypt(password) == nil {
			UI(win)
		}
	}

	gridLogin := layout.NewGridLayout(2)
	containerLogin := fyne.NewContainerWithLayout(gridLogin, lblDatabase, entryDatabase, lblMasterPass, entryMasterPass, logo, builtBy)

	win.Resize(fyne.NewSize(600, 400))
	win.SetContent(containerLogin)
	data = make(map[string][]string)
}

func UI(win fyne.Window) {
	var mainSplit *container.Split
	buildDataID()

	gridContent := layout.NewGridLayout(2)
	containerContent := fyne.NewContainerWithLayout(gridContent)

	var containerTitles *fyne.Container
	var scrollTitles *container.Scroll

	btnChangePassword := widget.NewButton("Change MPW", func() {
		changeMasterPass(containerTitles, containerContent, win)
	})
	btnAddTitle := widget.NewButton("Add", func() {
		addTitle(containerTitles, containerContent)
	})
	btnSave := widget.NewButton("Save and close", func() {
		save(containerTitles, containerContent)
	})
	topBox := container.NewHBox(btnAddTitle, btnSave, btnChangePassword)

	containerTitles = container.NewHBox()
	scrollTitles = container.NewHScroll(containerTitles)

	buildTitles(containerTitles, containerContent)

	topSplit := container.NewHSplit(topBox, scrollTitles)
	topSplit.SetOffset(0)

	mainSplit = container.NewVSplit(topSplit, containerContent)
	mainSplit.SetOffset(0.12)
	win.SetContent(mainSplit)

}

func addTitle(titles, content *fyne.Container) {
	str := "new"
	if data[str] == nil {
		data[str] = nil
		dataID = append(dataID, str)
		ent := customWidget.NewTitleEntry()
		ent.SetContent = buildContent
		ent.OnSubmitted = func(string) {
			ent.Submitted(data, dataID)
		}
		ent.ContentContainer = content
		ent.TitlesContainer = titles
		ent.Text = str
		ent.ID = str
		titles.Add(ent)
		titles.Refresh()
	}
}

func save(containerTitles, contentContainer *fyne.Container) {
	/* clean data */
	delete(data, "")
	for k := range data {
		length := 0
		for _, i := range data[k]{
			if i != "" {
				data[k][length] = i
				length++
			}
		}
		data[k] = data[k][:length]
	}

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
	os.Mkdir(dir, 0777)
	ioutil.WriteFile(dataPath, cipherText, 0660)
	buildDataID()
	buildTitles(containerTitles, contentContainer)
}

func changeMasterPass(titles, content *fyne.Container, win fyne.Window) {
	logo := customWidget.NewIcon(resourceLogoPng, website)
	builtBy := customWidget.NewIcon(resourceBuiltByPng, website)

	lblMasterPass := widget.NewLabel("New Master Password:")
	entryMasterPass := customWidget.NewLoginEntry()
	entryMasterPass.OnSubmitted = func(password string) {
		key = sha256.Sum256([]byte(password))
		UI(win)
	}

	gridLogin := layout.NewGridLayout(2)
	containerLogin := fyne.NewContainerWithLayout(gridLogin, lblMasterPass, entryMasterPass, logo, builtBy)

	win.SetContent(containerLogin)
	win.Show()
}

func buildTitles(titles, content *fyne.Container) {
	/* remove all */
	for i := len(titles.Objects); i > 0; i-- {
		titles.Remove(titles.Objects[i-1])
	}

	/* Build all titles */
	for i := len(data); i > 0; i-- {
		ent := customWidget.NewTitleEntry()
		ent.OnSubmitted = func(title string) {
			ent.Submitted(data, dataID)
		}
		ent.SetContent = buildContent
		ent.ContentContainer = content
		ent.TitlesContainer = titles
		ent.Text = dataID[i-1]
		ent.ID = dataID[i-1]

		titles.Add(ent)
		ent.Refresh()
	}
	titles.Refresh()
}

func buildContent(chosenTitle string, titles, content *fyne.Container) {
	/* Mark selected item */
	for i := len(titles.Objects); i>0; i-- {
		titles.Objects[i-1].(*customWidget.TitleEntry).TextStyle.Bold = false
		if titles.Objects[i-1].(*customWidget.TitleEntry).Text == chosenTitle {
			titles.Objects[i-1].(*customWidget.TitleEntry).TextStyle.Bold = true
		}
		titles.Objects[i-1].Refresh()
	}
	/* Remove previous items */
	for i := len(content.Objects); i > 0; i-- {
		content.Remove(content.Objects[i-1])
	}
	/* build new widgets */
	for i, v := range data[chosenTitle] {
		ent := customWidget.NewContentEntry()
		ent.ID = &data[chosenTitle][i]
		ent.Text = v
		ent.Password = i % 2 == 1
		ent.OnSubmitted = func(string) {
			ent.Submitted()
		}
		content.Add(ent)
	}
	/* add functionality */
	btnAdd := widget.NewButton("Add", func() {
		data[chosenTitle] = append(data[chosenTitle], "", "")
		buildContent(chosenTitle,titles, content)
	})
	btnGen := widget.NewButton("Generate", func() {
		generateRandom(chosenTitle, titles, content)
		buildContent(chosenTitle, titles, content)
	})
	if chosenTitle == titleConfig {
		return
	}
	content.Add(btnAdd)
	content.Add(btnGen)
	content.Refresh()
}

func generateRandom(title string, titles, content *fyne.Container) {
	word := ""
	/* make default ascii list */
	if len(data[titleConfig]) < 4 {
		data[titleConfig] = append(data[titleConfig], "Length", "21", "Chars", "")
		for i := 0x21; i < 0x7f; i++ {
			data[titleConfig][3] += string(i)
		}
		ent := customWidget.NewTitleEntry()
		ent.ID = titleConfig
		ent.Text = titleConfig
		ent.ContentContainer = content
		ent.TitlesContainer = titles
		ent.OnSubmitted = func(string) {
			ent.Submitted(data, dataID)
		}
		content.Add(ent)
		content.Refresh()
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
	data[title] = append(data[title], "Password:", word)
}

/* builds the slice from ground up new */
func buildDataID() {
	dataID = make([]string, 0, 10)
	for k := range data {
		dataID = append(dataID, k)
	}
	sort.Strings(dataID)
}

func decrypt(password string) error {
	/* Decryption of database */
	key = sha256.Sum256([]byte(password))
	cipherText, err := ioutil.ReadFile(dataPath)
	if err == nil {
		jsonData, err := dec(key[:], cipherText)
		if err != nil {
			return err
		}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			return err
		}
	} else {
		noFileErr := "open " + dataPath + ": no such file or directory"
		if fmt.Sprint(err) == noFileErr {
			return nil
		}
		if err != nil {
			win := a.NewWindow("error" )
			str := fmt.Sprint(err)
			win.SetContent(widget.NewLabel(string(str)))
			win.Show()
			return err
		}
	}
	return nil
}
