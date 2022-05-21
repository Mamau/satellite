package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("SOME_VAR", "test")
	result := getEnv("SOME_VAR", "default")
	assert.Equal(t, "test", result)
	result = getEnv("SOME_VAR_2", "default")
	assert.Equal(t, "default", result)
}
