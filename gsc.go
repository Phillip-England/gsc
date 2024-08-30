package gsc

import (
	"fmt"
	"strings"
)

//========================================
// COMPONENTS
//========================================

func Link(attr string) string {
	return fmt.Sprintf(`<link %s />`, attr)
}

func Meta(content string) string {
	return fmt.Sprintf(`<meta %s></meta>`, content)
}

func P(attr string, text string) string {
	return fmt.Sprintf(`<p %s>%s</p>`, attr, text)
}

func H1(attr string, text string) string {
	return fmt.Sprintf(`<h1 %s>%s</h1>`, attr, text)
}

func H2(attr string, text string) string {
	return fmt.Sprintf(`<h2 %s>%s</h2>`, attr, text)
}

func H3(attr string, text string) string {
	return fmt.Sprintf(`<h3 %s>%s</h3>`, attr, text)
}

func A(attr string, content ...string) string {
	return fmt.Sprintf(`<a %s>%s</a>`, attr, strings.Join(content, "\n"))
}

func Head(attr string, title string, content ...string) string {
	return fmt.Sprintf(`<head %s>%s<title>%s</title></head>`, attr, strings.Join(content, "\n"), title)
}

func HTMLDoc(content ...string) string {
	return fmt.Sprintf(`<!DOCTYPE><html>%s</html>`, strings.Join(content, "\n"))
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
	return fmt.Sprintf(`<main %s>%s</main>`, attr, strings.Join(content, "\n"))
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
	return fmt.Sprintf(`<article %s>%s</article>`, attr, output)
}

func Body(attr string, content ...string) string {
	return fmt.Sprintf(`<body %s>%s</body>`, attr, strings.Join(content, "\n"))
}

//========================================
// ATTRIBUTES
//========================================

func Attr(funcs ...string) string {
	output := ""
	for _, str := range funcs {
		output = output + str
	}
	return output
}

func Class(className string) string {
	return `class="` + className + `" `
}

func Src(s string) string {
	return `src="` + s + `" `
}

func Rel(className string) string {
	return `rel="` + className + `" `
}

func Href(s string) string {
	return `href="` + s + `" `
}

func Name(s string) string {
	return `name="` + s + `" `
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
