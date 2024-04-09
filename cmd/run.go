package cmd

import (
	"github.com/Ted-bug/open-api/internal"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a web application.",
	Long:  "Run a web application.",
	Run:   runCmdExcutefunc,
}

func runCmdExcutefunc(cmd *cobra.Command, args []string) {
	internal.Run()
}
