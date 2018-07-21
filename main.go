package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/CJ-Jackson/gobox/tool"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need at least one argument")
	}

	env := tool.GetEnv()
	file := env.ProjectBinPath() + "/" + strings.Trim(os.Args[1], "/")

	cmd := exec.Command(tool.FixPath(file), os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
