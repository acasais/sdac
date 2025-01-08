package view

import (
	"bytes"
	"image/color"
	"image/png"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/acasais/sdac/configs"
	"github.com/go-analyze/charts"
)

type PricesView struct {
	PriceDate           binding.String
	PriceList           binding.StringList
	RefreshButtonTapped binding.Bool
	plot                *canvas.Image
}

func NewPricesView() *PricesView {
	return &PricesView{
		PriceDate:           binding.NewString(),
		PriceList:           binding.NewStringList(),
		RefreshButtonTapped: binding.NewBool(),
	}
}
func (pv *PricesView) Build() fyne.CanvasObject {

	refreshButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		pv.RefreshButtonTapped.Set(true)
	})

	refreshButton.Importance = widget.LowImportance

	mkPriceDate := widget.NewRichTextFromMarkdown("# Precios ")
	pv.PriceDate.AddListener(binding.NewDataListener(func() {
		pd, _ := pv.PriceDate.Get()
		mkPriceDate.ParseMarkdown("# Precios **" + pd + "**")
	}))

	pv.plot = pv.makePlot()
	return container.NewBorder(
		container.NewBorder(nil, nil, nil, refreshButton, container.NewCenter(mkPriceDate)),
		container.NewCenter(canvas.NewText("fuente: omie.es", &color.RGBA{100, 78, 21, 255})),
		nil,
		nil,
		pv.plot,
	)
}

func (pv *PricesView) makePlot() *canvas.Image {
	var hours []string
	var prices1, prices2 []float64
	prices, _ := pv.PriceList.Get()
	if len(prices) == 0 {
		logo := canvas.NewImageFromResource(configs.APP_LOGO)
		logo.FillMode = canvas.ImageFillContain
		return logo
	}
	for i, p := range prices {
		f, err := strconv.ParseFloat(p, 64)
		if err != nil {
			continue
		}
		if i%3 == 0 {
			hours = append(hours, p)
		} else {
			//prices = append(prices, f)
			if i%2 == 0 {
				prices1 = append(prices1, f)
			} else {
				prices2 = append(prices2, f)
			}
		}
	}
	values := [][]float64{prices1, prices2}

	opt := charts.BarChartOption{}
	opt.SeriesList = charts.NewSeriesListDataFromValues(values, charts.ChartTypeBar)
	opt.XAxis.Data = hours

	p, err := charts.NewPainter(charts.PainterOptions{
		OutputFormat: charts.ChartOutputPNG,
		Width:        configs.WIDTH_WINDOW_SIZE - 10,
		Height:       configs.HEIGHT_WINDOW_SIZE - 10,
	})

	if err != nil {
		return nil
	}
	_, err = charts.NewBarChart(p, opt).Render()
	if err != nil {
		return nil
	}

	buf, err := p.Bytes()
	if err != nil {
		return nil
	}

	img, err := png.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil
	}
	plotImage := canvas.NewImageFromImage(img)
	plotImage.Image = img
	plotImage.FillMode = canvas.ImageFillOriginal

	return plotImage
}
