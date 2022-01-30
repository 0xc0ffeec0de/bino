package neoengine

import (
	"fmt"
	"strings"
)

func (cpu *CPU) GetRegsRefs() (refs []RegRef) {
	// Parse the register struct using reflection
	var bin *Binary = cpu.Bin
	refs = []RegRef{}

	bin.r2.CmdjStruct("arrj", &refs)

	for i := 0; i < len(refs); i++ {
		regName := refs[i].Reg
		regRef := refs[i].RefStr

		// Parse reg ref
		// 1540039 rdi R W 0x6e696874656d6f73
		/*
			0 - value
			1 - regName
			2-5 - permissions
			6 - refValue
			7 - stringValue

		*/
		refTokens := strings.Split(regRef, " ")

		// Fix ref values
		if len(refTokens) >= 5 && refTokens[1] == regName {
			refString, _ := bin.r2.Cmdf("ps @ %s", regName)
			if refString != "" {
				refs[i].RefStr = fmt.Sprintf("\"%s\"", strings.ReplaceAll(refString, "\n", "\\n"))
				continue
			}
		}

		refs[i].RefStr = ""
	}

	return refs
}

// Beautify this
func (cpu CPU) String() string {

	// First print register references
	var refs []RegRef = cpu.GetRegsRefs()

	regState := ""
	out := ""
	for _, reg := range refs {
		regState += fmt.Sprintf("%s = 0x%s", reg.Reg, reg.Value)
		if reg.RefStr != "" {
			regState += fmt.Sprintf(" => %s", reg.RefStr)
		}
		regState += "\n"

	}
	out += regState

	if len(cpu.Bin.StackFrame) > 0 {
		out += cpu.Bin.StackFrameStr
	}

	return out
}
