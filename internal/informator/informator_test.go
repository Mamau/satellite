package informator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEntity struct {
	SomeString string   `yaml:"some_string"`
	SomeBool   bool     `yaml:"some_bool"`
	SomeInt    int64    `yaml:"some_int"`
	SomeSlice  []string `yaml:"some_slice"`
}

func TestScanEntity(t *testing.T) {
	en := testEntity{
		SomeString: "test",
		SomeBool:   true,
		SomeInt:    11,
		SomeSlice: []string{
			"8.8.8.8",
			"0.0.0.0",
		},
	}
	informator := NewInformator(en)
	assert.NotEmpty(t, informator.EntityName)
	assert.Equal(t, informator.EntityName, "testEntity")

	assert.NotEmpty(t, informator.Strings["some_string"])
	assert.Equal(t, informator.Strings["some_string"], "test")

	assert.NotEmpty(t, informator.Booleans["some_bool"])
	assert.True(t, informator.Booleans["some_bool"])

	assert.NotEmpty(t, informator.Integers["some_int"])
	assert.Equal(t, informator.Integers["some_int"], int64(11))

	assert.NotEmpty(t, informator.Slices["some_slice"])
	assert.Equal(t, informator.Slices["some_slice"][0], "8.8.8.8")
	assert.Equal(t, informator.Slices["some_slice"][1], "0.0.0.0")

	en = testEntity{}
	informator = NewInformator(en)
	assert.NotEmpty(t, informator.EntityName)
	assert.Equal(t, informator.EntityName, "testEntity")

	assert.Empty(t, informator.Strings["some_string"])
	assert.Empty(t, informator.Integers["some_int"])
	assert.Empty(t, informator.Slices["some_slice"])
	assert.Empty(t, informator.Booleans["some_bool"])
}
