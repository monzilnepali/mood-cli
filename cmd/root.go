package cmd

import (
	"fmt"
	"os"

	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mood-cli",
	Short: "relaxing cli",
	Long:  `Listen to relaxing ambiances in your cli`,
	Run: func(cmd *cobra.Command, args []string) {
		emoji.Println("\n Mood Cli :notes:")
		Menu()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
