package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "qp_laravel",
		Short: "QP Laravel CLI tool",
	}

	// Add commands from other files
	rootCmd.AddCommand(createLaravelCmd())

	rootCmd.Execute()
}
