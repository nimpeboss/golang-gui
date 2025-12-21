package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"advanced-gui/state"
)

func Dashboard(a fyne.App, s *state.AppState) fyne.CanvasObject {
	status := widget.NewLabel("Idle")
	bar := widget.NewProgressBar()

	btn := widget.NewButton("Start job", func() {
		status.SetText("Running...")
		bar.SetValue(0)

		go func() {
			for i := 0; i <= 100; i += 10 {
				time.Sleep(300 * time.Millisecond)
				progress := float64(i) / 100

				bar.SetValue(progress)

				s.AddLog("Progress updated")
			}
			status.SetText("Completed")
		}()
	})

	return container.NewVBox(status, bar, btn)
}
