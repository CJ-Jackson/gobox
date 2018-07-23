package tool

import (
	"os"
	"runtime"
)

func userHomeDir() string {
	env := "HOME"
	switch runtime.GOOS {
	case "windows":
		env = "USERPROFILE"
	case "plan9":
		env = "home"
	}
	return os.Getenv(env)
}

func PrependPath(path string) {
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", path+string(os.PathListSeparator)+oldPath)
}
