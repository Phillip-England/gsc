# gsc

go simple component, keeping things.. simple ⚡

## Philosophy

1. components are functions (which resolve to strings)
2. tailwind is used for styling
3. components contain and expose their own javascript


## Examples

a simple component
```go
func SimpleComponent(word string) string {
    return fmt.Sprinf(`
        <p>the word is: %s</p>
    `, word)
}
```

usage
```go
fmt.Println(SimpleComponent("i dont wanna taco bout it 🌮"))
```