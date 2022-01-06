package neoengine

import (
	"fmt"

	"github.com/0xc0ffeec0de/bino/pkg/r2pipe"
)

type Binary struct {
	r2   *r2pipe.Pipe
	path string
}

func NewBinary() *Binary {
	return &Binary{}
}

func (n *Binary) Open(binaryPath string) error {
	r2, err := r2pipe.NewNativePipe(binaryPath)

	if err != nil {
		return err
	}
	n.r2 = r2

	// Quiet mode

	// Set up the in memory cache and code analysis
	n.r2.Cmd("e io.cache=true; e bin.cache=true; aaaa 2> /dev/null")

	return nil
}

func (n *Binary) Emulate(startAddress uint, numInstructions uint) {
	// reset state machine and start the stack
	n.r2.Cmd("aei;aeim")

	// seek

	var i uint = 0
	for ; i < numInstructions; i++ {
		out, err := n.r2.Cmd("aeso; aer")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(out)
	}
}
