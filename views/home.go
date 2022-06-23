package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/uzair-ashraf/vrc-spotify-status/spotify"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

func Home(a fyne.App, w fyne.Window, vrc *vrchat.VRChat, s *spotify.Spotify, r *Router) fyne.CanvasObject {
	Logo.SetMinSize(fyne.NewSize(305, 142))
	currentSong := widget.NewLabel("")
	btn := widget.NewButton("Start", func() {})
	btn.OnTapped = func() {
		if s.IsRunning() {
			btn.Text = "Start"
			s.Stop()
		} else {
			btn.Text = "Stop"
			s.Start()
		}
	}
	s.Subscribe(func(nextSong string) {
		currentSong.Text = nextSong
		currentSong.Refresh()
		vrc.SetStatus(nextSong)
	})
	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), Logo, layout.NewSpacer()),
		container.NewHBox(layout.NewSpacer(), currentSong, layout.NewSpacer()),
		container.NewHBox(layout.NewSpacer(), btn, layout.NewSpacer()),
		layout.NewSpacer(),
	)
}
