package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Branch  string
	Date    string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "App Version.",
	Long:  "App Version.",
	Run:   versionCmdExcutefunc,
}

func versionCmdExcutefunc(cmd *cobra.Command, args []string) {
	fmt.Println("Version is 0.1.0.")
}
