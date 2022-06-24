package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

func Login(l *canvas.Image, a fyne.App, w fyne.Window, vrc *vrchat.VRChat, r *Router) fyne.CanvasObject {
	errorMessage := widget.NewLabel("")
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	loading := widget.NewProgressBarInfinite()
	loading.Hidden = true
	form := widget.NewForm(
		widget.NewFormItem("Username", usernameEntry),
		widget.NewFormItem("Password", passwordEntry),
	)
	form.SubmitText = "Login"
	form.OnCancel = func() {
		errorMessage.Text = ""
		errorMessage.Refresh()
		usernameEntry.Text = ""
		passwordEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		if usernameEntry.Text == "" || passwordEntry.Text == "" {
			return
		}
		errorMessage.Text = ""
		loading.Hidden = false
		form.Disable()
		err := vrc.Login(usernameEntry.Text, passwordEntry.Text)
		if err != nil {
			errorMessage.Text = err.Error()
			errorMessage.Refresh()
		} else {
			if vrc.IsLoggedIn() {
				if vrc.IsTwoFactorAuthEnabled() {
					r.SetRoute(RouteTwoFactorAuth)
				} else {
					r.SetRoute(RouteHome)
				}
			}
		}
		loading.Hidden = true
		form.Enable()
	}
	usernameEntry.OnChanged = func(_ string) {
		errorMessage.Text = ""
	}
	passwordEntry.OnChanged = func(_ string) {
		errorMessage.Text = ""
	}
	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), l, layout.NewSpacer()),
		form,
		loading,
		container.NewHBox(layout.NewSpacer(), errorMessage, layout.NewSpacer()),
		layout.NewSpacer(),
	)
}
