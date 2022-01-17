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
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
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
