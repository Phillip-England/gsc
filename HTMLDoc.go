package gsc

import (
	"fmt"
	"strings"
)

func HTMLDoc(content ...string) string {
	return fmt.Sprintf( /*html*/ `
		<!DOCTYPE><html>%s</html>
	`, strings.Join(content, "\n"))
}
