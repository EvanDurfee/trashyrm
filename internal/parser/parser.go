package parser

import (
	"github.com/pborman/getopt/v2"
	"strings"
)

// ReorderOpts takes a list of command line arguments and reorders the options
// in front of the keyword arguments, while preserving the overall order of
// keyword arguments and options. Stops reordering if "--" is encountered, but
// treats "-" like a normal argument
func ReorderOpts(opts []string) []string {
	reordered := opts[:1]
	endOpts := false
	var kwArgs []string
	for _, opt := range opts[1:] {
		if !endOpts && strings.HasPrefix(opt, "-") {
			if "--" == opt {
				endOpts = true
			}
			reordered = append(reordered, opt)
		} else {
			kwArgs = append(kwArgs, opt)
		}
	}
	return append(reordered, kwArgs...)
}

type Parser interface {
	Parse(args []string) (Opts, []string, error)
	Usage() string
}

type parserInternal struct {
	opts      *Opts
	optParser *getopt.Set
}

// Parse parses the given
func (p *parserInternal) Parse(args []string) (Opts, []string, error) {
	args = ReorderOpts(args)
	err := p.optParser.Getopt(args, nil)
	return *p.opts, p.optParser.Args(), err
}

func (p *parserInternal) Usage() string {
	b := strings.Builder{}
	getopt.PrintUsage(&b)
	return b.String()
}

func NewParser() Parser {
	opts := NewOpts()
	optParser := getopt.New()
	optParser.SetProgram("trashyrm")
	optParser.FlagLong(&opts.Help, "help", 'h', "Print this message and exit").SetFlag()
	optParser.FlagLong(&opts.Version, "version", 'v', "Print version information and exit").SetFlag()
	optParser.FlagLong(&opts.Verbose, "verbose", 0, "Increase verbosity").SetFlag()
	optParser.Flag(&opts.Interactive, 'i', "Prompt before removal for every file").SetFlag()
	optParser.Flag(&opts.Interactive, 'I', "Prompt once before removing more than CUTOFF files, or when removing recursively. This is the default mode.").SetFlag()
	optParser.FlagLong(&opts.Interactive, "interactive", 0, "Prompt according to WHEN: never, once (-I), or always (-i). Without WHEN, prompt always.", "WHEN").SetOptional()
	optParser.FlagLong(&opts.Force, "force", 'f', "Ignore nonexistent files and arguments, never prompt").SetFlag()
	optParser.FlagLong(&opts.Recurse, "recursive", 'r', "Remove directories and their contents recursively").SetFlag()
	optParser.FlagLong(&opts.HandleDirs, "dir", 'd', "Remove empty directories").SetFlag()
	optParser.FlagLong(&opts.Recycle, "recycle", 0, "Move to trash / recycle bin according to WHEN: never, whitelist, or always. Without WHEN, recycle always.", "WHEN").SetOptional()
	optParser.Flag(&opts.Recycle, 'c', "Move to trash / recycle bin").SetFlag()
	optParser.FlagLong(&opts.Recycle, "unlink", 'u', "Remove (unlink) the file(s) even if they're in the trash path").SetFlag()
	optParser.FlagLong(&opts.Recycle, "direct", 0, "Remove (unlink) the file(s) even if they're in the trash path").SetFlag()
	optParser.FlagLong(&opts.Shred, "shred", 's', "Shred the files for more secure deletion").SetFlag()
	optParser.FlagLong(&opts.DryRun, "dryrun", 0, "Dry run, does not modify the file system").SetFlag()
	optParser.FlagLong(&opts.Preserve, "preserve-root", 0, "Do not remove '/' (default); with 'all', reject any command line argument on a separate device from its parent", "all").SetOptional()
	optParser.FlagLong(&opts.Preserve, "no-preserve-root", 0, "Do not treat '/' specially").SetFlag()
	optParser.FlagLong(&opts.Preserve, "one-file-system", 0, "When removing a hierarchy recursively, skip any directory that is on a file system different from that of the corresponding command line argument").SetFlag()
	return &parserInternal{
		opts:      opts,
		optParser: optParser,
	}
}
