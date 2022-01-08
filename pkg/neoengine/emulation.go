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

	return true
}
