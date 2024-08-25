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

### Styles are Contained
gsc encourages sytles to be contains to the respective component. Tailwind makes this easy, but in some cases, like using CSS grid templates, plain CSS might be preferable. In such cases, the use of random class names may be used as an alternative.

```go

```

### Components Expose Hooks
If the outside world needs to interact with a component, it should do so through a **client primitive**. Client primitives are hooks which enable components to communicate and interact with each other.


## Components

a simple component
```go
package main

import "fmt"

func P(content string) string {
	return fmt.Sprintf( `<p>%s</p>`, content)
}

func main() {
    p := P("i dont wanna taco bout it 🌮")
    fmt.Println(p)
}
```
