package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/CJ-Jackson/gobox/tool"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need at least one argument")
	}

	env := tool.GetEnv()
	file := env.ProjectBinPath() + "/" + strings.Trim(os.Args[1], "/")
	tool.AppendPath(tool.FixPath(path.Dir(file)))

	cmd := exec.Command(path.Base(file), os.Args[2:]...)
	cmd.Env = os.Environ()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get working directory: %s", err)
	}
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
