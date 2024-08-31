package gsc

func Map(f func(item string) Component, items ...string) Component {
	output := ""
	for _, item := range items {
		output = output + string(f(item))
	}
	return NewComponent(output)
}

// Repeat generates a Component by repeating a given Component a specified number of times.
func Repeat(c Component, count int) Component {
	output := ""
	for i := 0; i < count; i++ {
		output = output + string(c) + "\n"
	}
	return NewComponent(output)
}

// Join concatenates multiple Components into a single Component with an optional separator.
func Join(separator string, components ...Component) Component {
	output := ""
	for i, component := range components {
		output = output + string(component)
		if i < len(components)-1 {
			output = output + separator
		}
	}
	return NewComponent(output)
}

// If generates a Component based on a condition.
func If(condition bool, trueComponent Component) Component {
	if condition {
		return trueComponent
	}
	return NewComponent("")
}

// If generates a Component based on a condition with an else option.
func IfElse(condition bool, trueComponent, falseComponent Component) Component {
	if condition {
		return trueComponent
	}
	return falseComponent
}

// Each applies a function to each Component in a slice and concatenates the results.
func Each(f func(item Component) Component, components ...Component) Component {
	output := ""
	for _, component := range components {
		output = output + string(f(component)) + "\n"
	}
	return NewComponent(output)
}

// Filter returns a Component containing only the Components that satisfy a predicate function.
func Filter(predicate func(item Component) bool, components ...Component) Component {
	output := ""
	for _, component := range components {
		if predicate(component) {
			output = output + string(component) + "\n"
		}
	}
	return NewComponent(output)
}

// WithAttr creates a Component with an additional attribute added.
func WithAttr(c Component, attrName, attrValue string) Component {
	return c.Attr(attrName, attrValue)
}

// WithClass creates a Component with an additional class added.
func WithClass(c Component, className string) Component {
	return c.Class(className)
}

// ConditionalClass adds a class to a Component if a condition is true.
func ConditionalClass(c Component, condition bool, className string) Component {
	if condition {
		return c.Class(className)
	}
	return c
}

// Wrap wraps a Component with another Component.
func Wrap(wrapper, content Component) Component {
	return wrapper.In(content)
}

// Merge merges the content of multiple Components into a single Component.
func Merge(components ...Component) Component {
	output := ""
	for _, component := range components {
		output = output + string(component) + "\n"
	}
	return NewComponent(output)
}
