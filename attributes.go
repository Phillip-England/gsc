package gsc

import (
	"fmt"
	"strings"
)

func (c Component) Attr(attrName string, value string) Component {
	strComponent := string(c)
	rightBracketIndex := strings.Index(strComponent, ">")
	if rightBracketIndex == -1 {
		return ""
	}
	if strings.Contains(strComponent, attrName+"=\"") {
		return NewComponent(strings.Replace(strComponent, ``+attrName+`="`, ``+attrName+`="`+value+` `, 1))
	}
	if string(strComponent[rightBracketIndex-1]) == "/" {
		return NewComponent(strings.Replace(strComponent, "/>", ` `+attrName+`="`+value+`"/>`, 1))
	}
	return NewComponent(strings.Replace(strComponent, ">", ` `+attrName+`="`+value+`">`, 1))
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
