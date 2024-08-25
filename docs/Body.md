```go
func Body(attr string, content ...string) string {
	return fmt.Sprintf( /*html*/ `<body %s>%s</body>`, attr, strings.Join(content, "\n"))
}
```