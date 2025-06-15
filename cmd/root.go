package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/finchknox/fnx/internal/config"
)

var (
	cfg     *config.ProjectConfig
	cfgErr  error
	flagEnv string
	flagRepo string
)

var rootCmd = &cobra.Command{
	Use:   "fnx",
	Short: "FinchKnox CLI – zero-friction secrets runner",
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		cfg, cfgErr = config.Load(flagRepo)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&flagEnv,  "env",  "", "Override environment (dev, prod, …)")
	rootCmd.PersistentFlags().StringVar(&flagRepo, "repo", "", "Override secrets repo URL")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
