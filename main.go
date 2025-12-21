package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	"advanced-gui/state"
	"advanced-gui/ui"
)

func main() {
	a := app.NewWithID("com.example.advancedgui")
	a.Settings().SetTheme(theme.DefaultTheme())

	appState := state.NewAppState()
	ui.BuildUI(a, appState)
}
