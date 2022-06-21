package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

type routes int64

const (
	RouteLogin routes = iota
	RouteTwoFactorAuth
)

type Router struct {
	currentRoute routes
	window       fyne.Window
	app          fyne.App
	vrc          *vrchat.VRChat
}

func NewRouter(a fyne.App, w fyne.Window, vrc *vrchat.VRChat) *Router {
	return &Router{
		currentRoute: RouteLogin,
		app:          a,
		window:       w,
		vrc:          vrc,
	}
}

func (r *Router) SetRoute(route routes) {
	var view fyne.CanvasObject
	switch route {
	case RouteLogin:
		view = Login(r.app, r.window, r.vrc, r)
	}
	if view != nil {
		r.window.SetContent(view)
	} else {
		fyne.LogError(fmt.Sprintf("Unexpected Route: %d", route), nil)
	}
}
