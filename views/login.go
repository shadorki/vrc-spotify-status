package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

var logo = canvas.NewImageFromFile("./images/logo.png")

func Login(a fyne.App, w fyne.Window, vrc *vrchat.VRChat, r *Router) fyne.CanvasObject {
	logo.SetMinSize(fyne.NewSize(305, 142))
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
			fmt.Println(vrc)
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
		container.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		form,
		loading,
		container.NewHBox(layout.NewSpacer(), errorMessage, layout.NewSpacer()),
	)
}
