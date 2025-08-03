package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

type AppState struct {
	openScreen     Screen
	selectedMember int
}

type Screen int

const (
	Screen_None Screen = iota
	Screen_Chat
	Screen_Dialog
)

var appSettings AppSettings
var appState AppState
var defaultAvatar *fyne.StaticResource
var avatarSettingsSize []float32

func main() {
	avatarSettingsSize = []float32{300.0, 300.0}
	app := app.New()
	mainWindow := app.NewWindow("GoPlural")
	loadSettings(app)
	loadDefaultAvatar()
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.SetMainMenu(makeMenu(app, mainWindow))
	mainWindow.SetContent(makeChatScreen(app, mainWindow))
	mainWindow.Show()
	app.Run()
	saveSettings(app)
}

func makeChatScreen(app fyne.App, window fyne.Window) fyne.CanvasObject {
	memberList := widget.NewList(
		func() int {
			return len(appSettings.Members)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			var avatarURI fyne.URI

			if appSettings.Members[id].Avatar != "" {
				if appSettings.Members[id].Avatar[0] == '/' {
					avatarURI = storage.NewFileURI(appSettings.Members[id].Avatar)
				} else {
					avatarURI = storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), appSettings.Members[id].Avatar))
				}
			}

			o.(*fyne.Container).Objects[0].(*canvas.Image).Image = (*canvas.NewImageFromURI(avatarURI)).Image
			o.(*fyne.Container).Objects[1].(*widget.Label).Text = appSettings.Members[id].Name
		},
	)

	memberList.OnSelected = func(id widget.ListItemID) {
		appState.selectedMember = id
	}

	chatContainer := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("System Members"),
		widget.NewLabel("Chat History"),
		memberList,
		makeChats(app, window),
	)

	return chatContainer
}

func makeChats(app fyne.App, window fyne.Window) fyne.CanvasObject {

}

func makeMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	settingsItem := fyne.NewMenuItem("Settings", func() { openSettings(app) })

	settingsShortcut := &desktop.CustomShortcut{
		KeyName:  fyne.KeyComma,
		Modifier: fyne.KeyModifierShortcutDefault,
	}

	settingsItem.Shortcut = settingsShortcut
	window.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings(app)
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

func openSettings(app fyne.App) {
	settingsWindow := app.NewWindow("Settings")
	setting := settings.NewSettings()

	tabs := container.NewAppTabs(
		container.NewTabItem("Appearance", setting.LoadAppearanceScreen(settingsWindow)),
		container.NewTabItem("Members", buildMemberSettings(app, settingsWindow)),
	)

	settingsWindow.SetContent(tabs)
	settingsWindow.Resize(fyne.NewSize(440, 520))
	settingsWindow.Show()
}

func buildMemberSettings(app fyne.App, window fyne.Window) fyne.CanvasObject {
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

	var tabContainer *fyne.Container

	nameEntry := widget.NewEntry()
	avatarEntry := widget.NewEntry()
	pronounEntry := widget.NewEntry()

	avatarImage := canvas.NewImageFromResource(defaultAvatar)
	avatarImage.ScaleMode = canvas.ImageScaleFastest

	updateImageSize := func() {
		avatarImage.SetMinSize(createSize(avatarSettingsSize))
	}

	updateImageSize()

	avatarValidator := func(text string) error {
		var imagePath string

		if text == "" {
			imagePath = text
		} else {
			if text[0] == '/' {
				imagePath = text
			} else {
				imagePath = path.Join(app.Storage().RootURI().Path(), text)
			}
		}

		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			return errors.New("file does not exist")
		}

		if imageURI := storage.NewFileURI(imagePath); !strings.Contains(imageURI.MimeType(), "image") {
			return errors.New("text is not an image (found " + imageURI.MimeType() + ")")
		}

		return nil
	}

	avatarEntry.OnChanged = func(text string) {
		if err := avatarEntry.Validate(); err != nil {
			log.Println("Error updating image:", err)
			return
		}

		var avatarURI fyne.URI

		if text[0] == '/' {
			avatarURI = storage.NewFileURI(text)
		} else {
			avatarURI = storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), text))
		}

		*avatarImage = *canvas.NewImageFromURI(avatarURI)
		avatarImage.Refresh()
		tabContainer.Refresh()
		avatarSettingsSize[1] = (300.0 * float32(avatarImage.Image.Bounds().Max.Y)) / float32(avatarImage.Image.Bounds().Max.X)
		updateImageSize()
	}

	memberForm := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("Name"),
		nameEntry,
		widget.NewLabel("Avatar"),
		avatarEntry,
		widget.NewLabel("Pronouns"),
		pronounEntry,
	)

	memberForm.Hidden = true

	list.OnSelected = func(i widget.ListItemID) {
		nameEntry.Unbind()
		avatarEntry.Unbind()
		pronounEntry.Unbind()
		memberForm.Hidden = false
		nameEntry.Bind(binding.BindString(&appSettings.Members[i].Name))
		avatarEntry.Bind(binding.BindString(&appSettings.Members[i].Avatar))
		avatarEntry.Validator = avatarValidator
		pronounEntry.Bind(binding.BindString(&appSettings.Members[i].Pronouns))
		var avatarURI fyne.URI

		if appSettings.Members[i].Avatar[0] == '/' {
			avatarURI = storage.NewFileURI(appSettings.Members[i].Avatar)
		} else {
			avatarURI = storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), appSettings.Members[i].Avatar))
		}

		*avatarImage = *canvas.NewImageFromURI(avatarURI)
	}

	outerRight := container.NewVBox(
		memberForm,
		avatarImage,
		widget.NewButton("Create Member", func() {
			appSettings.Members = append(appSettings.Members, Member{
				Name: "New Member",
			})

			newItem := list.CreateItem()
			list.UpdateItem(list.Length()-1, newItem)
		}),
		widget.NewButton("PluralKit Import", func() {
			PKImport(app, window, &appSettings.Members)
		}),
		widget.NewButton("SimplyPlural Import", func() {
			SPImport(app, window, &appSettings.Members)
		}),
	)

	tabContainer = container.New(
		layout.NewFormLayout(),
		list,
		outerRight,
	)

	return tabContainer
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

func loadDefaultAvatar() {
	data, err := os.ReadFile("avatars/default_avatar.png")

	if err != nil {
		log.Println("Error reading default avatar image:", err)
		return
	}

	defaultAvatar = fyne.NewStaticResource("defaultAvatar", data)
}

func createSize(size []float32) fyne.Size {
	return fyne.NewSize(size[0], size[1])
}
