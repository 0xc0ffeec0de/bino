package neoengine

import (
	"strings"
)

func (e *EmulationProfile) Emulate() (CPU, error) {

	if e.UntilAddress != "" {
		e.hasKnownEnd = true
	}

	var bin *Binary = e.Binary
	// Set up ESIL
	bin.SeekTo(e.StartAddress)
	bin.SetUpEsil()

	for {
		bin.Step()
		shouldCont, _ := e.handleExec() // Apply any kind of constraints in this emulation scenario
		if !shouldCont {
			bin.BuildStackFrame()
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

	//go First handle
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
			return true, Continue
		} else {
			e.Binary.retAddr = e.Binary.NextInstAddr()
			e.Binary.LocalCalls++
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
		if e.Binary.LocalCalls == 0 {
			return false, ReachEnd
		} else {
			e.Binary.LocalCalls--
		}
	}

	return true, Continue
}
