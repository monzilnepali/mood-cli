package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mood-cli",
	Short: "Listen some relaxing sound in the terminal",
	Long:  `relaxing ambiances from terminal`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("hello")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
