package config

import (
	"errors"
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
)

type ProjectConfig struct {
	Repo    string `toml:"repo"`
	Project string `toml:"project"`
	Env     string `toml:"env"`
	Branch  string `toml:"branch"`
	Tag     string `toml:"tag"`
	Push    bool   `toml:"push"`
}

// Load finds fnx.toml or .fnx/default.toml and parses it.
func Load(repoOverride string) (*ProjectConfig, error) {
	if _, err := os.Stat("fnx.toml"); err == nil {
		return parseFile("fnx.toml", repoOverride)
	}
	if _, err := os.Stat(".fnx"); err == nil {
		return parseFile(filepath.Join(".fnx", "default.toml"), repoOverride)
	}
	return nil, errors.New("no fnx config found (fnx.toml or .fnx/default.toml)")
}

func parseFile(path, repoOverride string) (*ProjectConfig, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg ProjectConfig
	if err := toml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	if repoOverride != "" {
		cfg.Repo = repoOverride
	}
	return &cfg, nil
}
