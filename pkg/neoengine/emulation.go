package neoengine

import (
	"fmt"

	"github.com/0xc0ffeec0de/bino/pkg/r2pipe"
)

type Register struct {
	RegName string
	Arch    int
	Mode    int
}

type EmulationProfile struct {
	Binary       *Binary
	StartAddress uint
	UntilAddress uint
	NumSteps     uint
	// hooks []CustomHooks
	IgnoreExtCalls  bool
	MonitorRegister []Register
	ReadRegister    Register
}

func (e *EmulationProfile) Emulate() {
	var pipe *r2pipe.Pipe = e.Binary.r2

	// Set up ESIL
	pipe.Cmd("aei;aeim")
	pipe.Cmd(fmt.Sprintf("s %d\n", e.StartAddress))

	// var i uint = 0
	// var steps := e.unt
	// for ; i <
}
