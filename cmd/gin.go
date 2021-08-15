package cmd

import (
	"github.com/raw34/gin-test/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ginCmd)
}

var ginCmd = &cobra.Command{
	Use:   "gin",
	Short: "gin server",
	Long:  `gin api server`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Execute()
	},
}
