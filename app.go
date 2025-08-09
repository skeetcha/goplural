package main

import (
	"context"
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

func (a *App) GetFAIcon(id string) string {
	if id == "xmark" {
		return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512" class="icon"><!--!Font Awesome Free v7.0.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M55.1 73.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3L147.2 256 9.9 393.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0L192.5 301.3 329.9 438.6c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L237.8 256 375.1 118.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L192.5 210.7 55.1 73.4z"/></svg>`
	}

	return ""
}
