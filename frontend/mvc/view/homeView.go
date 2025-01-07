package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/acasais/sdac/configs"
)

type HomeView struct {
	WelcomeButtonTapped binding.Bool
}

func NewHomeView() *HomeView {
	return &HomeView{
		WelcomeButtonTapped: binding.NewBool(),
	}
}
func (hv *HomeView) Build() fyne.CanvasObject {

	//Texto Página Bienvenida
	textoBienvenida := widget.NewRichText(
		&widget.TextSegment{
			Text: "Bienvenido a Precio Electricidad\n\n",
			Style: widget.RichTextStyle{
				TextStyle: fyne.TextStyle{Bold: true},
				Alignment: fyne.TextAlignCenter,
			},
		},
		&widget.TextSegment{
			Text: "Esta aplicación muestra \n" +
				"el precio intradiario del mercado eléctrico \n" +
				"en España y Portugal.",
			Style: widget.RichTextStyle{
				Alignment: fyne.TextAlignCenter,
			}})

	logo := canvas.NewImageFromResource(configs.APP_LOGO)
	logo.FillMode = canvas.ImageFillOriginal

	botonBienvenida := hv.makeBotonBienvenida()
	return container.NewBorder(
		nil, container.NewCenter(container.NewGridWithRows(2, botonBienvenida, layout.NewSpacer())), nil, nil,
		container.NewCenter(container.NewGridWithRows(2, logo, textoBienvenida)))
}

func (hv *HomeView) makeBotonBienvenida() *widget.Button {
	botonBienvenida := widget.NewButton("                Continuar                ", func() {
		hv.WelcomeButtonTapped.Set(true)
	})
	botonBienvenida.Importance = widget.HighImportance
	return botonBienvenida
}
