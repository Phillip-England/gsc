```go
func Div(attr string, content ...string) string {
	return fmt.Sprintf( /*html*/ `<div %s>%s</div>`, attr, strings.Join(content, "\n"))
}
```