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
			H1().Text("Hello, World!"),
		),
	)
}

func Test_G(t *testing.T) {

	c := P().TagName()
	fmt.Println(c)

}
