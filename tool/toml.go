package tool

type TomlSupplement struct {
	ProjectInstalls []string     `toml:"projectInstalls"`
	Modules         []TomlModule `toml:"module"`
}

type TomlModule struct {
	Repo     string   `toml:"repo"`
	Tag      string   `toml:"tag"`
	BinPath  string   `toml:"binPath"`
	Installs []string `toml:"installs"`
}

func (t TomlModule) RepoAndTag() string {
	if t.Tag == "" {
		return t.Repo
	}

	return t.Repo + "@" + t.Tag
}
