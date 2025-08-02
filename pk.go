package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func PKImport(parent fyne.Window) {
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

		fmt.Println(idEntry.Text)
	}, parent)
}
