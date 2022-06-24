package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/uzair-ashraf/vrc-spotify-status/spotify"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

type routes int64

const (
	RouteLogin routes = iota
	RouteTwoFactorAuth
	RouteHome
)

type Router struct {
	currentRoute routes
	window       fyne.Window
	app          fyne.App
	vrc          *vrchat.VRChat
	spotify      *spotify.Spotify
	logo         *canvas.Image
}

func NewRouter(l *canvas.Image, a fyne.App, w fyne.Window, vrc *vrchat.VRChat, s *spotify.Spotify) *Router {
	return &Router{
		currentRoute: RouteLogin,
		app:          a,
		window:       w,
		vrc:          vrc,
		spotify:      s,
		logo:         l,
	}
}

func (r *Router) SetRoute(route routes) {
	var view fyne.CanvasObject
	switch route {
	case RouteLogin:
		view = Login(r.logo, r.app, r.window, r.vrc, r)
	case RouteTwoFactorAuth:
		view = TwoFactorAuth(r.logo, r.app, r.window, r.vrc, r)
	case RouteHome:
		view = Home(r.logo, r.app, r.window, r.vrc, r.spotify, r)
	}
	if view != nil {
		r.window.SetContent(view)
	} else {
		fyne.LogError(fmt.Sprintf("Unexpected Route: %d", route), nil)
	}
}
