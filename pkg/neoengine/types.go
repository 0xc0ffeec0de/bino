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

	// "private" information for execution
	hasKnownEnd bool
}

// Change to support more arch
type Context struct {
	RegisterState x8664Registers
}

type x8664Registers struct {
	RAX uint64 `json:"rax"`
	RBX uint64 `json:"rbx"`
	RDI uint64 `json:"rdi"`
	RCX uint64 `json:"rcx"`
	RSI uint64 `json:"rsi"`
	RIP uint64 `json:"rip"`
	R8  uint64 `json:"r8"`
	R9  uint64 `json:"r9"`
	R10 uint64 `json:"r10"`
	R11 uint64 `json:"r11"`
	R12 uint64 `json:"r12"`
	R13 uint64 `json:"r13"`
	R14 uint64 `json:"r14"`
	R15 uint64 `json:"r15"`
}
