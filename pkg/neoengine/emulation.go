package neoengine

import (
	"strings"
)

func (e *EmulationProfile) Emulate() (CPU, error) {

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
		shouldCont, reason := e.handleExec() // Apply any kind of constraints in this emulation scenario
		if !shouldCont {
			switch reason {
			// Just hit a hunted call, build the stack frame here
			// because the stackframe yet exists
			case HitCall:
				bin.BuildStackFrame()
			}
			break
		}
	}

	cpuCtx := CPU{
		Bin:           bin,
		RegisterState: bin.Getx8664RegState(),
	}

	return cpuCtx, nil
}

func (e *EmulationProfile) handleExec() (bool, FinishEmuReason) {
	// First handle
	if e.hasKnownEnd {
		currentAddr := e.Binary.CurrentAddress()
		currentAddr = strings.Trim(currentAddr, "\n")
		if currentAddr == e.UntilAddress {
			e.Binary.Step()
			return false, HitTarget
		}
	}

	// Get current instruction
	currInst := e.Binary.GetCurrInstruction()

	// Hit invalid code aka end
	if currInst.Type == "invalid" || currInst.Disasm == "invalid" {
		// e.Binary.SetRegister()
		return false, ReachEnd
	}

	//Check if is the target call
	if currInst.Type == "call" {
		extCall, found := e.Binary.imports[currInst.Jump]

		if found { // just ignore imp call
			if extCall.Name == e.UntilCall {
				return false, HitCall
			}
			e.Binary.StepOver()
		} else {
			e.Binary.retAddr = e.Binary.NextInstAddr()
		}

		return true, Continue
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

	// intel specific
	if currInst.Disasm == "ret" {

		e.Binary.SetRegister("rsp", e.Binary.retAddr)
	}

	return true, Continue
}
