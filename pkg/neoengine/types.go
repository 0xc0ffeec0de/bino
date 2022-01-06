package neoengine

import "github.com/0xc0ffeec0de/bino/pkg/r2pipe"

type Binary struct {
	r2   *r2pipe.Pipe
	path string
}

type Register struct {
	RegName string
	Arch    int
	Mode    int
}

type EmulationProfile struct {
	Binary       *Binary
	StartAddress string // string type because it's more easy to deal inside r2 shell
	UntilAddress string
	NumSteps     uint
	// hooks []CustomHooks
	IgnoreExtCalls  bool
	MonitorRegister []Register
	ReadRegister    Register
}

// Change to support more arch
type Context struct {
	RegisterState x8664Registers
}

type x8664Registers struct {
	RAX uint64 `json:"rax"`
	RBX uint64 `json:"rbx"`
	RDI uint64 `json:"rdi"`
}
