package gsc

import (
	"fmt"
	"strings"
)

//========================================
//
//========================================

type Component string

func NewComponent(s string, args ...string) Component {
	if len(args) != 0 {
		return Component(fmt.Sprintf(strings.TrimSpace(s), args))
	} else {
		return Component(strings.TrimSpace(s))
	}
}

func (c Component) Text(text string) Component {
	s := string(c)
	s = strings.Replace(s, "><", `>`+text+`<`, 1)
	return NewComponent(s)
}

func (c Component) Content(components ...Component) Component {
	componentStr := ""
	for _, component := range components {
		componentStr = componentStr + string(component)
	}
	s := string(c)
	s = strings.Replace(s, "><", `>`+componentStr+`<`, 1)
	return NewComponent(s)
}

func (c Component) InsertAttr(attrName string, value string) Component {
	strComponent := string(c)
	rightBracketIndex := strings.Index(strComponent, ">")
	if rightBracketIndex == -1 {
		return ""
	}
	if string(strComponent[rightBracketIndex-1]) == "/" {
		c = NewComponent(strings.Replace(strComponent, "/>", ` `+attrName+`='`+value+`'/>`, 1))
	} else {
		c = NewComponent(strings.Replace(strComponent, ">", ` `+attrName+`='`+value+`'>`, 1))
	}
	return c
}

func (c Component) Class(className string) Component {
	c = Component(c.InsertAttr("class", className))
	return c
}

func (c Component) Src(s string) Component {
	c = Component(c.InsertAttr("src", s))
	return c
}

func (c Component) Href(s string) Component {
	c = Component(c.InsertAttr("href", s))
	return c
}

func (c Component) Rel(s string) Component {
	c = Component(c.InsertAttr("rel", s))
	return c
}

func (c Component) Name(s string) Component {
	c = Component(c.InsertAttr("name", s))
	return c
}

//========================================
// COMPONENTS
//========================================

func Link() Component {
	return NewComponent(`<link />`)
}

func Meta() Component {
	return NewComponent(`<meta />`)
}

func P() Component {
	return NewComponent(`<p></p>`)
}

func H1() Component {
	return NewComponent(`<h1></h1>`)
}

func H2() Component {
	return NewComponent(`<h2></h2>`)
}

func H3() Component {
	return NewComponent(`<h3></h3>`)
}

func A() Component {
	return NewComponent(`<a></a>`)
}

func Head() Component {
	return NewComponent(`<head></head>`)
}

func HTMLDoc() Component {
	return NewComponent(`<html></html>`)
}

func Div() Component {
	return NewComponent(`<div></div>`)
}

func Header() Component {
	return NewComponent(`<header></header>`)
}

func Li() Component {
	return NewComponent(`<li></li>`)
}

func Main() Component {
	return NewComponent(`<main></main>`)
}

func Nav() Component {
	return NewComponent(`<nav></nav>`)
}

func Ol() Component {
	return NewComponent(`<ol></ol>`)
}

func Script() Component {
	return NewComponent(`<script></script>`)
}

func Ul() Component {
	return NewComponent(`<ul></ul>`)
}

func Article() Component {
	return NewComponent(`<article></article>`)
}

func Body() Component {
	return NewComponent(`<body></body>`)
}

//========================================
// ATTRIBUTES
//========================================

func Attr(attributes ...string) string {
	output := ""
	for _, str := range attributes {
		output = output + str
	}
	return output
}

func Class(className string) string {
	return `class="` + className + `" `
}

func Src(s string) string {
	return `src="` + s + `" `
}

func Rel(className string) string {
	return `rel="` + className + `" `
}

func Href(s string) string {
	return `href="` + s + `" `
}

func Name(s string) string {
	return `name="` + s + `" `
}

//========================================
// TEMPLATING FUNCTIONS
//========================================

func Group(components ...string) string {
	return strings.Join(components, "\n")
}

func Map(f func(item string) string, items ...string) string {
	output := ""
	for _, item := range items {
		output = output + f(item) + "\n"
	}
	return output
}
