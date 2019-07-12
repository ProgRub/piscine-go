package piscine

import "strings"

func Split(str, charset string) []string {
	return SplitWhiteSpaces(strings.Replace(str, charset, " ", -1))
}
