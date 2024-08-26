package gsc

import (
	"fmt"
	"testing"
)

func Test_G(t *testing.T) {

	f, err := NewGscFile("./component.html")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Content)
}
