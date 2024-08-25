```go
func Article(attr string, content ...string) string {
	output := strings.Join(content, "\n")
	return fmt.Sprintf( /*html*/ `
		<article %s>
			%s
		</article>
	`, attr, output)
}
```