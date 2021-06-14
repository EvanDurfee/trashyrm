package parser

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

type Opts struct {
	Help        bool
	Version     bool
	Verbose     bool
	Interactive InteractiveMode
	Force       bool
	Recurse     bool
	HandleDirs  bool
	Recycle     RecycleMode
	Shred       bool
	DryRun      bool
	Preserve    PreserveMode
}

func NewOpts() *Opts {
	return &Opts{
		Help:        false,
		Version:     false,
		Verbose:     false,
		Interactive: AskOnce,
		Force:       false,
		Recurse:     false,
		HandleDirs:  false,
		Recycle:     RecycleWhitelist,
		Shred:       false,
		DryRun:      false,
		Preserve:    PreserveRoot,
	}
}

type InteractiveMode int

const (
	AskNever InteractiveMode = iota + 1
	AskOnce
	AskAlways
)

func (m *InteractiveMode) Set(value string, opt getopt.Option) error {
	switch opt.Name() {
	case "-i":
		*m = AskAlways
	case "-I":
		*m = AskOnce
	case "--interactive":
		switch value {
		case "never", "no", "none":
			*m = AskNever
		case "once":
			*m = AskOnce
		case "always", "yes", "":
			*m = AskAlways
		default:
			return fmt.Errorf("invalid interactive mode %s", value)
		}
	default:
		return fmt.Errorf("invalid argument %s", opt.Name())
	}
	return nil
}

func (m *InteractiveMode) String() string {
	switch *m {
	case AskNever:
		return "Never"
	case AskOnce:
		return "Once"
	case AskAlways:
		return "Always"
	default:
		return "UNDEFINED"
	}
}

type RecycleMode int

const (
	RecycleNever RecycleMode = iota + 1
	RecycleWhitelist
	RecycleAlways
)

func (m *RecycleMode) Set(value string, opt getopt.Option) error {
	switch opt.Name() {
	case "-c":
		*m = RecycleAlways
	case "--direct", "--unlink", "-u":
		*m = RecycleNever
	case "--recycle":
		switch value {
		case "never", "no":
			*m = RecycleNever
		case "whitelist", "trashpath":
			*m = RecycleWhitelist
		case "always", "yes", "":
			*m = RecycleAlways
		default:
			return fmt.Errorf("invalid recycle mode %s", value)
		}
	default:
		return fmt.Errorf("invalid argument %s", opt.Name())
	}
	return nil
}

func (m *RecycleMode) String() string {
	switch *m {
	case RecycleNever:
		return "Never"
	case RecycleWhitelist:
		return "Whitelist"
	case RecycleAlways:
		return "Always"
	default:
		return "UNDEFINED"
	}
}

type PreserveMode int

const (
	PreserveNone    PreserveMode = 0
	PreserveRoot                 = PreserveNone | 1
	PreserveOtherFs              = PreserveRoot << 1
	PreserveAll                  = PreserveRoot | PreserveOtherFs
)

func (m *PreserveMode) Set(value string, opt getopt.Option) error {
	switch opt.Name() {
	case "--no-preserve-root":
		*m &= PreserveAll ^ PreserveRoot
	case "--one-file-system":
		*m |= PreserveOtherFs
	case "--preserve-root":
		switch value {
		case "":
			*m |= PreserveRoot
		case "all":
			*m = PreserveAll
		default:
			return fmt.Errorf("invalid preserve mode %s", value)
		}
	default:
		return fmt.Errorf("invalid argument %s", opt.Name())
	}
	return nil
}

func (m *PreserveMode) String() string {
	switch *m {
	case PreserveNone:
		return "None"
	case PreserveRoot:
		return "Root"
	case PreserveOtherFs:
		return "Other FS"
	case PreserveAll:
		return "All"
	default:
		return "UNDEFINED"
	}
}
