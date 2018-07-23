package tool

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

type Env struct {
	GoMod string `json:"GOMOD"`
}

func userHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func (e Env) dirPath() string { return path.Dir(RevFixPath(e.GoMod)) }

func (e Env) ProjectBinPath() string { return e.dirPath() + "/bin" }

func (e Env) ProjectConfigFile() string { return e.dirPath() + "/gobox.toml" }

func (e Env) SandboxLocation() string {
	homeDir := RevFixPath(userHomeDir())
	hash := sha256.New()
	hash.Write([]byte(e.GoMod))
	return fmt.Sprintf("%s/.gobox/%s/sandbox", homeDir, base64.URLEncoding.EncodeToString(hash.Sum(nil)))
}

func GetEnv() Env {
	buf := &bytes.Buffer{}

	cmd := exec.Command("vgo", "env", "-json")
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Could not run vgo command: %s", err)
	}

	env := Env{}
	err = json.NewDecoder(buf).Decode(&env)
	if err != nil {
		log.Fatalf("Could not parse json: %s", err)
	}

	if env.GoMod == "" {
		log.Fatal("Not a go module.")
	}

	return env
}
