package formatters

import (
	"strconv"
	"unicode"
)

func CapitalizeFirstLetter(str string) string {
	if len(str) == 0 {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func StringToUint(str string) uint {
	if str == "" {
		return 0
	}
	val, _ := strconv.ParseUint(str, 10, 64)
	return uint(val)
}

func UintToString(val uint) string {
	return strconv.FormatUint(uint64(val), 10)
}
