package Repeat

import "strings"

func Repeat(leter string, times int) string {
	var repeated string
	repeated = strings.Repeat(leter, times)
	return repeated
}
