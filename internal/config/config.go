package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type ProjectConfig struct {
	Repo    string `json:"repo"`
	Project string `json:"project"`
	Env     string `json:"env"`
	Branch  string `json:"branch"`
	Tag     string `json:"tag"`
	Push    bool   `json:"push"`
}

func Load(repoOverride string) (*ProjectConfig, error) {
	if _, err := os.Stat("fnx.json"); err == nil {
		return parseFile("fnx.json", repoOverride)
	}
	if _, err := os.Stat(".fnx"); err == nil {
		return parseFile(filepath.Join(".fnx", "default.json"), repoOverride)
	}
	return nil, errors.New("no fnx config found (fnx.json or .fnx/default.json)")
}

func parseFile(path, repoOverride string) (*ProjectConfig, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// allow $schema present in JSON without failing
	var raw map[string]any
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}
	delete(raw, "$schema")
	out, _ := json.Marshal(raw)

	var cfg ProjectConfig
	if err := json.Unmarshal(out, &cfg); err != nil {
		return nil, err
	}
	if repoOverride != "" {
		cfg.Repo = repoOverride
	}
	return &cfg, nil
}
