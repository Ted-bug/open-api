package cmd

import (
	"fmt"
	"os"

	"github.com/Ted-bug/open-api/internal/model"
	"github.com/Ted-bug/open-api/internal/pkg/db"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database.",
	Long:  "Migrate database.",
	Run:   migrateCmdExcutefunc,
}

func migrateCmdExcutefunc(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("./config/config.yaml"); err != nil && os.IsNotExist(err) {
		fmt.Println("there is not a config/config.yaml in this dir")
		return
	} else {
		if err := db.InitMysql(); err != nil {
			fmt.Println(err)
			return
		}
		db.DB.AutoMigrate(&model.AkSk{}, &model.ShortUrl{}, &model.UniqueNum{})
		fmt.Println("done!")
	}
}
