package neoengine

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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

	// code analysis
	n.r2.Cmd("e scr.color=0; e io.cache=true; aaaa 2> /dev/null")

	// Map all imports
	impList := []Import{}
	n.imports = make(map[uint]Import)

	n.r2.CmdjStruct("iij", &impList)

	for _, imp := range impList {
		n.imports[imp.Plt] = imp
	}

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
	n.r2.Cmd("aes;so 1")
}

func (n *Binary) SetUpEsil() {
	n.r2.Cmd("aei;aeim;aeip")
}

func (n *Binary) StepOver() {
	// n.r2.Cmd("aess; s @ rip")
	n.r2.Cmd("so 1; aeip")
}

func (n *Binary) SeekTo(addr string) {
	seekAddr := fmt.Sprintf("s %s", addr)
	n.r2.Cmd(seekAddr)
}

func (n *Binary) CurrentAddress() string {
	curr, _ := n.r2.Cmd("s")
	return curr
}

func (n *Binary) GetCurrInstruction() Instruction {
	inst := Instruction{}
	currInst, _ := n.r2.Cmd("pdj 1 ~{0}")
	json.Unmarshal([]byte(currInst), &inst)

	return inst
}

func (n *Binary) DisasmAt(address uint, numOpcodes uint) Instruction {
	inst := Instruction{}
	n.r2.Cmdjf("pdj %d @ %d ~{0}", &inst, numOpcodes, address)

	return inst
}

func (n *Binary) FlipZeroFlagIfSet() {
	zf, _ := n.r2.Cmd("?vi `aer zf`")
	zflag, _ := strconv.ParseInt(zf, 10, 8)
	zflag = zflag & 0 // yep, make sure to always be zero

	cmd := fmt.Sprintf("aer zf=%d", zflag)
	n.r2.Cmd(cmd)
}

func (n *Binary) NextInstAddr() uint64 {
	nextAddrStr, _ := n.r2.Cmd("so 1; ?vi `s` ;so -1")
	nextAddrStr = strings.Trim(nextAddrStr, "\n")
	nextAddr, _ := strconv.ParseUint(nextAddrStr, 10, 64)

	return nextAddr
}

func (n *Binary) BuildStackFrame() {
	regs := n.Getx8664RegState()

	n.StackFrame = []uint8{}
	stackSize := regs.RSP - regs.RBP
	n.r2.CmdjfStruct("xj %d @ rbp", &n.StackFrame, stackSize)
	n.StackFrameStr, _ = n.r2.Cmdf("x %d @ rbp", stackSize)

}

func (n *Binary) SetRegister(regName string, value uint64) {
	cmd := fmt.Sprintf("aer %s=%d", regName, value)
	n.r2.Cmd(cmd)
}
