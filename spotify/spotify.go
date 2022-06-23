package spotify

import (
	"os/exec"
	"strings"
	"time"

	"fyne.io/fyne/v2"
)

type Spotify struct {
	hasStarted   bool
	isRunning    bool
	subscribers  []func(nextSong string)
	lastSongName string
}

func New() *Spotify {
	return &Spotify{
		hasStarted:   false,
		isRunning:    false,
		subscribers:  []func(nextSong string){},
		lastSongName: "",
	}
}

func (s *Spotify) IsRunning() bool {
	return s.isRunning
}

func (s *Spotify) HasStarted() bool {
	return s.hasStarted
}

func (s *Spotify) HasSubscribers() bool {
	return len(s.subscribers) != 0
}

func (s *Spotify) Start() {
	s.isRunning = true
	if s.hasStarted {
		return
	}
	go s.run()
	s.hasStarted = true
}

func (s *Spotify) Stop() {
	s.isRunning = false
}

func (s *Spotify) run() {
	for {
		if !s.IsRunning() || !s.HasSubscribers() {
			time.Sleep(time.Second * 3)
			continue
		}
		songName, err := s.getCurrentSpotifySong()
		if err != nil {
			fyne.LogError("Unable to retrieve song name", err)
			time.Sleep(time.Second * 3)
			continue
		}
		if songName == s.lastSongName {
			continue
		}
		s.lastSongName = songName
		s.Dispatch(songName)
		time.Sleep(time.Second * 3)
	}
}

func (s *Spotify) Dispatch(songName string) {
	for _, sub := range s.subscribers {
		sub(songName)
	}
}

// Takes a function to be called when a new spotify song is discovered
// Returns a function to be called when you would like to unsubscribe
func (s *Spotify) Subscribe(listener func(nextSong string)) func() {
	s.subscribers = append(s.subscribers, listener)
	return func() {
		newSubscribers := []func(nextSong string){}
		for _, sub := range s.subscribers {
			if &listener != &sub {
				newSubscribers = append(newSubscribers, sub)
			}
		}
		s.subscribers = newSubscribers
	}
}

func (s *Spotify) getCurrentSpotifySong() (string, error) {
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
			s = strings.TrimSpace(s)
			if strings.HasPrefix(s, "Spotify Premium") {
				continue
			}
			songName = s
		}
	}
	return songName, nil
}
