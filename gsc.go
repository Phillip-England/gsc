package gsc

import (
	"fmt"
	"strings"
)

//========================================
// COMPONENTS
//========================================

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

func A(attr string, content ...string) string {
	return fmt.Sprintf(`<a %s>%s</a>`, attr, strings.Join(content, "\n"))
}

func Head(attr string, title string, content ...string) string {
	return fmt.Sprintf(`<head %s>%s<title>%s</title></head>`, attr, strings.Join(content, "\n"), title)
}

func HTMLDoc(content ...string) string {
	return fmt.Sprintf(`
		<!DOCTYPE><html>%s</html>
	`, strings.Join(content, "\n"))
}

func Div(attr string, content ...string) string {
	return fmt.Sprintf(`<div %s>%s</div>`, attr, strings.Join(content, "\n"))
}

func Header(attr string, content ...string) string {
	return fmt.Sprintf(`<header %s>%s</header>`, attr, strings.Join(content, "\n"))
}

func Li(attr string, content ...string) string {
	return fmt.Sprintf(`<li %s>%s</li>`, attr, strings.Join(content, "\n"))
}

func Main(attr string, content ...string) string {
	return fmt.Sprintf(`
		<main %s>
			%s
		</main>
	`, attr, strings.Join(content, "\n"))
}

func Nav(attr string, content ...string) string {
	return fmt.Sprintf(`<nav %s>%s</nav>`, attr, strings.Join(content, "\n"))
}

func Ol(attr string, items ...string) string {
	return fmt.Sprintf(`<ol %s>%s</ol>`, attr, strings.Join(items, "\n"))
}

func Script(attr string, text string) string {
	return fmt.Sprintf(`<script %s>%s</script>`, attr, text)
}

func Ul(attr string, items ...string) string {
	return fmt.Sprintf(`<ul %s>%s</ul>`, attr, strings.Join(items, "\n"))
}

func Article(attr string, content ...string) string {
	output := strings.Join(content, "\n")
	return fmt.Sprintf(`
		<article %s>
			%s
		</article>
	`, attr, output)
}

func Body(attr string, content ...string) string {
	return fmt.Sprintf(`<body %s>%s</body>`, attr, strings.Join(content, "\n"))
}

//========================================
// ATTRIBUTES
//========================================

func Attr(funcs ...func(string) string) string {
	output := ""
	for _, fn := range funcs {
		output = fn(output)
	}
	return output
}

func Class(className string) string {
	return `class="` + className + `" `
}

//========================================
// TEMPLATING FUNCTIONS
//========================================

func Group(components ...string) string {
	return strings.Join(components, "\n")
}

func Map(f func(item string) string, items ...string) string {
	output := ""
	for _, item := range items {
		output = output + f(item) + "\n"
	}
	return output
}
