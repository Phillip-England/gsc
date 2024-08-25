```go
func Style(content string) string {
	return fmt.Sprintf( /*html*/ `
		<style>%s</style>
	`, content)
}
```