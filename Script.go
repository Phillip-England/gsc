package gsc

import "fmt"

func Script(attr string, text string) string {
	return fmt.Sprintf( /*html*/ `<script %s>%s</script>`, attr, text)
}
