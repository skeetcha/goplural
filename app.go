package main

import (
	"context"
	"encoding/json"
	"os"
	"path"
)

// App struct
type App struct {
	ctx          context.Context
	currentTheme string
}

// NewApp creates a new App application struct
func NewApp() *App {
	app := &App{}
	app.currentTheme = "superhero"

	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetCurrentTheme() string {
	return a.currentTheme
}

func (a *App) GetCustomThemes() (out map[string]string, err error) {
	out = make(map[string]string)

	var exe string
	var oerr error

	if exe, oerr = os.Executable(); oerr != nil {
		err = oerr
		return
	}

	if _, oerr = os.Stat(path.Join(path.Dir(exe), "data")); os.IsNotExist(oerr) {
		if oerr = os.Mkdir(path.Join(path.Dir(exe), "data"), 0755); oerr != nil {
			err = oerr
			return
		}

		if oerr = os.Mkdir(path.Join(path.Dir(exe), "data", "themes"), 0755); oerr != nil {
			err = oerr
			return
		}

		f, oerr := os.Create(path.Join(path.Dir(exe), "data", "themes", "themes.json"))

		if oerr != nil {
			err = oerr
			return
		}

		_, oerr = f.Write([]byte("[]"))

		if oerr != nil {
			err = oerr
			return
		}

		oerr = f.Close()

		if oerr != nil {
			err = oerr
			return
		}
	}

	f, oerr := os.Open(path.Join(path.Dir(exe), "data", "themes", "themes.json"))

	if oerr != nil {
		err = oerr
		return
	}

	stats, oerr := os.Stat(f.Name())

	if oerr != nil {
		err = oerr
		return
	}

	themeBData := make([]byte, stats.Size())

	_, oerr = f.Read(themeBData)

	if oerr != nil {
		err = oerr
		return
	}

	oerr = f.Close()

	if oerr != nil {
		err = oerr
		return
	}

	var themes []string

	oerr = json.Unmarshal(themeBData, &themes)

	if oerr != nil {
		err = oerr
		return
	}

	for _, theme := range themes {
		f, oerr = os.Open(path.Join(path.Dir(exe), "data", "themes", theme+".css"))

		if oerr != nil {
			err = oerr
			return

		}

		stats, oerr = os.Stat(f.Name())

		if oerr != nil {
			err = oerr
			return
		}

		customData := make([]byte, stats.Size())

		_, oerr = f.Read(customData)

		if oerr != nil {
			err = oerr
			return
		}

		oerr = f.Close()

		if oerr != nil {
			err = oerr
			return
		}

		out[theme] = string(customData)
	}

	return
}

func (a *App) GetFAIcon(id string) string {
	if id == "xmark" {
		return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512" class="icon"><!--!Font Awesome Free v7.0.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M55.1 73.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L147.2 256 9.9 393.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0L192.5 301.3 329.9 438.6c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L237.8 256 375.1 118.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L192.5 210.7 55.1 73.4z"/></svg>`
	}

	return ""
}
