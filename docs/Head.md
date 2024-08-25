```go
func Head(attr string, title string, content ...string) string {
	return fmt.Sprintf( /*html*/ `<head %s>%s<title>%s</title></head>`, attr, strings.Join(content, "\n"), title)
}
```