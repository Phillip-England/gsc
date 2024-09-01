package gsc

import "strings"

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
