package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed themes/default-themes.json
var defaultThemes string

func main() {
	// Create an instance of the app structure
	app := NewApp()

	appMenu := menu.NewMenu()

	if runtime.GOOS == "darwin" {
		appMenu.Append(menu.AppMenu())
	}

	fileMenu := appMenu.AddSubmenu("File")

	fileMenu.AddText("Settings", keys.CmdOrCtrl(","), func(_ *menu.CallbackData) {
		rt.WindowExecJS(app.ctx, `window.openSettings()`)
	})

	fileMenu.AddSeparator()

	fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		rt.Quit(app.ctx)
	})

	if runtime.GOOS == "darwin" {
		appMenu.Append(menu.EditMenu())
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "GoPlural",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Menu:             appMenu,
		Bind: []interface{}{
			app,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
