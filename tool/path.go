package tool

import "os"

func PrependPath(path string) {
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", path+string(os.PathListSeparator)+oldPath)
}
