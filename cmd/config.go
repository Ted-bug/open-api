package cmd

import (
	_ "embed"
	"errors"
	"fmt"
	"os"

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
	if err := CreateConfigFile(cmd.Flag("name").Value.String()); err != nil {
		fmt.Println("create config.yaml failed: ", err)
	} else {
		fmt.Println("create config.yaml success")
	}
}

//go:embed example-file/config_example.yaml
var configExample string

// 创建配置文件示例
func CreateConfigFile(filename string) error {
	if filename == "" {
		filename = "config"
	}
	path := "./config/" + filename + ".yaml"
	if _, err := os.Stat(path); err == nil || !os.IsNotExist(err) {
		return errors.New("the file is exist: " + path)
	}
	if err := os.WriteFile(path, []byte(configExample), 0644); err != nil {
		return err
	}
	return nil
}
