package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/0xc0ffeec0de/bino/pkg/neoengine"
	"github.com/spf13/cobra"
)

type Emulation struct {
	binaryPath string
	startAddr  string
	endAddr    string

	logSteps  bool
	untilCall string
}

var emulationStruct = Emulation{}
var binary *neoengine.Binary = neoengine.NewBinary()

var emulateCmd = &cobra.Command{
	Use:   "emulate [flags] binary",
	Short: "Emulate binary executable files",
	Args: func(cmd *cobra.Command, args []string) error {
		if emulationStruct.startAddr == "0x0" && len(args) > 0 {
			return errors.New("a start address is needed")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		target := args[0]
		// TODO: Create a log library
		fmt.Printf("Opening %s...\n", target)
		if err := binary.Open(target); err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		emuProfile := neoengine.EmulationProfile{
			Binary:       binary,
			StartAddress: emulationStruct.startAddr,
			UntilAddress: emulationStruct.endAddr,
			UntilCall:    emulationStruct.untilCall,
		}
		// Emulate
		cpuState, err := emuProfile.Emulate()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Reading content of RDI...")
		data, err := binary.ReadStrAt(uint(cpuState.RegisterState.RDI))

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("RDI => %s\n", data)
	},
}

func init() {
	emulateCmd.Flags().StringVar(&emulationStruct.binaryPath, "binary", "", "Binary path to be analyzed and emulated")
	emulateCmd.Flags().StringVar(&emulationStruct.startAddr, "start-at", "0x0", "Start address of the emulation")
	emulateCmd.Flags().StringVar(&emulationStruct.endAddr, "until", "0x0", "Emulate until this address")
	emulateCmd.Flags().StringVar(&emulationStruct.untilCall, "until-call", "", "Emulate until a function call")
	emulateCmd.Flags().BoolVar(&emulationStruct.logSteps, "log", false, "Log each step emulated")
}
