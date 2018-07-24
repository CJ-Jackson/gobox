// +build windows

package tool

import (
	"strings"
)

func FixPath(str string) string      { return strings.Replace(str, "/", "\\", -1) }
func RevFixPath(str string) string   { return strings.Replace(str, "\\", "/", -1) }
func FixOutput(output string) string { return output + ".exe" }
