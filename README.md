# gsc

go simple component, keeping things.. simple ⚡

## Philosophy

### Components are Functions
gsc encourages the use of functions to compose and output HTML. Functions are simple. Simple good.

echo, echo 👂
```go
package main

import "fmt"

func Echo(phrase string) string {
	return fmt.Sprintf( `<p>%s</p>`, phrase)
}

func main() {
    fmt.Println(Echo("moms spaghetti 🍝"))
}
```