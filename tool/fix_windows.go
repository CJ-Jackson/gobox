// +build windows

package tool

import (
	"os"
	"strings"
)

func FixPath(str string) string { return strings.Replace(str, "/", "\\", -1) }

func RevFixPath(str string) string { return strings.Replace(str, "\\", "/", -1) }

func PrependPath(path string) {
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", path+";"+oldPath)
}

func FixOutput(output string) string { return output + ".exe" }
