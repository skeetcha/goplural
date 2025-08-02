package main

import (
	"github.com/skeetcha/goplural/themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type GoPlural struct {
	themes []themes.Theme
}

func main() {
	appData := GoPlural{
		themes: themes.SetupThemes(),
	}

	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(800, 600))
	w.SetMainMenu(makeMenu(a, w))
	w.Show()
	a.Run()
}

func makeMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)

	file := fyne.NewMenu("File", newItem)

	main := fyne.NewMainMenu(
		file,
		// something else, presumably
	)

	return main
}
