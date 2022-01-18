package customTheme

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Nord struct {
	ColorNameBackground string

}

func (t Nord) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x2e, G: 0x34, B: 0x40, A: 0xff}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xec, G: 0xef, B: 0xf4, A: 0xff}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x2e, G: 0x34, B: 0x40, A: 0xff}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x88, G: 0xc0, B: 0xd0, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0xec, G: 0xef, B: 0xf4, A: 0xff}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0x2e, G: 0x34, B: 0x40, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x88, G: 0xc0, B: 0xd0, A: 0xff}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 0x4c, G: 0x56, B: 0x6a, A: 0xff}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0xd8, G: 0xde, B: 0xe9, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0x3b, G: 0x42, B: 0x52, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0x3a, G: 0xbe, B: 0x8c, A: 0xff}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}
func (t Nord) Icon(name fyne.ThemeIconName) fyne.Resource {
	if name == theme.IconNameHome {
		//fyne.NewStaticResource("myHome", homeBytes)
	}
	return theme.DefaultTheme().Icon(name)
}
func (t Nord) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t Nord) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
