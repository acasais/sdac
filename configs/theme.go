package configs

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type SdacTheme struct {
	Theme string
}

func (t *SdacTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch t.Theme {
	case DARK_THEME:
		variant = theme.VariantDark
		switch name {
		case theme.ColorNameDisabled:
			return color.NRGBA{R: 0x55, G: 0x55, B: 0x55, A: 0xff}
		case theme.ColorNameBackground:
			return color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff}
		case theme.ColorNameButton:
			return color.NRGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{R: 0x26, G: 0x26, B: 0x26, A: 0xff}
		case theme.ColorNameOverlayBackground:
			return color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff}
		case theme.ColorNameMenuBackground:
			return color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff}
		}

	case LIGHT_THEME:
		variant = theme.VariantLight
		switch name {
		case theme.ColorNameDisabled:
			return color.NRGBA{R: 0xab, G: 0xab, B: 0xab, A: 0xff}
		case theme.ColorNameInputBorder:
			return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{R: 0xe5, G: 0xe5, B: 0xe5, A: 0xff}
		}
	}
	theme.InnerPadding()
	return theme.DefaultTheme().Color(name, variant)
}

func (t *SdacTheme) Font(s fyne.TextStyle) fyne.Resource {

	if s.Monospace {
		if s.Italic {
			return resourceGoMonoItalicTtf
		}
		if s.Bold {
			return resourceGoMonoBoldTtf
		}
		return resourceGoMonoTtf
	}
	if s.Bold {
		if s.Italic {
			return resourceGoBoldItalicTtf
		}
		return resourceGoBoldTtf
	}
	if s.Italic {
		return resourceGoItalicTtf
	}

	return resourceGoRegularTtf

}

func (f *SdacTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (r *SdacTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
