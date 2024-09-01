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
