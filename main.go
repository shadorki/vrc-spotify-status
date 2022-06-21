package main

import (
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/uzair-ashraf/vrc-spotify-status/views"
	"github.com/uzair-ashraf/vrc-spotify-status/vrchat"
)

var lastSongName string

const statusDescriptionLimit = 31

func main() {
	a := app.New()
	w := a.NewWindow("VRCSS")
	w.Resize(fyne.NewSize(400, 400))
	vrc := vrchat.New()
	router := views.NewRouter(a, w, vrc)
	router.SetRoute(views.RouteLogin)
	w.ShowAndRun()
}

func getCurrentSpotifySong() (string, error) {
	cmd := exec.Command("powershell", "-nologo", "-noprofile", `
		Get-Process | Where {
			$_.MainWindowTitle
		} |
		Select-Object ProcessName, MainWindowTitle
	`)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	defer stdin.Close()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(out), "\n")
	var songName string
	for _, l := range lines {
		if strings.HasPrefix(l, "Spotify") {
			s := strings.Replace(l, "Spotify", "", 1)
			s = strings.Trim(s, " ")
			s = strings.Trim(s, "\n")
			if strings.HasPrefix(s, "Spotify Premium") {
				continue
			}
			songName = s
		}
	}
	return songName, nil
}
