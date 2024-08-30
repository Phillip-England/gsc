package gsc

import (
	"fmt"
	"testing"
)

func Test_G(t *testing.T) {
	c := HTMLDoc().Content(
		Div().Content(
			P().Class("text-sm"),
			H1().Class("text-xl font-bold"),
		),
	)
	fmt.Println(c)
}
