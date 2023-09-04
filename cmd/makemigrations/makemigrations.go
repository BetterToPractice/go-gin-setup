package makemigrations

import (
	"fmt"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/spf13/cobra"
)

var configFile string
var fileName string

func init() {
	pf := StartCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c", "config/config.yaml", "parameter used to start service")
	pf.StringVarP(&fileName, "fileName", "f", "", "name of migration file")

	cobra.MarkFlagRequired(pf, "fileName")
}

var StartCmd = &cobra.Command{
	Use:          "makemigrations",
	Short:        "Create migration file",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		lib.SetConfigPath(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.NewConfig()
		migration := lib.NewMigration(config)

		if err := migration.Create(fileName); err != nil {
			fmt.Printf("Error to create migration file: %v", err)
		}
	},
}
