package gsc

import "fmt"

func CSSLink(path string) string {
	return fmt.Sprintf(`<link rel="stylesheet" href="%s">`, path)
}

func Meta(content string) string {
	return fmt.Sprintf(`<meta %s></meta>`, content)
}

func P(text string) string {
	return fmt.Sprintf(`<p class="text-sm">%s</p>`, text)
}

func H1(text string) string {
	return fmt.Sprintf(`<h1 class="text-2xl font-bold">%s</h1>`, text)
}

func H2(text string) string {
	return fmt.Sprintf(`<h2 class="text-xl font-bold">%s</h2>`, text)
}

func H3(text string) string {
	return fmt.Sprintf(`<h3 class="text-lg font-bold">%s</h3>`, text)
}
