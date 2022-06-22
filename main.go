package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/uzair-ashraf/vrc-spotify-status/spotify"
	"github.com/uzair-ashraf/vrc-spotify-status/views"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

const statusDescriptionLimit = 31

func main() {
	s := spotify.New()
	a := app.New()
	w := a.NewWindow("VRCSS")
	w.Resize(fyne.NewSize(400, 400))
	vrc := vrchat.New()
	router := views.NewRouter(a, w, vrc, s)
	router.SetRoute(views.RouteLogin)
	w.ShowAndRun()
}
