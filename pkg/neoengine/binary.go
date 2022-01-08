package neoengine

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/0xc0ffeec0de/bino/pkg/r2pipe"
)

func NewBinary() *Binary {
	return &Binary{}
}

func (n *Binary) Open(binaryPath string) error {
	r2, err := r2pipe.NewNativePipe(binaryPath)

	if err != nil {
		return err
	}
	n.r2 = r2
	n.path = binaryPath
	// Quiet mode

	// Set up the in memory cache and code analysis
	n.r2.Cmd("e io.cache=true; e bin.cache=true; aaaa 2> /dev/null")

	return nil
}

func (n *Binary) ReadStrAt(address uint) (string, error) {
	cmd := fmt.Sprintf("ps @ %d", address)
	str, err := n.r2.Cmd(cmd)

	return str, err
}

func (n *Binary) Getx8664RegState() x8664Registers {
	// Get current register state
	regs, err := n.r2.Cmd("aerj")
	if err != nil {
		log.Fatalln(err)
	}
	regsByteArray := []byte(regs)
	registers := x8664Registers{}
	json.Unmarshal(regsByteArray, &registers)

	return registers
}

func (n *Binary) Step() {
	n.r2.Cmd("aeso;so 1")
}

func (n *Binary) SetUpEsil() {
	n.r2.Cmd("aei;aeim;aeip")
}

func (n *Binary) SeekTo(addr string) {
	seekAddr := fmt.Sprintf("s %s", addr)
	n.r2.Cmd(seekAddr)
}

func (n *Binary) CurrentAddress() string {
	curr, _ := n.r2.Cmd("s")
	return curr
}
