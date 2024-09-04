package gsc

import (
	"fmt"
	"testing"
)

func Layout() Component {
	return HTMLDoc().In(
		Head().In(
			Meta().Name("viewport").Content("width=device-width, initial-scale=1.0"),
			Link().Rel("stylesheet").Href("/static/css/output.css"),
			Title().Text("Receipt Tracker"),
		),
		Body().In(
			H1().ID("header").Text("yuooooooo"),
		),
	)
}

func Test_G(t *testing.T) {

	c := Layout().Query("#header", func(child Component) Component {
		return child.Class("text-xl font-bold")
	})
	fmt.Println(c)

}
