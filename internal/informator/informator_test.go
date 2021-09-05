package informator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEntity struct {
	SomeString string
	SomeBool   bool
	SomeSlice  []string
	SomeInt    int64
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

	assert.NotEmpty(t, informator.Strings["SomeString"])
	assert.Equal(t, informator.Strings["SomeString"], "test")

	assert.NotEmpty(t, informator.Booleans["SomeBool"])
	assert.True(t, informator.Booleans["SomeBool"])

	assert.NotEmpty(t, informator.Integers["SomeInt"])
	assert.Equal(t, informator.Integers["SomeInt"], int64(11))

	assert.NotEmpty(t, informator.Slices["SomeSlice"])
	assert.Equal(t, informator.Slices["SomeSlice"][0], "8.8.8.8")
	assert.Equal(t, informator.Slices["SomeSlice"][1], "0.0.0.0")

	en = testEntity{}
	informator = NewInformator(en)
	assert.NotEmpty(t, informator.EntityName)
	assert.Equal(t, informator.EntityName, "testEntity")

	assert.Empty(t, informator.Strings["SomeString"])
	assert.Empty(t, informator.Integers["SomeInt"])
	assert.Empty(t, informator.Slices["SomeSlice"])
	assert.Empty(t, informator.Booleans["SomeBool"])
}
