package gsc

import "fmt"

func H1(attr string, text string) string {
	return fmt.Sprintf(`<h1 %s>%s</h1>`, attr, text)
}
