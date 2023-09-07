package cmd

import (
	"errors"
	"github.com/BetterToPractice/go-gin-setup/cmd/makemigrations"
	"github.com/BetterToPractice/go-gin-setup/cmd/migrate"
	"github.com/BetterToPractice/go-gin-setup/cmd/runserver"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(runserver.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(makemigrations.StartCmd)
}

var rootCmd = &cobra.Command{
	Use: "go-gin-setup",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
