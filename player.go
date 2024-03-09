package modparser

import (
	"strings"
	"unsafe"
)

// ParsePlayerTxt parse player.txt file, returns slice of klei ID
func ParsePlayerTxt(content string) []string {
	if len(content) == 0 {
		return nil
	}
	return strings.Split(content, "\n")
}

// ToPlayerTxt converts kleiIDs to player.txt
func ToPlayerTxt(kleiIDs []string) ([]byte, error) {
	joinStr := strings.Join(kleiIDs, "\n")
	return unsafe.Slice(unsafe.StringData(joinStr), len(joinStr)), nil
}
