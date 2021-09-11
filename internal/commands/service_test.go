package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareArgs(t *testing.T) {
	args := []string{"composer-cmd-name", "composer", "i", "--ignore-platform-reqs"}
	sn, command := prepareArgs(args)
	assert.Equal(t, sn, args[0])
	assert.Equal(t, command, args[1:])

	args = []string{"composer-cmd-name"}
	sn, command = prepareArgs(args)
	assert.Equal(t, sn, args[0])
	assert.Empty(t, command)
}
