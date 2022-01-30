package neoengine

import "github.com/0xc0ffeec0de/bino/pkg/r2pipe"

type Binary struct {
	r2            *r2pipe.Pipe
	Arch          string
	path          string
	imports       map[uint]Import
	retAddr       uint64
	StackFrame    []uint8
	StackFrameStr string
	LocalCalls    uint
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
	UntilCall    string
	// hooks []CustomHooks
	IgnoreExtCalls  bool
	MonitorRegister []Register
	ReadRegister    Register

	// "private" information for execution
	hasKnownEnd bool
}

// Instruction struct that wrapp the current opcode information
type Instruction struct {
	Offset   int64  `json:"offset"`
	Esil     string `json:"esil"`
	Refptr   bool   `json:"refptr"`
	FcnAddr  int64  `json:"fcn_addr"`
	FcnLast  int64  `json:"fcn_last"`
	Size     int    `json:"size"`
	Opcode   string `json:"opcode"`
	Disasm   string `json:"disasm"`
	Bytes    string `json:"bytes"`
	Family   string `json:"family"`
	Type     string `json:"type"`
	Reloc    bool   `json:"reloc"`
	TypeNum  int    `json:"type_num"`
	Type2Num int    `json:"type2_num"`
	Jump     uint   `json:"jump"`
	Fail     uint   `json:"fail"`
	Refs     []struct {
		Addr int64  `json:"addr"`
		Type string `json:"type"`
	} `json:"refs"`
}

// Import function struct
type Import struct {
	Ordinal int    `json:"ordinal"`
	Bind    string `json:"bind"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Plt     uint   `json:"plt"`
}

type Stack struct {
	StartOffset uint
	EndOffset   uint
}

// Change to support more arch
type CPU struct {
	Bin           *Binary
	RegisterState x8664Registers
}

type x8664Registers struct {
	RAX uint64 `json:"rax"`
	RBX uint64 `json:"rbx"`
	RDI uint64 `json:"rdi"`
	RCX uint64 `json:"rcx"`
	RSI uint64 `json:"rsi"`
	RIP uint64 `json:"rip"`
	RBP uint64 `json:"rbp"`
	RSP uint64 `json:"rsp"`
	R8  uint64 `json:"r8"`
	R9  uint64 `json:"r9"`
	R10 uint64 `json:"r10"`
	R11 uint64 `json:"r11"`
	R12 uint64 `json:"r12"`
	R13 uint64 `json:"r13"`
	R14 uint64 `json:"r14"`
	R15 uint64 `json:"r15"`
}

type RegRef struct {
	Role   string `json:"role"`
	Reg    string `json:"reg"`
	Value  string `json:"value"`
	RefStr string `json:"refstr"`
}

//

type FinishEmuReason int

const (
	HitCall FinishEmuReason = iota
	ReachEnd
	HitTarget
	Continue
)
