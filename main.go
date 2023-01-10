package main

import (
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"golang.design/x/clipboard"

	"github.com/sqweek/dialog"
	"github.com/timelessnesses/osu_beatmaps_downloader/download"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range c {
			os.Exit(0)
		}
	}()
	a := app.New()
	w := a.NewWindow("A osu! beatmap downloader written in Fyne and Go")
	selections := []string{"Kitsu", "Chimu"}
	selected := ""
	download_source := widget.NewRadioGroup(selections, func(b string) {
		selected = b
	})
	chosen_id := string(clipboard.Read(clipboard.FmtText))
	chosen_id = convert(chosen_id)
	beatmap_id := widget.NewEntry()
	beatmap_id.SetPlaceHolder("Beatmap ID")
	beatmap_id.SetText(chosen_id)
	status_text := widget.NewLabel("Status: Doing nothing (yet)")
	button := widget.NewButton("Download", func() {
		a, err := dialog.Directory().Title("Select a folder to save the beatmap").Browse()
		if err != nil {
			status_text.SetText("Status: Failed to select a folder to save the beatmap.\n" + err.Error())
			return
		}
		if selected != "Kitsu" && selected != "Chimu" {
			status_text.SetText("Status: Source isn't selected")
			return
		}
		download.Set_Download_Source(download.Sources(selected))
		if beatmap_id.Text == "" {
			status_text.SetText("Status: Beatmap ID isn't entered")
			return
		}
		if !download.Check_Beatmap_ID(beatmap_id.Text) {
			status_text.SetText("Status: Beatmap ID isn't valid")
		}
		b, err := download.Download_Beatmap(beatmap_id.Text, a, status_text)
		if err != nil {
			status_text.SetText("Status: Failed to download the beatmap.\n" + err.Error())
			return
		}
		status_text.SetText("Status: Downloaded the beatmap to " + b)
		download.Open_file_explorer(b)

	})
	w.SetContent(container.NewVBox(
		widget.NewLabel("Select a download source:"),
		download_source,
		widget.NewLabel("Enter a beatmap id:"),
		beatmap_id,
		button,
		status_text,
	))
	w.Show()
	a.Run()
}

func convert(thing string) string {
	// try convert to int
	_, err := strconv.Atoi(thing)
	if err != nil {
		return ""
	}
	// well that failed. try convert this to url object then let's check its domain
	url, err := url.Parse(thing)
	if err != nil {
		return ""
	}
	// check if the domain is osu.ppy.sh
	if url.Host != "osu.ppy.sh" {
		return ""
	}
	// now iterate through the path and find the beatmap id
	for _, v := range url.Path {
		// try convert to int
		_, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		} else {
			return string(v)
		}
	}
	return ""
}
