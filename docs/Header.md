```go
func Header(attr string, content ...string) string {
	return fmt.Sprintf( /*html*/ `<header %s>%s</header>`, attr, strings.Join(content, "\n"))
}
```