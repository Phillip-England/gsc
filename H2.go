package gsc

import "fmt"

func H2(attr string, text string) string {
	return fmt.Sprintf(`<h2 %s>%s</h2>`, attr, text)
}
