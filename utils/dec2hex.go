package utils

import "fmt"

func Dec2Hex(TagGroup int, TagElement int) string {
	return fmt.Sprintf("%04X,%04X", TagGroup, TagElement)
}
