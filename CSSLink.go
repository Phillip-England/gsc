package gsc

import "fmt"

func CSSLink(path string) string {
	return fmt.Sprintf(`<link rel="stylesheet" href="%s">`, path)
}
