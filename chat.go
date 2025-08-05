package main

import (
	"database/sql"
	"log"
	"path"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func makeChatScreen(app fyne.App, window fyne.Window, state *AppState) fyne.CanvasObject {
	memberList := widget.NewList(
		func() int {
			res, err := state.db.Query("select count(*) from members")

			if err != nil {
				log.Fatal("Error getting count of members:", err)
			}

			res.Next()
			var result int
			err = res.Scan(&result)

			if err != nil {
				log.Fatal("Error scanning count of members:", err)
			}

			return result
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				canvas.NewImageFromResource(state.defaultAvatar),
				widget.NewLabel("template"),
			)
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			var avatarUri fyne.URI

			res, err := state.db.Query("select name, avatar_url from members where id = " + strconv.Itoa(id))

			if err != nil {
				log.Println("Error getting avatar url:", err)
				return
			}

			res.Next()
			var avatarUrl sql.NullString
			var name string
			err = res.Scan(&name, &avatarUrl)

			if err != nil {
				log.Println("Error scanning avatar url:", err)
				return
			}

			if avatarUrl.Valid {
				if avatarUrl.String[0] == '/' {
					avatarUri = storage.NewFileURI(avatarUrl.String)
				} else {
					avatarUri = storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), avatarUrl.String))
				}

				o.(*fyne.Container).Objects[0].(*canvas.Image).Image = (*canvas.NewImageFromURI(avatarUri)).Image
			} else {
				o.(*fyne.Container).Objects[0].(*canvas.Image).Image = (*canvas.NewImageFromResource(state.defaultAvatar)).Image
			}

			o.(*fyne.Container).Objects[1].(*widget.Label).Text = name
		},
	)

	memberList.OnSelected = func(id widget.ListItemID) {
		(*state).selectedMember = id
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
	return container.NewVBox()
}
