package gsc

import (
	"fmt"
	"testing"
)

func Test_G(t *testing.T) {
	c := HTMLDoc().In(
		Div().In(
			P().Class("text-sm"),
			H1().Class("text-xl font-bold"),
		),
		Map(func(item string) Component {
			return P().Text(item)
		}, "apple", "orange", "banana"),
	)
	fmt.Println(c)
}
