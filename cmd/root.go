package cmd

import "github.com/spf13/cobra"

func Init() {
	rootCmd := cobra.Command{}
	rootCmd.AddCommand(StartCommand())
}
