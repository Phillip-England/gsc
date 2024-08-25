package gsc

import "fmt"

func Meta(attr string) string {
	return fmt.Sprintf(`<meta %s></meta>`, attr)
}
