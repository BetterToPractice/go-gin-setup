package migrate

import (
	"fmt"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		db := lib.NewDatabase()
		if err := db.ORM.AutoMigrate(
			&models.User{},
		); err != nil {
			fmt.Println("Error when migrate", err)
		}
	},
}
