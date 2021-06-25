package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReorderOpts(t *testing.T) {
	args := []string{"cmd", "-short", "--long", "./file.txt", "/tmp", "--verbose", "hello", "-a", "dash-here", "--", "-b"}
	expected := []string{"cmd", "-short", "--long", "--verbose", "-a", "--", "./file.txt", "/tmp", "hello", "dash-here", "-b"}
	actual := ReorderOpts(args)
	assert.Equal(t, expected, actual)
}

func Parse(args []string) (Opts, []string, error) {
	return NewParser().Parse(args)
}

func TestInteractiveMode_Set(t *testing.T) {
	expectedOpts := *NewOpts()
	var expectedArgs []string

	opts, args, err := Parse([]string{"cmd"})
	expectedOpts.Interactive = AskOnce
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--interactive"})
	expectedOpts.Interactive = AskAlways
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--interactive="})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--interactive=always"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--interactive=never"})
	expectedOpts.Interactive = AskNever
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--interactive=no"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--interactive=none"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-i"})
	expectedOpts.Interactive = AskAlways
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-Ii"})
	expectedOpts.Interactive = AskAlways
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-I"})
	expectedOpts.Interactive = AskOnce
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-iI"})
	expectedOpts.Interactive = AskOnce
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)
}

func TestRecycleMode_Set(t *testing.T) {
	expectedOpts := *NewOpts()
	var expectedArgs []string

	opts, args, err := Parse([]string{"cmd"})
	expectedOpts.Recycle = RecycleWhitelist
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--recycle"})
	expectedOpts.Recycle = RecycleAlways
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--recycle="})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--recycle=always"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--recycle=yes"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-c"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-uc"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--recycle=whitelist"})
	expectedOpts.Recycle = RecycleWhitelist
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--recycle=trashpath"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--direct"})
	expectedOpts.Recycle = RecycleNever
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--unlink"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-u"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "-cu"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)
}

func TestPreserveMode_Set(t *testing.T) {
	expectedOpts := *NewOpts()
	var expectedArgs []string

	opts, args, err := Parse([]string{"cmd"})
	expectedOpts.Preserve = PreserveRoot
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--no-preserve-root"})
	expectedOpts.Preserve = PreserveNone
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--no-preserve-root", "--one-file-system"})
	expectedOpts.Preserve = PreserveOtherFs
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--one-file-system"})
	expectedOpts.Preserve = PreserveAll
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--preserve-root", "--one-file-system"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)

	opts, args, err = Parse([]string{"cmd", "--preserve-root=all"})
	assert.Equal(t, expectedOpts, opts)
	assert.ElementsMatch(t, expectedArgs, args)
	assert.NoError(t, err)
}
