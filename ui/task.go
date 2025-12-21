package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"advanced-gui/state"
)

func Tasks(s *state.AppState) fyne.CanvasObject {
	return widget.NewTable(
		func() (int, int) {
			return len(s.Tasks) + 1, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			if id.Row == 0 {
				labels := []string{"ID", "Task", "Status"}
				label.SetText(labels[id.Col])
				return
			}
			task := s.Tasks[id.Row-1]
			switch id.Col {
			case 0:
				label.SetText(strconv.Itoa(task.ID))
			case 1:
				label.SetText(task.Name)
			case 2:
				label.SetText(task.Status)
			}
		},
	)
}
