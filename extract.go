package gsc

import (
	"fmt"
	"strings"
)

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

func (c Component) Query(selector string, fn func(child Component) Component) Component {
	str := c.ToString()
	child := c.ExtractChild(selector)
	newChild := fn(child)
	str = strings.Replace(str, child.ToString(), newChild.ToString(), 1)
	return NewComponent(str)
}

func (c Component) ExtractChild(selector string) Component {
	if selector == "" {
		return c
	}
	str := c.ToString()
	selectorPrefix := string(selector[0])
	selectorAttribute := ""
	if selectorPrefix == "#" || selectorPrefix == "." {
		selector = selector[1:]
		if selectorPrefix == "#" {
			selectorAttribute = "id=\""
		}
		if selectorPrefix == "." {
			selectorAttribute = "class=\""
		}
	}
	selector = selectorAttribute + selector
	componentStart := 0
	selectorIndex := strings.Index(str, selector)
	for i := selectorIndex; i > 0; i-- {
		char := string(str[i])
		if char == "<" {
			componentStart = i
			break
		}
	}
	foundFirstClose := false
	foundStartOfClosingTag := false
	for i := 0; i < len(str); i++ {
		if i < componentStart {
			continue
		}
		char := string(str[i])
		if char == ">" {
			foundFirstClose = true
			// is this a self closing tag?
			if string(str[i-1]) == "/" {
				// return the self closing tag here!
				fmt.Println(componentStart, i)
				return NewComponent(str[componentStart : i+1])
			}
		}
		if foundFirstClose && char == "<" && string(str[i+1]) == "/" {
			foundStartOfClosingTag = true
		}
		if foundFirstClose && foundStartOfClosingTag && char == ">" {
			return NewComponent(str[componentStart : i+1])
		}
	}
	if componentStart == 0 {
		return c
	}
	return c
}
