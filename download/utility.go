package download

import (
	"math/rand"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
)

func Random_id(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	f := ""
	for i := 0; i < length; i++ {
		f += string(characters[rand.Intn(len(characters))])
	}
	return f
}

func Open_file_explorer(path string) { // really want this to multi platform
	if runtime.GOOS == "windows" {
		exec.Command("explorer.exe", path).Start()
	} else if runtime.GOOS == "linux" {
		exec.Command("xdg-open", path).Start()
	} else if runtime.GOOS == "darwin" {
		exec.Command("open", path).Start()
	}
}

func Check_Beatmap_ID(beatmap_id string) bool {
	// TODO: Check the url too
	_, err := strconv.Atoi(beatmap_id)
	if err != nil {
		return false
	}
	// now try check if its url
	url, err := url.Parse(beatmap_id)
	if err != nil {
		return false
	}
	if url.Host != "osu.ppy.sh" {
		return false
	}
	// now iterate through the path and find the beatmap id
	for _, v := range url.Path {
		// try convert to int
		_, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		} else {
			return true
		}
	}
	return true
}
