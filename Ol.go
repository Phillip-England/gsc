package gsc

import (
	"fmt"
	"strings"
)

func Ol(attr string, items ...string) string {
	return fmt.Sprintf(`<ol %s>%s</ol>`, attr, strings.Join(items, "\n"))
}
