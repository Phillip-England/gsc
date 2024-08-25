package gsc

import "strings"

func Group(components ...string) string {
	return strings.Join(components, "\n")
}
