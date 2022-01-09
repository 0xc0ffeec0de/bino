package neoengine

import (
	"strings"
)

func (e *EmulationProfile) Emulate() (Context, error) {

	if e.UntilAddress != "" {
		e.hasKnownEnd = true
	}

	var bin *Binary = e.Binary
	bin.SetUpEsil()

	// Set up ESIL
	bin.SeekTo(e.StartAddress)
	bin.SetUpEsil()

	for {
		bin.Step()
		shouldCont := e.handleExec() // Apply any kind of constraints in this emulation scenario
		if !shouldCont {
			break
		}
	}

	ctx := Context{
		RegisterState: bin.Getx8664RegState(),
	}
	return ctx, nil
}

func (e *EmulationProfile) handleExec() bool {
	// First handle
	if e.hasKnownEnd {
		currentAddr := e.Binary.CurrentAddress()
		currentAddr = strings.Trim(currentAddr, "\n")
		if currentAddr == e.UntilAddress {
			e.Binary.Step()
			return false
		}
	}

	// Get current instruction
	currInst := e.Binary.GetCurrInstruction()

	// Hit invalid code aka end
	if currInst.Type == "invalid" || currInst.Disasm == "invalid" {
		return false
	}

	//Check if is the target call
	if currInst.Type == "call" {
		_, found := e.Binary.imports[currInst.Jump]
		if found { // just ignore imp call
			e.Binary.StepOver()
		} else {
			e.Binary.retAddr = e.Binary.NextInstAddr()
		}

		return true
	}

	// Look-ahead to bypass stack checking

	if currInst.Type == "cjmp" {
		// Disasm a single instruction at the jump address
		// to check if is a call to __stack_chk_fail
		// if is, ignore by flipping the ZF bit
		asmAt := e.Binary.DisasmAt(currInst.Jump, 1)
		if asmAt.Type == "call" {

			imp, found := e.Binary.imports[asmAt.Jump]
			if found && imp.Name == "__stack_chk_fail" {
				e.Binary.FlipZeroFlagIfSet()
			}
		}
	}

	// fmt.Printf("0x%x\t%s\n", currInst.Offset, currInst.Disasm)

	if currInst.Disasm == "ret" {
		e.Binary.SetRegister("rsp", e.Binary.retAddr)
	}

	return true
}
