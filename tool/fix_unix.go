// +build !windows

package tool

import "os"

func FixPath(str string) string { return str }

func RevFixPath(str string) string { return str }

func AppendPath(path string) {
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", path+":"+oldPath)
}
