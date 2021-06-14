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

// Parse parses the given
func Parse(args []string) (Opts, []string) {
	args = ReorderOpts(args)
	opts := *NewOpts()
	optParser := getopt.New()
	optParser.FlagLong(&opts.Help, "help", 'h', "Print this message and exit").SetFlag()
	optParser.FlagLong(&opts.Version, "version", 'v', "Print version information and exit").SetFlag()
	optParser.FlagLong(&opts.Verbose, "verbose", 0, "Increase verbosity").SetFlag()
	optParser.FlagLong(&opts.Force, "force", 'f', "Ignore nonexistent files and arguments, never prompt").SetFlag()
	optParser.Flag(&opts.Interactive, 'i', "Prompt before removal for every file").SetFlag()
	optParser.Flag(&opts.Interactive, 'I', "Prompt once before removing more than CUTOFF files, or when removing recursively. This is the default mode.").SetFlag()
	optParser.FlagLong(&opts.Interactive, "interactive", 0, "Prompt according to WHEN: never, once (-I), or always (-i). Without WHEN, prompt always.", "WHEN").SetOptional()
	optParser.FlagLong(&opts.Recurse, "recursive", 'r', "Remove directories and their contents recursively").SetFlag()
	optParser.FlagLong(&opts.HandleDirs, "dir", 'd', "Remove empty directories").SetFlag()
	optParser.FlagLong(&opts.Recycle, "recycle", 0, "Move to trash / recycle bin according to WHEN: never, whitelist, or always. Without WHEN, recycle always.", "WHEN").SetOptional()
	optParser.Flag(&opts.Recycle, 'c', "Move to trash / recycle bin").SetFlag()
	optParser.FlagLong(&opts.Recycle, "unlink", 'u', "Remove (unlink) the file(s) even if they're in the trash path").SetFlag()
	optParser.FlagLong(&opts.Recycle, "direct", 0, "Remove (unlink) the file(s) even if they're in the trash path").SetFlag()
	optParser.FlagLong(&opts.Shred, "shred", 's', "Shred the files for more secure deletion").SetFlag()
	optParser.FlagLong(&opts.DryRun, "dryrun", 0, "Dry run, does not modify the file system").SetFlag()
	optParser.Parse(args)
	return opts, optParser.Args()
}
