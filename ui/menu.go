package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"advanced-gui/state"
	"advanced-gui/utils"
)

func BuildUI(a fyne.App, s *state.AppState) {
	w := a.NewWindow("Advanced Go GUI (No DB)")
	w.Resize(fyne.NewSize(900, 600))

	tabs := container.NewAppTabs(
		container.NewTabItem("Dashboard", Dashboard(a, s)),
		container.NewTabItem("Tasks", Tasks(s)),
		container.NewTabItem("Logs", Logs(s)),
	)

	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Save Config", func() {
				utils.SaveConfig(s)
			}),
			fyne.NewMenuItem("Load Config", func() {
				if err := utils.LoadConfig(s); err != nil {
					dialog.ShowError(err, w)
				}
			}),
			fyne.NewMenuItem("Quit", w.Close),
		),
	)

	w.SetMainMenu(menu)
	w.SetContent(tabs)
	w.ShowAndRun()
}
