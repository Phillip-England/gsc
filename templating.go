package gsc

import "strings"

func (c Component) In(components ...Component) Component {
	componentStr := ""
	for _, component := range components {
		componentStr = componentStr + string(component)
	}
	s := string(c)
	rightAngleIndex := strings.Index(s, ">")
	if rightAngleIndex == -1 {
		return c
	}
	if s[rightAngleIndex-1] == '/' {
		return c
	}
	s = s[:rightAngleIndex+1] + componentStr + s[rightAngleIndex+1:]
	return NewComponent(s)
}

func Map(f func(item string) Component, items ...string) Component {
	output := ""
	for _, item := range items {
		output = output + string(f(item))
	}
	return NewComponent(output)
}

func If(condition bool, trueComponent Component) Component {
	if condition {
		return trueComponent
	}
	return NewComponent("")
}

func IfElse(condition bool, trueComponent, falseComponent Component) Component {
	if condition {
		return trueComponent
	}
	return falseComponent
}

func (c Component) Of(funcs ...func(component Component) Component) Component {
	for _, fn := range funcs {
		c = fn(c)
	}
	return c
}
