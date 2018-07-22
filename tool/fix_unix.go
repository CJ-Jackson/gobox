// +build !windows

package tool

import "os"

func FixPath(str string) string { return str }

func RevFixPath(str string) string { return str }

func PrependPath(path string) {
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", path+":"+oldPath)
}

func FixOutput(output string) string { return output }
