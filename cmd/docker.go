package cmd

import (
	"github.com/spf13/cobra"
	"scripts/helpers"
)

// nosCmd represents the nos command
var nosCmd = &cobra.Command{
	Use:   "docker",
	Short: "A brief description of your command",
	Long:  `Docker stuff`,
	Run: func(cmd *cobra.Command, args []string) {

		helpers.Docker()
	},
}

func init() {
	rootCmd.AddCommand(nosCmd)
}
