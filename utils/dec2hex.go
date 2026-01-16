package utils

import "fmt"

func Dec2Hex(TagGroup int, TagElement int) string {
	return fmt.Sprintf("%04x,%04x", TagGroup, TagElement)
}
