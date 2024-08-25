package gsc

import (
	"fmt"
	"strings"
)

func Ul(attr string, items ...string) string {
	return fmt.Sprintf(`<ul %s>%s</ul>`, attr, strings.Join(items, "\n"))
}
