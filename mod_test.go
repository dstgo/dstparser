package modparser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestParseModInfoLua(t *testing.T) {
	dir := "testdata/workshop"
	entries, err := os.ReadDir(dir)
	assert.Nil(t, err)

	for _, entry := range entries {
		t.Log(entry.Name())
		infopath := filepath.Join(dir, entry.Name(), "modinfo.lua")
		bytes, err := os.ReadFile(infopath)
		assert.Nil(t, err)

		modInfo, err := ParseModInfoWithEnv(fmt.Sprintf("workshop-%s", entry.Name()), "zh", string(bytes))
		t.Log(modInfo.Name, modInfo.Author, modInfo.Version)
	}
}
