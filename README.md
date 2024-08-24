# gsc

go simple component, keeping things.. simple ⚡

## Philosophy

1. components are functions (which return strings)
2. tailwind is used for styling
3. components contain and expose their own javascript


## Examples

a simple component
```go
func P(content string) string {
	return fmt.Sprintf( /*html*/ `<p>%s</p>`, content)
}
```

usage
```go
package main

import "fmt"

func main() {
    p := P("i dont wanna taco bout it 🌮")
    fmt.Println(p)
}
```