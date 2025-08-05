package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func openSettings(app fyne.App, state *AppState) {
	settingsWindow := app.NewWindow("Settings")
	setting := settings.NewSettings()

	tabs := container.NewAppTabs(
		container.NewTabItem("Appearance", setting.LoadAppearanceScreen(settingsWindow)),
		container.NewTabItem("Members", buildMemberSettings(app, settingsWindow, state)),
	)

	settingsWindow.SetContent(tabs)
	settingsWindow.Resize(fyne.NewSize(440, 520))
	settingsWindow.Show()
}

func buildMemberSettings(app fyne.App, window fyne.Window, state *AppState) fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			res, err := state.db.Query("select count(*) from members")

			if err != nil {
				log.Fatal("Error getting count of members:", err)
			}

			defer res.Close()
			res.Next()
			var result int
			err = res.Scan(&result)

			if err != nil {
				log.Fatal("Error getting count of numbers (while scanning):", err)
			}

			return result
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			res, err := state.db.Query("select name from members where id=" + strconv.Itoa(i+1))

			if err != nil {
				log.Println("Error getting name of member "+strconv.Itoa(i+1)+":", err)
				return
			}

			defer res.Close()
			res.Next()
			var name string
			err = res.Scan(&name)

			if err != nil {
				log.Println("Error getting name of member "+strconv.Itoa(i+1)+" (while scanning):", err)
				return
			}

			o.(*widget.Label).Text = name
			o.(*widget.Label).Wrapping = fyne.TextWrapBreak
			o.Refresh()
		},
	)

	var tabContainer *fyne.Container

	nameEntry := widget.NewEntry()
	avatarEntry := widget.NewEntry()
	pronounEntry := widget.NewEntry()

	avatarImage := canvas.NewImageFromResource(state.defaultAvatar)
	avatarImage.ScaleMode = canvas.ImageScaleFastest

	updateImageSize := func() {
		avatarImage.SetMinSize(state.avatarSettingsSize)
		canvas.Refresh(avatarImage)
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
		if state.avatarChangedText {
			if err := avatarEntry.Validate(); err != nil {
				log.Println("Error updating image:", err)
				return
			}

			_, err := state.db.Exec("update members set avatar_path = '" + text + "' where id = " + strconv.Itoa(state.currentSettingsMember+1))

			if err != nil {
				log.Println("Error updating avatar url:", err)
				return
			}
		}

		var avatarURI fyne.URI

		if text[0] == '/' {
			avatarURI = storage.NewFileURI(text)
		} else {
			avatarURI = storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), text))
		}

		img, err := LoadImage(avatarURI.Path())

		if err != nil {
			log.Println("Error loading image:", err)
			return
		}

		(*avatarImage).Image = img
		tabContainer.Refresh()
		state.avatarSettingsSize.Height = (300.0 * float32(img.Bounds().Max.Y)) / float32(img.Bounds().Max.X)
		updateImageSize()
		state.avatarChangedText = true
	}

	nameEntry.OnChanged = func(text string) {
		if state.avatarChangedText {
			_, err := state.db.Exec("update members set name = '" + text + "' where id = " + strconv.Itoa(state.currentSettingsMember+1))

			if err != nil {
				log.Println("Error updating name:", err)
				return
			}
		}
	}

	pronounEntry.OnChanged = func(text string) {
		if state.avatarChangedText {
			_, err := state.db.Exec("update members set pronouns = '" + text + "' where id = " + strconv.Itoa(state.currentSettingsMember+1))

			if err != nil {
				log.Println("Error updating pronouns:", err)
				return
			}
		}
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
		memberForm.Hidden = false

		res, err := state.db.Query("select name, avatar_path, pronouns from members where id=" + strconv.Itoa(i+1))

		if err != nil {
			log.Println("Error getting info from member "+strconv.Itoa(i+1)+":", err)
			return
		}

		defer res.Close()
		var name string
		var avatarUrl sql.NullString
		var pronouns sql.NullString
		res.Next()
		err = res.Scan(&name, &avatarUrl, &pronouns)

		if err != nil {
			log.Println("Error scanning info from member "+strconv.Itoa(i+1)+":", err)
			return
		}

		state.avatarChangedText = false
		nameEntry.Text = name
		nameEntry.Refresh()
		pronounEntry.Text = pronouns.String
		pronounEntry.Refresh()
		avatarEntry.Text = avatarUrl.String
		avatarEntry.Validator = avatarValidator
		avatarEntry.Refresh()
		var avatarURI fyne.URI

		if avatarUrl.Valid {
			if avatarUrl.String[0] == '/' {
				avatarURI = storage.NewFileURI(avatarUrl.String)
			} else {
				avatarURI = storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), avatarUrl.String))
			}

			img, err := LoadImage(avatarURI.Path())

			if err != nil {
				log.Println("Error loading image:", err)
				return
			}

			(*avatarImage).Image = img
			state.avatarSettingsSize.Height = (300.0 * float32(img.Bounds().Max.Y)) / float32(img.Bounds().Max.X)
			updateImageSize()
		} else {
			(*avatarImage).Image = (*canvas.NewImageFromResource(state.defaultAvatar)).Image
			state.avatarSettingsSize.Height = 300.0
			updateImageSize()
		}
	}

	outerRight := container.NewVBox(
		memberForm,
		avatarImage,
		widget.NewButton("Create Member", func() {
			_, err := state.db.Exec("insert into members(name) values('New Member')")

			if err != nil {
				log.Println("Error creating new member:", err)
				return
			}

			newItem := list.CreateItem()
			list.UpdateItem(list.Length()-1, newItem)
		}),
		widget.NewButton("PluralKit Import", func() {
			PKImport(app, window, state)
		}),
		widget.NewButton("SimplyPlural Import", func() {
			SPImport(app, window, state)
		}),
	)

	tabContainer = container.New(
		layout.NewFormLayout(),
		list,
		outerRight,
	)

	return tabContainer
}
