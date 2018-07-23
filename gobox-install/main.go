package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/CJ-Jackson/gobox/tool"
)

const mainGo = `
package main

func main() {
	
}
`

const goMod = `module sandbox`

func main() {
	env := tool.GetEnv()

	userConfigFile, err := os.Open(env.ProjectConfigFile())
	if err != nil {
		log.Fatalf("Unable to open gobox.toml: %s", err)
	}

	userConfig := tool.TomlSupplement{}
	_, err = toml.DecodeReader(userConfigFile, &userConfig)
	if err != nil {
		log.Fatalf("Unable to parse gobox.toml: %s", err)
	}

	if userConfig.ProjectInstalls != nil {
		localInstall(env, userConfig)
	}

	if userConfig.Modules != nil {
		externalInstall(env, userConfig)
	}
}

func localInstall(env tool.Env, userConfig tool.TomlSupplement) {
	goBin := "GOBIN=" + tool.FixPath(env.ProjectBinPath())
	for _, install := range userConfig.ProjectInstalls {
		execCommand("vgo", []string{"install", install}, []string{goBin})
	}
}

func externalInstall(env tool.Env, userConfig tool.TomlSupplement) {
	sandBoxLocation := env.SandboxLocation()
	if _, err := os.Stat(sandBoxLocation); os.IsNotExist(err) {
		os.MkdirAll(sandBoxLocation, 0755)
	}
	if _, err := os.Stat(sandBoxLocation + "/main.go"); os.IsNotExist(err) {
		file, err := os.Create(sandBoxLocation + "/main.go")
		if err != nil {
			log.Fatalf("Unable to create main.go: %s", err)
		}
		file.Write([]byte(mainGo))
		file.Close()
	}
	if _, err := os.Stat(sandBoxLocation + "/go.mod"); os.IsNotExist(err) {
		file, err := os.Create(sandBoxLocation + "/go.mod")
		if err != nil {
			log.Fatalf("Unable to create go.mod: %s", err)
		}
		file.Write([]byte(goMod))
		file.Close()
	}
	err := os.Chdir(sandBoxLocation)
	if err != nil {
		log.Fatalf("Unable to change to sandbox directory: %s", err)
	}

	binPath := env.ProjectBinPath()
	for _, module := range userConfig.Modules {
		execCommand("vgo", []string{"get", "-d", module.RepoAndTag()}, []string{})
		moduleBinPath := binPath
		if module.BinPath != "" {
			moduleBinPath += "/" + strings.Trim(module.BinPath, "/")
		}
		for _, install := range module.Installs {
			output := moduleBinPath
			if install == "" || install == "." {
				output += "/" + tool.FixOutput(path.Base(module.Repo))
				install = module.Repo
			} else {
				output += "/" + tool.FixOutput(path.Base(install))
				install = module.Repo + "/" + strings.Trim(install, "/")
			}
			execCommand("vgo", []string{"build", "-o", tool.FixPath(output), "-i", install}, []string{})
		}
	}
}

func execCommand(name string, args []string, environ []string) {
	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ(), environ...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Unable to run command: %s", err)
	}
}
