package main

import (
	"github.com/skeetcha/goplural/themes"
)

type GoPlural struct {
	themes []themes.Theme
}

func main() {
	appData := GoPlural{
		themes: themes.SetupThemes(),
	}
}
