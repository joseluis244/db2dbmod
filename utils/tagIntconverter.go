package utils

import "strconv"

func TagIntConverter(list map[string]bool, tag string, value string) any {
	if _, ok := list[tag]; ok {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return int64(intValue)
		}
	}
	return value
}
