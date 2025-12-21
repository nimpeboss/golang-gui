package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"advanced-gui/state"
)

func Logs(s *state.AppState) fyne.CanvasObject {
	logs := widget.NewMultiLineEntry()
	logs.Disable()

	go func() {
		for {
			logs.SetText(strings.Join(s.Logs, "\n"))
		}
	}()

	return logs
}