package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func PKImport(app fyne.App, parent fyne.Window, state *AppState) {
	idEntry := widget.NewEntry()
	warningLabel := widget.NewLabel("WARNING: This will erase all system data to import from PluralKit. Make sure you know what you're doing before you click confirm.")
	warningLabel.Wrapping = fyne.TextWrapWord
	infoLabel := widget.NewLabel("Enter your 5 or 6 character PluralKit ID below (shown at the very bottom of the pk;s command):")
	infoLabel.Wrapping = fyne.TextWrapWord

	dataContainer := container.NewVBox(
		warningLabel,
		infoLabel,
		idEntry,
	)

	dataContainer.Resize(fyne.NewSize(440, dataContainer.MinSize().Height))

	dialog.ShowCustomConfirm("Import from PluralKit", "Confirm", "Cancel", dataContainer, func(result bool) {
		if !result {
			return
		}

		resp, err := http.Get("https://api.pluralkit.me/v2/systems/" + idEntry.Text + "/members")

		if err != nil {
			log.Println("Error getting system data:", err)
			return
		}

		defer resp.Body.Close()
		var members pkMembers

		err = json.NewDecoder(resp.Body).Decode(&members)

		if err != nil {
			log.Println("Error parsing system data:", err)
			return
		}

		_, err = state.db.Exec(`
		drop table members;
		
		create table members (
			id integer not null primary key,
			name text unique not null,
			pronouns text,
			avatar_path text,
			color text,
			proxy_tags text
		);`)

		if err != nil {
			log.Println("Error clearing members table:", err)
			return
		}

		createStringValue := func(val *string) string {
			if val == nil {
				return "null"
			}

			return "'" + (*val) + "'"
		}

		for _, v := range members {
			var name string
			var pronouns string
			var avatar string
			var proxies string

			name = v.Name
			pronouns = createStringValue(v.Pronouns)

			if v.Avatar != nil {
				uri, err := loadAvatar(*v.Avatar, v.Id, app)
				time.Sleep(1 * time.Second)

				if err == nil {
					avatar = "'" + uri + "'"
				}
			} else {
				avatar = "null"
			}

			data, err := json.Marshal(v.ProxyTags)

			if err != nil {
				log.Println("Error marshaling proxy tags:", err)
				continue
			}

			proxies = string(data)

			_, err = state.db.Exec("insert into members(name, pronouns, avatar_url, proxy_tags) values(" + name + ", " + pronouns + ", " + avatar + ", '" + proxies + "')")

			if err != nil {
				log.Println("Error inserting new member into table:", err)
				continue
			}
		}

		log.Printf("Imported system %s from PluralKit", idEntry.Text)
	}, parent)
}

type pkMembers []pkMember

type pkMember struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Pronouns  *string      `json:"pronouns,omitempty"`
	ProxyTags []pkProxyTag `json:"proxy_tags"`
	Avatar    *string      `json:"avatar_url,omitempty"`
}

type pkProxyTag struct {
	Prefix *string `json:"prefix,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
}

func loadAvatar(uri string, id string, app fyne.App) (string, error) {
	log.Printf("Loading avatar %s from %s...", id, uri)

	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		log.Println("Error making request:", err)
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := new(http.Client).Do(req)

	if err != nil {
		log.Println("Error downloading avatar:", err)
		return "", err
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error reading avatar data:", err)
		return "", err
	}

	imageURI := storage.NewFileURI(path.Join(app.Storage().RootURI().Path(), id+path.Ext(uri)))
	f, err := storage.Writer(imageURI)

	if err != nil {
		log.Println("Error opening file for writing:", err)
		return "", err
	}

	defer f.Close()
	val, err := f.Write(data)

	if err != nil {
		log.Println("Error writing file:", err)
		return "", err
	}

	fmt.Println(val)
	return imageURI.Name(), nil
}
