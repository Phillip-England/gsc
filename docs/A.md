```go
func A(attr string, content ...string) string {
	return fmt.Sprintf(`<a %s>%s</a>`, attr, strings.Join(content, "\n"))
}
```