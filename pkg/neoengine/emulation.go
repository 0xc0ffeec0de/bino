package neoengine

type Register struct {
	RegName string
	Arch    int
	Mode    int
}

type EmulationProfile struct {
	NEngine      *Binary
	StartAddress uint
	UntilAddress uint
	// hooks []CustomHooks
	IgnoreExtCalls  bool
	MonitorRegister []Register
}
