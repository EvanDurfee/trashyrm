package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReorderOpts(t *testing.T) {
	assert := assert.New(t)
	args := []string {"cmd", "-short", "--long", "./file.txt", "/tmp", "--verbose", "hello", "-a", "dash-here", "-b"}
	expected := []string {"cmd", "-short", "--long", "--verbose", "-a", "-b", "./file.txt", "/tmp", "hello", "dash-here"}
	actual := ReorderOpts(args)
	assert.Equal(expected, actual)
}