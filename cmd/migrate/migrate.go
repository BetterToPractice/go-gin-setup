package migrate

import (
	"fmt"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/spf13/cobra"
)

var configFile string

func init() {
	pf := StartCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c", "config/config.yaml", "this parameter is used to start the service application")
}

var StartCmd = &cobra.Command{
	Use: "migrate",
	PreRun: func(cmd *cobra.Command, args []string) {
		lib.SetConfigPath(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		db := lib.NewDatabase(lib.NewConfig())
		if err := db.ORM.AutoMigrate(
			&models.User{},
		); err != nil {
			fmt.Println("Error when migrate", err)
		}
	},
}
