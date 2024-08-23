package gsc

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func StrJoin(strs ...string) string {
	return strings.Join(strs, "\n")
}

func HTMLDoc(meta string, content ...string) string {
	output := ""
	for _, str := range content {
		output = output + str
	}
	return fmt.Sprintf( /*html*/ `
		<!DOCTYPE>
		<html>
			<head>%s</head>
			<body class='select-none'>%s</body>
		</html>
	`, meta, output)
}

func CodeBlock(language string, content string) string {
	lines := []string{}
	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	content = strings.Join(lines, "\n")
	content = fmt.Sprintf("\n```%s\n%s\n```", language, content)
	markdown2 := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf2 bytes.Buffer
	if err := markdown2.Convert([]byte(content), &buf2); err != nil {
		panic(err)
	}
	title2 := buf2.String()
	title2 = strings.Replace(title2, "style=\"", `class="text-sm flex flex-col relative p-2"  style="overflow-x:auto;`, 1)
	title2 = strings.Replace(title2, "<code>", `<div class='absolute right-2'><svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M18 3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1V9a4 4 0 0 0-4-4h-3a1.99 1.99 0 0 0-1 .267V5a2 2 0 0 1 2-2h7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 7.054V11H4.2a2 2 0 0 1 .281-.432l2.46-2.87A2 2 0 0 1 8 7.054ZM10 7v4a2 2 0 0 1-2 2H4v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3Z" clip-rule="evenodd"/></svg></div><code>`, 1)
	return title2
}

func Head(content ...string) string {
	output := ""
	for _, str := range content {
		output = output + str
	}
	return fmt.Sprintf(`<head>%s</head>`, output)
}

func CSSLink(path string) string {
	return fmt.Sprintf(`<link rel="stylesheet" href="%s">`, path)
}

func Navbar(title string, subtext string) string {
	return fmt.Sprintf( /*html*/ `
		<nav id='navbar' class='p-4 border-b flex flex-row justify-between'>
			<div>
				<h1 class='text-lg font-bold'>%s</h1>
				<p>%s</p>
			</div>
			<div id='navbar-bars' class='flex items-center'>
				<svg class="w-8 h-8" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
					<path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 7h14M5 12h14M5 17h14"/>
				</svg>
			</div>
			<div id='navbar-x' class='flex items-center hidden'>
				<svg class="w-8 h-8" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
			  		<path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18 17.94 6M18 18 6.06 6"/>
				</svg>
			</div>
		</nav>

		<script>
			(() => {
				let navbar = document.querySelector("#navbar")
				let bars = navbar.querySelector('#navbar-bars')
				let x = navbar.querySelector('#navbar-x')
				let toggleNav = (e) => {
					bars.classList.toggle('hidden')
					x.classList.toggle('hidden')
				}
				bars.addEventListener('click', toggleNav)
				x.addEventListener('click', toggleNav)
			})()
		</script>

	`, title, subtext)
}

func Root(content ...string) string {
	contentStr := strings.Join(content, "\n")
	return fmt.Sprintf( /*html*/ `
		<div id="root" class='p-4 flex flex-col gap-4'>
			%s
		</div>
	`, contentStr)
}

func Article(content ...string) string {
	contentStr := strings.Join(content, "\n")
	return fmt.Sprintf( /*html*/ `
	
	`, contentStr)
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
