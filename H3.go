package gsc

import "fmt"

func H3(attr string, text string) string {
	return fmt.Sprintf(`<h3 %s>%s</h3>`, attr, text)
}
