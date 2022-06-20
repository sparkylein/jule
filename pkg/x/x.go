package x

import "github.com/the-xlang/xxc/pkg/xset"

// X constants.
const (
	Version       = `@developer_beta 0.0.1`
	SrcExt        = `.xx`
	DocExt        = ".xdoc"
	SettingsFile  = "x.set"
	Stdlib        = "lib"
	Localizations = "localization"

	EntryPoint = "main"

	Anonymous = "<anonymous>"

	DocPrefix = "doc:"

	PlatformWindows = "windows"
	PlatformLinux   = "linux"
	PlatformDarwin  = "darwin"
)

// Environment Variables.
var (
	LangsPath  string
	StdlibPath string
	ExecPath   string
	Set        *xset.XSet
)
