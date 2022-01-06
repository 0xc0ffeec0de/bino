package neoengine

import (
	"fmt"
	"strings"

	"github.com/0xc0ffeec0de/bino/pkg/r2pipe"
)

func (e *EmulationProfile) Emulate() (Context, error) {
	var pipe *r2pipe.Pipe = e.Binary.r2

	// Set up ESIL
	pipe.Cmd(fmt.Sprintf("s %s\n", e.StartAddress))
	pipe.Cmd("aei;aeim;aeip")

	for {
		pipe.Cmd("aeso;so 1") // emu and seek 1
		currentAddr, _ := pipe.Cmd("s")
		currentAddr = strings.Trim(currentAddr, "\n")
		if currentAddr == e.UntilAddress {
			pipe.Cmd("aeso") // last emu
			break
		}
	}

	ctx := Context{
		RegisterState: e.Binary.Getx8664RegState(),
	}
	return ctx, nil
}
