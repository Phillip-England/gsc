package gsc

import (
	"fmt"
	"strings"
)

type Component string

func NewComponent(s string, args ...string) Component {
	if len(args) != 0 {
		return Component(fmt.Sprintf(strings.TrimSpace(s), args))
	} else {
		return Component(strings.TrimSpace(s))
	}
}

func (c Component) ToString() string {
	return string(c)
}

func (c Component) Text(text string) Component {
	s := string(c)
	rightAngleIndex := strings.Index(s, ">")
	if rightAngleIndex == -1 {
		return c
	}
	if s[rightAngleIndex-1] == '/' {
		return c
	}
	s = s[:rightAngleIndex+1] + text + s[rightAngleIndex+1:]
	return NewComponent(s)
}

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

func (c Component) Of(funcs ...func(component Component) Component) Component {
	for _, fn := range funcs {
		c = fn(c)
	}
	return c
}

func (c Component) GetAttr(name string) string {
	s := c.ToString()
	if !strings.Contains(s, name+"=") {
		return ""
	}
	qTypeIndex := strings.Index(s, name+"=") + len(name) + 1
	qType := string(s[qTypeIndex])
	attr := ""
	for i := qTypeIndex + 1; i < len(s); i++ {
		char := string(s[i])
		if char == qType && string(s[i-1]) != `\` {
			break
		}
		attr = attr + char
	}
	return attr
}

func (c Component) TagName() string {
	s := c.ToString()
	name := ""
	for i := 1; i < len(s); i++ {
		char := string(s[i])
		if char == " " || char == ">" || char == "/" {
			if i > 1 {
				break
			}
			continue
		}
		name = name + char
	}
	return name
}
