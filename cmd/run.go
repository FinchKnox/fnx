package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/finchknox/fnx/internal/secrets" // you’ll create this soon
)

var runCmd = &cobra.Command{
	Use:   "run -- <command> [args]",
	Short: "Execute a command with secrets injected (dotenv replacement)",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgErr != nil {
			return cfgErr
		}
		if len(args) == 0 {
			return fmt.Errorf("missing command after --")
		}

		// 1. Pull secrets → map[string]string
		envVals, err := secrets.Pull(cfg, flagEnv)
		if err != nil {
			return err
		}

		// 2. Merge into child environment
		childEnv := os.Environ()
		for k, v := range envVals {
			childEnv = append(childEnv, fmt.Sprintf("%s=%s", k, v))
		}

		// 3. Replace current PID with child
		bin, lookErr := exec.LookPath(args[0])
		if lookErr != nil {
			return lookErr
		}
		return syscall.Exec(bin, args, childEnv)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
