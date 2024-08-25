package gsc

import "strings"

func FlexStack(items ...string) string {
	return Div("class='flex flex-col gap-4'", strings.Join(items, "\n"))
}
