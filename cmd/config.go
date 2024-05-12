package cmd

import (
	"fmt"

	"github.com/Ted-bug/open-api/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create a config file.",
	Long:  "Create a config file.",
	Run:   configCmdExcutefunc,
}

// 初始化配置文件命令
func configCmdExcutefunc(cmd *cobra.Command, args []string) {
	if err := config.CreateConfig(); err != nil {
		fmt.Println("create config.yaml failed: ", err)
	} else {
		fmt.Println("create config.yaml success")
	}
}
