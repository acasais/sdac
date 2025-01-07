package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/acasais/sdac/configs"
)

type MainController struct {
}

func NewMainController() *MainController {
	return &MainController{}
}

func (mc *MainController) BuildAndRunUI() {
	a := app.NewWithID(configs.APP_ID)
	w := a.NewWindow(configs.APP_NAME)

	a.Settings().SetTheme(&configs.SdacTheme{Theme: configs.LIGHT_THEME})

	pref := a.Preferences()
	hvc := newHomeVC(a, w)

	if pref.Bool(configs.PREFERENCES_WELCOME_DISPLAYED) {
		hvc.showPricesView()
		hvc.a.Run()
		return
	}
	hvc.addListenerWelcomeButton()
	w.SetContent(hvc.v.Build())
	if fyne.CurrentApp().Driver().Device().IsMobile() {
		w.SetFullScreen(true)
	} else {
		w.Resize(fyne.NewSize(configs.WIDTH_WINDOW_SIZE, configs.HEIGHT_WINDOW_SIZE))
	}
	w.ShowAndRun()
}
