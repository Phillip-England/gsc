package gsc

import (
	"fmt"
	"testing"
)

func FlexCol(c Component) Component {
	c = c.Class("flex flex-col")
	return c
}

func TextSm(c Component) Component {
	c = c.Class("text-sm")
	return c
}

func BoldHeader(c Component) Component {
	return c.Class("text-lg font-bold")
}

func Test_G(t *testing.T) {
	c := HTMLDoc().In(
		Div().Of(FlexCol).In(
			P().Of(TextSm),
			H1().Of(BoldHeader),
		),
		Map(func(item string) Component {
			return P().Text(item)
		}, "apple", "orange", "banana"),
	)
	fmt.Println(c)
}
