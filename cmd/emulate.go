package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Emulation struct {
	binaryPath string
	startAddr  uint
	numInst    uint
	endAddr    uint
	logSteps   bool
}

var emulationStruct = Emulation{}

var emulateCmd = &cobra.Command{
	Use:   "emulate [flags] binary",
	Short: "Emulate binary executable files",
	Args: func(cmd *cobra.Command, args []string) error {

		if emulationStruct.startAddr == 0 {
			return errors.New("a start address is needed")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Root().Help()
			os.Exit(1)
		}

		target := args[0]
		// TODO: Create a log library
		fmt.Printf("Opening %s...\n", target)

	},
}

func init() {
	emulateCmd.Flags().StringVar(&emulationStruct.binaryPath, "binary", "", "Binary path to be analyzed and emulated")
	emulateCmd.Flags().UintVar(&emulationStruct.startAddr, "start-at", 0, "Start address of the emulation")
	emulateCmd.Flags().UintVar(&emulationStruct.endAddr, "until", 0, "Emulate until this address")
	emulateCmd.Flags().UintVar(&emulationStruct.numInst, "num-instructions", 0, "Number of instructions to emulate")
	emulateCmd.Flags().BoolVar(&emulationStruct.logSteps, "log", false, "Log each step emulated")
}
