package download

import (
	"io"
	"net/http"
	"os"

	"fyne.io/fyne/widget"
)

type Sources string

const (
	Kitsu Sources = "https://kitsu.moe/api/d/"
	Chimu Sources = "https://api.chimu.moe/v1/download/"
)

var current_source Sources

func Set_Download_Source(source Sources) {
	current_source = source
}

func (s Sources) String() string {
	if s == "Kitsu" {
		return "https://kitsu.moe/api/d/"
	} else if s == "Chimu" {
		return "https://api.chimu.moe/v1/download/"
	} else {
		panic("Unknown source")
	}
}

func Download_Beatmap(beatmap_id string, path string, status_text *widget.Label) (string, error) {
	status_text.SetText("Status: Fetching osz from " + current_source.String() + beatmap_id + ". (" + path + ")")
	req, err := http.Get(current_source.String() + beatmap_id)
	print(current_source.String() + beatmap_id)
	if err != nil {
		status_text.SetText("Status: Failed to fetch osz from source: " + current_source.String() + "\n" + err.Error())
		return "", err
	}
	defer req.Body.Close()
	print(path)
	file, _ := os.Create(path + "/" + Random_id(10) + ".osz")
	defer file.Close()
	io.Copy(file, req.Body)
	return file.Name(), nil
}
