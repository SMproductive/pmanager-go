package customWidget

import (
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var SendTitle chan string

type TitleEntry struct {
	widget.Entry
	ID string
}
/* Creates an Entry with some custom functionality. */
func NewTitleEntry() *TitleEntry {
	ent := &TitleEntry{}
	ent.ExtendBaseWidget(ent)
	ent.Disable()
	ent.Text = "new"
	ent.ID = ent.Text
	return ent
}
func (ent *TitleEntry) Submitted(data map[string] []string, dataID []string) {
	data[ent.Text] = append(data[ent.ID])
	delete(data, ent.ID)
	for i, v := range dataID {
		if ent.ID == v {
			dataID[i] = ent.Text
			break
		}
	}
	ent.ID = ent.Text
	ent.Disable()
	i := 0
	for k := range data {
		dataID[i] = k
	}
}
/* Sends the "id" through the channel "SendTitle". */
func (ent *TitleEntry) Tapped(_ *fyne.PointEvent) {
	if ent.Disabled() {
		SendTitle <- ent.ID
	}
}
/* Enables or disables the widget. */
func (ent *TitleEntry) TappedSecondary(_ *fyne.PointEvent) {
	if ent.Disabled() {
		ent.Enable()
	} else {
		ent.Disable()
	}
}

type ContentEntry struct {
	widget.Entry
	ID *string
}
/* Creates an Entry with some custom functionality. */
func NewContentEntry() *ContentEntry {
	ent := &ContentEntry{}
	ent.ExtendBaseWidget(ent)
	ent.Disable()
	return ent
}
func (ent *ContentEntry) Submitted() {
	*ent.ID = ent.Text
	ent.Disable()
}
/* If disabled: Copies the "Text" to clipboard else: it behaves as normal "widget.Entry". */
func (ent *ContentEntry) Tapped(_ *fyne.PointEvent) {
	if ent.Disabled() {
		clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
		fyne.Clipboard.SetContent(clipboard, ent.Text)
	}
}
/* Enables or disables the widget. */
func (ent *ContentEntry) TappedSecondary(_ *fyne.PointEvent) {
	if ent.Disabled() {
		ent.Enable()
	} else {
		ent.Disable()
	}
}
func (ent *ContentEntry) DoubleTapped(_ *fyne.PointEvent) {
	ent.Password = !ent.Password
	ent.Refresh()
}
