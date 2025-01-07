package controller

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/acasais/sdac/frontend/mvc/view"
	"github.com/acasais/sdac/service"
)

type pricesViewController struct {
	a           fyne.App
	w           fyne.Window
	v           *view.PricesView
	ws          *service.WebService
	pricesCache map[string][]string
}

func newPricesVC(a fyne.App, w fyne.Window) *pricesViewController {

	return &pricesViewController{
		a:           a,
		w:           w,
		v:           view.NewPricesView(),
		ws:          service.NewWS(),
		pricesCache: make(map[string][]string),
	}
}

func (pvc *pricesViewController) updatePrices() {
	t := time.Now()	
	if t.Hour() > 13 {
		t = t.Add(24 * time.Hour)
	}
	date := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
	pvc.v.PriceDate.Set(t.Local().Format("02/01/2006"))
	_, exists := pvc.pricesCache[date]
	if !exists {
		p, err := pvc.ws.GetPricesForDate(date)
		if err != nil || len(p) <= 0 {
			return
		}
		pvc.pricesCache[date] = p[1 : len(p)-1]
	}
	pvc.v.PriceList.Set(pvc.pricesCache[date])
	pvc.w.SetContent(pvc.v.Build())
}

func (pvc *pricesViewController) callbackRefresh() binding.DataListener {
	return binding.NewDataListener(func() {

		tapped, _ := pvc.v.RefreshButtonTapped.Get()
		if tapped {
			defer pvc.v.RefreshButtonTapped.Set(false)
			pvc.updatePrices()
		}
	})
}

func (pvc *pricesViewController) removeListenerRefreshButton() {
	pvc.v.RefreshButtonTapped.RemoveListener(pvc.callbackRefresh())
}
func (pvc *pricesViewController) addListenerRefreshButton() {
	pvc.v.RefreshButtonTapped.AddListener(pvc.callbackRefresh())
}
