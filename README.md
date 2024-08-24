# gsc

go simple component, keeping things.. simple ⚡

## Philosophy

1. components are functions (which return strings)
2. tailwind is used for styling
3. components contain and expose their own javascript


## Examples

a simple component
```go
func _(content string) string {
	return fmt.Sprintf( /*html*/ `<p>%s</p>`, content)
}
```

usage
```go
fmt.Println(SimpleComponent("i dont wanna taco bout it 🌮"))
```