package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

func TwoFactorAuth(l *canvas.Image, a fyne.App, w fyne.Window, vrc *vrchat.VRChat, r *Router) fyne.CanvasObject {
	twoFactorAuthEnabledMessage := widget.NewLabel("You have Two Factor Authentication Enabled")
	errorMessage := widget.NewLabel("")
	authCodeEntry := widget.NewPasswordEntry()
	loading := widget.NewProgressBarInfinite()
	loading.Hidden = true
	form := widget.NewForm(
		widget.NewFormItem("Auth Code", authCodeEntry),
	)
	form.SubmitText = "Authenticate"
	form.OnCancel = func() {
		errorMessage.Text = ""
		errorMessage.Refresh()
		authCodeEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		if authCodeEntry.Text == "" {
			return
		}
		errorMessage.Text = ""
		loading.Hidden = false
		form.Disable()
		err := vrc.TwoFactorAuthenticate(authCodeEntry.Text)
		if err != nil {
			authCodeEntry.Text = ""
			errorMessage.Text = err.Error()
			errorMessage.Refresh()
		} else {
			r.SetRoute(RouteHome)
		}
		loading.Hidden = true
		form.Enable()
	}
	authCodeEntry.OnChanged = func(_ string) {
		errorMessage.Text = ""
	}
	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), l, layout.NewSpacer()),
		container.NewHBox(layout.NewSpacer(), twoFactorAuthEnabledMessage, layout.NewSpacer()),
		form,
		loading,
		container.NewHBox(layout.NewSpacer(), errorMessage, layout.NewSpacer()),
		layout.NewSpacer(),
	)
}
