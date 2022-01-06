package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Bino",
		Short: "Emulate binary executable code easily",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Root().Help()
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(emulateCmd)
}
