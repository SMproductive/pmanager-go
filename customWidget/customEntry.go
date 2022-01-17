package customWidget

import (
	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var SendTitle chan string

type TitleEntry struct {
	widget.Entry
	IsBound bool
}
/* Creates an Entry with some custom functionality. */
func NewTitleEntry() *TitleEntry {
	ent := &TitleEntry{}
	ent.ExtendBaseWidget(ent)
	ent.Disable()
	return ent
}
/* Binds a "binding.String" and sets "IsBound" to true. */
func (ent *TitleEntry) BindStr(str binding.String) {
	ent.Bind(str)
	ent.IsBound = true
}
/* Unbinds a "binding.String" and sets "IsBound" to false. */
func (ent *TitleEntry) UnbindStr() {
	ent.Unbind()
	ent.IsBound = false
}
/* Sends the "Text" through the channel "SendTitle". */
func (ent *TitleEntry) Tapped(_ *fyne.PointEvent) {
	if ent.Disabled() {
		SendTitle <- ent.Text
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
	TitleEntry
}
/* Creates an Entry with some custom functionality. */
func NewContentEntry() *ContentEntry {
	ent := &ContentEntry{}
	ent.ExtendBaseWidget(ent)
	ent.Disable()
	return ent
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
