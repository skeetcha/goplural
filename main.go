package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

type AppSettings struct {
	Members []Member `json:"members"`
}

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Hello World")
	app.Preferences().SetInt("test.four", 4)
	app.Preferences().SetString("test.five", "five")
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.SetMainMenu(makeMenu(app, mainWindow))
	mainWindow.Show()
	app.Run()
	fmt.Println(app.Preferences())
}

func makeMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	settingsItem := fyne.NewMenuItem("Settings", func() { openSettings(app) })

	settingsShortcut := &desktop.CustomShortcut{
		KeyName:  fyne.KeyComma,
		Modifier: fyne.KeyModifierShortcutDefault,
	}

	settingsItem.Shortcut = settingsShortcut
	window.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings(app)
	})

	file := fyne.NewMenu("File", newItem)
	device := fyne.CurrentDevice()

	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}

	main := fyne.NewMainMenu(
		file,
	)

	return main
}

func openSettings(app fyne.App) {
	settingsWindow := app.NewWindow("Settings")
	setting := settings.NewSettings()

	tabs := container.NewAppTabs(
		container.NewTabItem("Appearance", setting.LoadAppearanceScreen(settingsWindow)),
		container.NewTabItem("Members", buildMemberSettings(app)),
	)

	settingsWindow.SetContent(tabs)
	settingsWindow.Resize(fyne.NewSize(440, 520))
	settingsWindow.Show()
}

func buildMemberSettings(app fyne.App) fyne.CanvasObject {
	return nil
}
