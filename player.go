package modparser

import "strings"

// ParsePlayerTxt parse player.txt file, returns slice of klei ID
func ParsePlayerTxt(content string) []string {
	if len(content) == 0 {
		return nil
	}
	return strings.Split(content, "\n")
}
