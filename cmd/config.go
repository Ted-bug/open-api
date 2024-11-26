package cmd

import (
	_ "embed"
	"fmt"

	"github.com/Ted-bug/open-api/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create a config file.",
	Long:  "Create a config file.",
	Run:   configCmdExcutefunc,
}

func InitConfigCmd() {
	configCmd.PersistentFlags().StringP("name", "n", "config.yaml", "config file name")
}

// 初始化配置文件命令
func configCmdExcutefunc(cmd *cobra.Command, args []string) {
	if err := config.CreateConfigFile(cmd.Flag("name").Value.String()); err != nil {
		fmt.Println("create config.yaml failed: ", err)
	} else {
		fmt.Println("create config.yaml success")
	}
}
