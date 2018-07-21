// +build windows

package tool

import (
	"os"
	"strings"
)

func FixPath(str string) string { return strings.Replace(str, "/", "\\", -1) }

func RevFixPath(str string) string { return strings.Replace(str, "\\", "/", -1) }

func AppendPath(path string) {
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", path+";"+oldPath)
}
