package parser

import (
	"github.com/pborman/getopt/v2"
	"strings"
)

// ReorderOpts takes a list of command line arguments and reorders the options
// in front of the keyword arguments, while preserving the overall order of
// keyword arguments and options
func ReorderOpts(opts []string) []string {
	reordered := opts[:1]
	var kwArgs []string
	for _, opt := range opts[1:] {
		if strings.HasPrefix(opt, "-") {
			reordered = append(reordered, opt)
		} else {
			kwArgs = append(kwArgs, opt)
		}
	}
	return append(reordered, kwArgs...)
}

type Opts struct {
	help bool

}

// TODO
func do() {
	optParser := getopt.New()
	optParser.BoolLong("--help", 'h', "Print this message and exit")

}