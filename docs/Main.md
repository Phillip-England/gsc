```go
func Main(attr string, content ...string) string {
	return fmt.Sprintf( /*html*/ `
		<main %s>
			%s
		</main>
	`, attr, strings.Join(content, "\n"))
}
```