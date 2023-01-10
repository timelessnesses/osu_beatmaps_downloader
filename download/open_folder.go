package download

import (
	"math/rand"
	"os/exec"
	"runtime"
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
