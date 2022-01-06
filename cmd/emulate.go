package cmd

import "github.com/spf13/cobra"

type Emulation struct {
	binaryPath string
	startAddr  uint
	endAddr    uint
}

var emulationStruct = Emulation{}

var emulateCmd = &cobra.Command{
	Use:   "emulate [flags] binary",
	Short: "Emulate binary executable files",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	emulateCmd.Flags().StringVarP(&emulationStruct.binaryPath, "binary", "", "", "Binary path to be analyzed and emulated")
	emulateCmd.Flags().UintVarP(&emulationStruct.startAddr, "start-at", "s", 0, "Start address of the emulation")
}
