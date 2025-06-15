package secrets

import "github.com/finchknox/fnx/internal/config"

// Pull is stubbed for now: returns an empty map so fnx run compiles.
func Pull(_ *config.ProjectConfig, _ string) (map[string]string, error) {
	return map[string]string{}, nil
}
