package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SPImport(app fyne.App, parent fyne.Window, memberData *[]Member) {
	idEntry := widget.NewEntry()
	tokenEntry := widget.NewEntry()
	warningLabel := widget.NewLabel("WARNING: This will erase all system data to import from PluralKit. Make sure you know what you're doing before you click confirm.")
	warningLabel.Wrapping = fyne.TextWrapWord
	info1Label := widget.NewLabel("Enter your account ID below (Settings > Account > Account Settings)")
	info1Label.Wrapping = fyne.TextWrapWord
	info2Label := widget.NewLabel("Enter your account token below (Settings > Account > Tokens)")
	info2Label.Wrapping = fyne.TextWrapWord

	dataContainer := container.NewVBox(
		warningLabel,
		info1Label,
		idEntry,
		info2Label,
		tokenEntry,
	)

	dialog.ShowCustomConfirm("Import from SimplyPlural", "Confirm", "Cancel", dataContainer, func(result bool) {
		if !result {
			return
		}

		req, err := http.NewRequest("GET", "https://api.apparyllis.com/v1/members/"+idEntry.Text, nil)

		if err != nil {
			log.Println("Error in creating request:", err)
			return
		}

		req.Header.Add("Authorization", tokenEntry.Text)

		res, err := new(http.Client).Do(req)

		if err != nil {
			log.Println("Error sending request:", err)
			return
		}

		defer res.Body.Close()
		var members spMembers

		err = json.NewDecoder(res.Body).Decode(&members)

		if err != nil {
			log.Println("Error parsing system data:", err)
			return
		}

		(*memberData) = make([]Member, len(members))

		for i, v := range members {
			(*memberData)[i].Name = v.Content.Name

			if v.Content.Avatar != nil {
				uri, err := loadAvatar(*v.Content.Avatar, v.Id, app)
				time.Sleep(1 * time.Second)

				if err == nil {
					(*memberData)[i].Avatar = uri
				}
			}

			if v.Content.Pronouns != nil {
				(*memberData)[i].Pronouns = *v.Content.Pronouns
			}
		}

		log.Println("Imported system from SimplyPlural")
	}, parent)
}

type spMembers []spMember

type spMember struct {
	Id      string          `json:"id"`
	Content spMemberContent `json:"content"`
}

type spMemberContent struct {
	Name     string  `json:"name"`
	Avatar   *string `json:"avatarUrl,omitempty"`
	Pronouns *string `json:"pronouns,omitempty"`
}
