package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "open-api",
	Short: "OpenAPI is a set of tools.",
	Long:  `OpenAPI is a set of tools.`,
	Run:   rootCmdExcutefunc,
}

func init() {
	rootCmd.AddCommand(versionCmd, runCmd)
}

func rootCmdExcutefunc(cmd *cobra.Command, args []string) {
	fmt.Println("Welcom to OpenAPI.")
}

func Excutefunc() error {
	return rootCmd.Execute()
}
