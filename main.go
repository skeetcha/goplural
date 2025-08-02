package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type AppSettings struct {
	Members []Member `json:"members"`
}

var appSettings AppSettings

func main() {
	app := app.New()
	mainWindow := app.NewWindow("GoPlural")
	loadSettings(app)
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.SetMainMenu(makeMenu(app, mainWindow))
	mainWindow.Show()
	app.Run()
	saveSettings(app)
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
		container.NewTabItem("Members", buildMemberSettings()),
	)

	settingsWindow.SetContent(tabs)
	settingsWindow.Resize(fyne.NewSize(440, 520))
	settingsWindow.Show()
}

func buildMemberSettings() fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return len(appSettings.Members)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(binding.BindString(&appSettings.Members[i].Name))
			o.(*widget.Label).Wrapping = fyne.TextWrapBreak
		},
	)

	nameEntry := widget.NewEntry()
	avatarEntry := widget.NewEntry()
	pronounEntry := widget.NewEntry()
	proxyEntry := widget.NewEntry()

	memberForm := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("Name"),
		nameEntry,
		widget.NewLabel("Avatar"),
		avatarEntry,
		widget.NewLabel("Pronouns"),
		pronounEntry,
		widget.NewLabel("Proxy"),
		proxyEntry,
	)

	memberForm.Hidden = true

	list.OnSelected = func(i widget.ListItemID) {
		nameEntry.Unbind()
		avatarEntry.Unbind()
		pronounEntry.Unbind()
		proxyEntry.Unbind()
		memberForm.Hidden = false
		nameEntry.Bind(binding.BindString(&appSettings.Members[i].Name))
		avatarEntry.Bind(binding.BindString(&appSettings.Members[i].Avatar))
		pronounEntry.Bind(binding.BindString(&appSettings.Members[i].Pronouns))
		proxyEntry.Bind(binding.BindString(&appSettings.Members[i].Proxy))
	}

	outerRight := container.NewVBox(
		memberForm,
		widget.NewButton("Create Member", func() {
			appSettings.Members = append(appSettings.Members, Member{
				Name: "New Member",
			})

			newItem := list.CreateItem()
			list.UpdateItem(list.Length()-1, newItem)
		}),
		widget.NewButton("PluralKit Import", func() {
			fmt.Println("PluralKit Import not implemented yet")
		}),
		widget.NewButton("SimplyPlural Import", func() {
			fmt.Println("SimplyPlural Import not implemented yet")
		}),
	)

	return container.New(
		layout.NewFormLayout(),
		list,
		outerRight,
	)
}

func loadSettings(app fyne.App) {
	settingsURI := storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), "settings.json"))
	f, err := storage.Reader(settingsURI)

	if err != nil {
		log.Println("Error opening config for reading:", err)
		return
	}

	defer f.Close()
	err = json.NewDecoder(f).Decode(&appSettings)

	if err != nil {
		log.Println("Error loading config:", err)
	}
}

func saveSettings(app fyne.App) {
	settingsURI := storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), "settings.json"))
	f, err := storage.Writer(settingsURI)

	if err != nil {
		log.Println("Error opening config for writing:", err)
		return
	}

	defer f.Close()

	err = json.NewEncoder(f).Encode(appSettings)

	if err != nil {
		log.Println("Error writing config:", err)
	}
}
