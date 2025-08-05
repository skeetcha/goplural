package main

import (
	"database/sql"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

type AppState struct {
	selectedMember        int
	db                    *sql.DB
	defaultAvatar         *fyne.StaticResource
	avatarSettingsSize    fyne.Size
	avatarChangedText     bool
	currentSettingsMember int
}

func main() {
	var appState AppState

	app := app.New()
	mainWindow := app.NewWindow("GoPlural")
	loadDatabases(app, &appState)
	defer appState.db.Close()

	defaultAvatar, err := loadDefaultAvatar()

	if err != nil {
		return
	}

	appState.defaultAvatar = defaultAvatar
	appState.avatarSettingsSize = fyne.NewSize(300., 300.)
	appState.avatarChangedText = true
	appState.currentSettingsMember = -1

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.SetMainMenu(makeMenu(app, mainWindow, &appState))
	mainWindow.SetContent(makeChatScreen(app, mainWindow, &appState))
	mainWindow.Show()
	app.Run()
}

func makeMenu(app fyne.App, window fyne.Window, state *AppState) *fyne.MainMenu {
	settingsItem := fyne.NewMenuItem("Settings", func() { openSettings(app, state) })

	settingsShortcut := &desktop.CustomShortcut{
		KeyName:  fyne.KeyComma,
		Modifier: fyne.KeyModifierShortcutDefault,
	}

	settingsItem.Shortcut = settingsShortcut
	window.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings(app, state)
	})

	file := fyne.NewMenu("File")
	device := fyne.CurrentDevice()

	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}

	main := fyne.NewMainMenu(
		file,
	)

	return main
}

func loadDefaultAvatar() (*fyne.StaticResource, error) {
	data, err := os.ReadFile("avatars/default_avatar.png")

	if err != nil {
		log.Println("Error reading default avatar image:", err)
		return nil, err
	}

	return fyne.NewStaticResource("defaultAvatar", data), nil
}
