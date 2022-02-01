package neoengine

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/olekukonko/tablewriter"
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
		refTokens := strings.Split(regRef, " ")

		// Fix ref values
		if len(refTokens) >= 5 {
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
func (cpu *CPU) PrintState() {

	// First print register references
	var refs []RegRef = cpu.GetRegsRefs()
	tabelReg := tablewriter.NewWriter(os.Stdout)
	fmt.Println("REGISTER STATE")
	tabelReg.SetHeader([]string{"REG", "Value", "Printable"})
	// tabelReg.SetBorder(false)

	values := make([][]string, 3)
	for _, reg := range refs {

		row := []string{
			reg.Reg, fmt.Sprintf("0x%s", reg.Value), "",
		}

		if reg.RefStr != "" {
			row[2] = reg.RefStr
		}
		values = append(values, row)
	}
	tabelReg.AppendBulk(values)
	tabelReg.Render() // Send output

	if len(cpu.Bin.StackFrame) > 0 {
		fmt.Println("STACK FRAME")
		stackTable := tablewriter.NewWriter(os.Stdout)
		stackTable.SetHeader([]string{"Address", "Value", "Printable"})

		stackValues := make([][]string, len(cpu.Bin.StackFrame)/8)
		// stackTable.SetBorder(false)

		addr := cpu.Bin.StackAddress
		joinValue := ""
		for _, v := range cpu.Bin.StackFrame {
			joinValue = "0x"
			ref := ""
			for _, b := range v {
				r := rune(b)
				if unicode.IsLetter(r) {
					ref += fmt.Sprintf("%c", b)
				} else {
					ref += " "
				}
				joinValue += fmt.Sprintf("%x", b)
			}

			row := []string{
				fmt.Sprintf("0x%x", addr), joinValue, ref}

			stackValues = append(stackValues, row)
			addr -= 8
		}
		stackTable.AppendBulk(stackValues)
		stackTable.Render()
	}

}
