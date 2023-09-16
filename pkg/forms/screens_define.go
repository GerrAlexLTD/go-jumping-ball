package screens

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LaunchMainMenu() {
	base := app.New()
	mainWindow := base.NewWindow("Jumping.io")
	mainWindow.CenterOnScreen()
	mainWindow.FullScreen()
	mainWindow.SetContent(container.NewHBox(
		container.NewVBox(
			widget.NewLabel("Jumping.io"),
		),
		container.NewHBox(
			widget.NewButton("Start", func() {
				return
			}),
			widget.NewButton("Settings", func() {

			}),
		),
	))
	mainWindow.ShowAndRun()
}
