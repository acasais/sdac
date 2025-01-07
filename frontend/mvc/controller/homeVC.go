package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/acasais/sdac/configs"
	"github.com/acasais/sdac/frontend/mvc/view"
)

type homeViewController struct {
	a fyne.App
	w fyne.Window
	v *view.HomeView
}

func newHomeVC(a fyne.App, w fyne.Window) *homeViewController {

	return &homeViewController{
		a: a,
		w: w,
		v: view.NewHomeView(),
	}
}

func (hvc *homeViewController) addListenerWelcomeButton() {
	callback := binding.NewDataListener(func() {
		tapped, _ := hvc.v.WelcomeButtonTapped.Get()
		if tapped {
			defer hvc.v.WelcomeButtonTapped.Set(false)
			pref := hvc.a.Preferences()
			pref.SetBool(configs.PREFERENCES_WELCOME_DISPLAYED, true)
			hvc.showPricesView()
		}
	})
	hvc.v.WelcomeButtonTapped.AddListener(callback)
}

func (hvc *homeViewController) showPricesView() {
	pvc := newPricesVC(hvc.a, hvc.w)
	go pvc.updatePrices()
	pvc.addListenerRefreshButton()
	hvc.w.SetContent(pvc.v.Build())
	if fyne.CurrentApp().Driver().Device().IsMobile() {
		pvc.w.SetFullScreen(true)
	} else {
		pvc.w.Resize(fyne.NewSize(configs.WIDTH_WINDOW_SIZE, configs.HEIGHT_WINDOW_SIZE))
	}
	pvc.w.SetOnClosed(func() {
		pvc.removeListenerRefreshButton()
	})
	hvc.w.Show()
}
