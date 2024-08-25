package gsc

import "fmt"

func P(attr string, text string) string {
	return fmt.Sprintf(`<p %s>%s</p>`, attr, text)
}
