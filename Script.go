package gsc

import "fmt"

func Script(text string) string {
	return fmt.Sprintf( /*html*/ `<script>%s</script>`, text)
}
