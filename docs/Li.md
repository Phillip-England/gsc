```go
func Li(attr string, content ...string) string {
	return fmt.Sprintf(`<li %s>%s</li>`, attr, strings.Join(content, "\n"))
}
```