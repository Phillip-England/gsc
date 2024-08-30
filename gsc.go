package gsc

import (
	"fmt"
	"strings"
)

//========================================
// Component Type
//========================================

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
	s = strings.Replace(s, "><", `>`+text+`<`, 1)
	return NewComponent(s)
}

func (c Component) In(components ...Component) Component {
	componentStr := ""
	for _, component := range components {
		componentStr = componentStr + string(component)
	}
	s := string(c)
	s = strings.Replace(s, "><", `>`+componentStr+`<`, 1)
	return NewComponent(s)
}

func (c Component) Of(funcs ...func(component Component) Component) Component {
	for _, fn := range funcs {
		c = fn(c)
	}
	return c
}

func (c Component) Attr(attrName string, value string) Component {
	strComponent := string(c)
	rightBracketIndex := strings.Index(strComponent, ">")
	if rightBracketIndex == -1 {
		return ""
	}
	if string(strComponent[rightBracketIndex-1]) == "/" {
		c = NewComponent(strings.Replace(strComponent, "/>", ` `+attrName+`="`+value+`"/>`, 1))
	} else {
		c = NewComponent(strings.Replace(strComponent, ">", ` `+attrName+`="`+value+`">`, 1))
	}
	return c
}

func (c Component) AddClass(classNames ...string) Component {
	str := c.ToString()
	for _, name := range classNames {
		str = strings.Replace(str, "class=\"", `class="`+name+` `, 1)
	}
	return NewComponent(str)
}

func (c Component) Class(className string) Component {
	return c.Attr("class", className)
}

func (c Component) Src(s string) Component {
	return c.Attr("src", s)
}

func (c Component) Href(s string) Component {
	return c.Attr("href", s)
}

func (c Component) Rel(s string) Component {
	return c.Attr("rel", s)
}

func (c Component) Name(s string) Component {
	return c.Attr("name", s)
}

func (c Component) ID(s string) Component {
	return c.Attr("id", s)
}

func (c Component) Type(s string) Component {
	return c.Attr("type", s)
}

func (c Component) Value(s string) Component {
	return c.Attr("value", s)
}

func (c Component) Placeholder(s string) Component {
	return c.Attr("placeholder", s)
}

func (c Component) Content(s string) Component {
	return c.Attr("content", s)
}

func (c Component) Alt(s string) Component {
	return c.Attr("alt", s)
}

func (c Component) Title(s string) Component {
	return c.Attr("title", s)
}

func (c Component) Fill(s string) Component {
	return c.Attr("fill", s)
}

func (c Component) ViewBox(s string) Component {
	return c.Attr("viewBox", s)
}

func (c Component) D(s string) Component {
	return c.Attr("d", s)
}

func (c Component) StrokeLineCap(s string) Component {
	return c.Attr("stroke-linecap", s)
}

func (c Component) StrokeWidth(s string) Component {
	return c.Attr("stroke-width", s)
}

func (c Component) StrokeLineJoin(s string) Component {
	return c.Attr("stroke-linejoin", s)
}

func (c Component) Stroke(s string) Component {
	return c.Attr("stroke", s)
}

func (c Component) Disabled() Component {
	return c.Attr("disabled", "disabled")
}

func (c Component) ReadOnly() Component {
	return c.Attr("readonly", "readonly")
}

func (c Component) Required() Component {
	return c.Attr("required", "required")
}

func (c Component) Checked() Component {
	return c.Attr("checked", "checked")
}

func (c Component) MaxLength(n int) Component {
	return c.Attr("maxlength", fmt.Sprintf("%d", n))
}

func (c Component) MinLength(n int) Component {
	return c.Attr("minlength", fmt.Sprintf("%d", n))
}

func (c Component) Pattern(s string) Component {
	return c.Attr("pattern", s)
}

func (c Component) Step(s string) Component {
	return c.Attr("step", fmt.Sprintf("%s", s))
}

func (c Component) Width(s string) Component {
	return c.Attr("width", fmt.Sprintf("%s", s))
}

func (c Component) Height(s string) Component {
	return c.Attr("height", fmt.Sprintf("%s", s))
}

func (c Component) Data(key, value string) Component {
	return c.Attr("data-"+key, value)
}

func (c Component) Aria(key, value string) Component {
	return c.Attr("aria-"+key, value)
}

func (c Component) Xmlns(s string) Component {
	return c.Attr("pattern", s)
}

//========================================
// Components
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

func Title() Component {
	return NewComponent(`<title></title>`)
}
func Span() Component {
	return NewComponent(`<span></span>`)
}

func Img() Component {
	return NewComponent(`<img />`)
}

func Button() Component {
	return NewComponent(`<button></button>`)
}

func Input() Component {
	return NewComponent(`<input />`)
}

func Form() Component {
	return NewComponent(`<form></form>`)
}

func Label() Component {
	return NewComponent(`<label></label>`)
}

func Footer() Component {
	return NewComponent(`<footer></footer>`)
}

func Section() Component {
	return NewComponent(`<section></section>`)
}

func Aside() Component {
	return NewComponent(`<aside></aside>`)
}

func Table() Component {
	return NewComponent(`<table></table>`)
}

func Tr() Component {
	return NewComponent(`<tr></tr>`)
}

func Td() Component {
	return NewComponent(`<td></td>`)
}

func Th() Component {
	return NewComponent(`<th></th>`)
}

func Thead() Component {
	return NewComponent(`<thead></thead>`)
}

func Tbody() Component {
	return NewComponent(`<tbody></tbody>`)
}

func Tfoot() Component {
	return NewComponent(`<tfoot></tfoot>`)
}

func Select() Component {
	return NewComponent(`<select></select>`)
}

func Option() Component {
	return NewComponent(`<option></option>`)
}

func Textarea() Component {
	return NewComponent(`<textarea></textarea>`)
}

func Strong() Component {
	return NewComponent(`<strong></strong>`)
}

func Em() Component {
	return NewComponent(`<em></em>`)
}

func B() Component {
	return NewComponent(`<b></b>`)
}

func I() Component {
	return NewComponent(`<i></i>`)
}

func Hr() Component {
	return NewComponent(`<hr />`)
}

func Br() Component {
	return NewComponent(`<br />`)
}

func Blockquote() Component {
	return NewComponent(`<blockquote></blockquote>`)
}

func Code() Component {
	return NewComponent(`<code></code>`)
}

func Pre() Component {
	return NewComponent(`<pre></pre>`)
}

func Figure() Component {
	return NewComponent(`<figure></figure>`)
}

func Figcaption() Component {
	return NewComponent(`<figcaption></figcaption>`)
}

func Audio() Component {
	return NewComponent(`<audio></audio>`)
}

func Video() Component {
	return NewComponent(`<video></video>`)
}

func Source() Component {
	return NewComponent(`<source />`)
}

func Canvas() Component {
	return NewComponent(`<canvas></canvas>`)
}

func Embed() Component {
	return NewComponent(`<embed />`)
}

func Object() Component {
	return NewComponent(`<object></object>`)
}

func Param() Component {
	return NewComponent(`<param />`)
}

func Svg() Component {
	return NewComponent(`<svg></svg>`)
}

func Path() Component {
	return NewComponent(`<path></path>`)
}

func Polygon() Component {
	return NewComponent(`<polygon></polygon>`)
}

func Circle() Component {
	return NewComponent(`<circle></circle>`)
}

func Rect() Component {
	return NewComponent(`<rect></rect>`)
}

func Line() Component {
	return NewComponent(`<line></line>`)
}

//========================================
// Templating Funcs
//========================================

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
func If(condition bool, trueComponent, falseComponent Component) Component {
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
