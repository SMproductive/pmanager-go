package customWidget

import (
	"net/url"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Icon struct {
	widget.Icon
	URL *url.URL
}
func NewIcon(res fyne.Resource, link string) *Icon {
	i := &Icon{}
	i.ExtendBaseWidget(i)
	i.SetResource(res)
	u, _ := url.Parse(link)
	i.URL = u
	return i
}
func (i *Icon) Tapped(_ *fyne.PointEvent) {
	fyne.CurrentApp().OpenURL(i.URL)
}
