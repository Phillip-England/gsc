```go
func Nav(attr string, content ...string) string {
	return fmt.Sprintf( /*html*/ `<nav %s>%s</nav>`, attr, strings.Join(content, "\n"))
}
```