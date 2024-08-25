package gsc

import "fmt"

func Style(content string) string {
	return fmt.Sprintf( /*html*/ `
		<style>%s</style>
	`, content)
}
