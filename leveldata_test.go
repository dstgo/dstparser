package modparser

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseLevelDataOverridesCave(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/leveldataoverride.cave.lua")
	assert.Nil(t, err)
	overrides, err := ParseLevelDataOverrides(string(bytes))
	assert.Nil(t, err)

	assert.NotEmpty(t, overrides.Id)
	t.Log(overrides)
}

func TestParseLevelDataOverridesMaster(t *testing.T) {
	bytes, err := os.ReadFile("testdata/cluster/leveldataoverride.master.lua")
	assert.Nil(t, err)
	overrides, err := ParseLevelDataOverrides(string(bytes))
	assert.Nil(t, err)

	assert.NotEmpty(t, overrides.Id)
	t.Log(overrides)
}
