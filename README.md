# gsc

go simple components 💥

keeping components in go as simple as possible

## Useful Skeletons

components (don't forget to name your component)
```go
func _(name string) string {
	return fmt.Sprintf( /*html*/ `
		<h1>Hello, %s!</h1>
	`, name)
}
```