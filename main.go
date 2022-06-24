package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/uzair-ashraf/vrc-spotify-status/spotify"
	"github.com/uzair-ashraf/vrc-spotify-status/views"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

const (
	LOGO_URL = "https://raw.githubusercontent.com/uzair-ashraf/vrc-spotify-status/master/images/logo.png"
	VERSION  = "v0.0.1"
)

func main() {
	u, _ := fyne.LoadResourceFromURLString(LOGO_URL)
	l := canvas.NewImageFromResource(u)
	l.SetMinSize(fyne.NewSize(305, 142))
	s := spotify.New()
	a := app.New()
	w := a.NewWindow("VRCSS")
	w.Resize(fyne.NewSize(400, 400))
	vrc := vrchat.New()
	router := views.NewRouter(l, a, w, vrc, s)
	router.SetRoute(views.RouteLogin)
	w.ShowAndRun()
}
