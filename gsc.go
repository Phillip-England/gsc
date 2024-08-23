package gsc

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func GetRandomString(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < n; i++ {
		bytes[i] = letters[int(bytes[i])%len(letters)]
	}

	return string(bytes), nil
}

func StrJoin(strs ...string) string {
	return strings.Join(strs, "\n")
}

func HTMLDoc(head string, content ...string) string {
	output := ""
	for _, str := range content {
		output = output + str
	}
	return fmt.Sprintf( /*html*/ `
		<!DOCTYPE>
		<html>
			%s
			<body class='select-none'>%s</body>
		</html>
	`, head, output)
}

func CodeBlock(language string, content string) string {
	lines := []string{}
	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			fmt.Println(line)
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
	title := buf2.String()
	randStr, _ := GetRandomString(32)
	title = strings.Replace(title, "<pre", fmt.Sprintf("<pre id='%s'", randStr), 1)
	title = strings.Replace(title, "style=\"", `class="text-sm flex flex-col relative p-2 select-text"  style="overflow-x:auto;`, 1)
	title = strings.Replace(title, "<code>", `<div class='absolute right-2'><div id='copy-btn'><svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M18 3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1V9a4 4 0 0 0-4-4h-3a1.99 1.99 0 0 0-1 .267V5a2 2 0 0 1 2-2h7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 7.054V11H4.2a2 2 0 0 1 .281-.432l2.46-2.87A2 2 0 0 1 8 7.054ZM10 7v4a2 2 0 0 1-2 2H4v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3Z" clip-rule="evenodd"/></svg></div><div id='copy-indicator' class='hidden'><svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M9 2a1 1 0 0 0-1 1H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2a1 1 0 0 0-1-1H9Zm1 2h4v2h1a1 1 0 1 1 0 2H9a1 1 0 0 1 0-2h1V4Zm5.707 8.707a1 1 0 0 0-1.414-1.414L11 14.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4Z" clip-rule="evenodd"/></svg></div></div><code>`, 1)

	title = title + fmt.Sprintf( /*html*/ `
		<script>
			let block = document.querySelector("#%s")
			let copyBtn = block.querySelector('#copy-btn')
			let copyIndicator = block.querySelector('#copy-indicator')
			console.log(copyBtn, copyIndicator)
			let lines = block.textContent.split("\n")
			let newLines = []
			for (let i = 0; i < lines.length; i++) {
				let line = lines[i]
				for (let i2 = 0; i2 < line.length; i2++) {
					let char = line[i2]
					let num = parseInt(char, 10)
					if (Number.isInteger(num) && num.toString() === char.trim() || char == " ") {
						continue
					} 
					newLines.push(line.slice(i2))
					break
z				}
			}
			let htmlData = newLines.join("\n")
			block.addEventListener('click', (e) => {
				copyBtn.classList.toggle('hidden')
				copyIndicator.classList.toggle('hidden')
				navigator.clipboard.writeText(htmlData)
				setTimeout(() => {
					copyBtn.classList.toggle('hidden')
					copyIndicator.classList.toggle('hidden')
				}, 1000)
			})
			
		</script>
	`, randStr)
	return title
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

func Navbar(title string, subtext string, navmenu string) string {
	return fmt.Sprintf( /*html*/ `
		<nav id='navbar' class='p-4 border-b flex flex-row justify-between z-30 bg-white relative'>
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
		<div id='navbar-overlay' class='absolute top-0 left-0 h-full w-full opacity-50 bg-black z-20 hidden'></div>

		<script>
			(() => {
				let navbar = document.querySelector("#navbar")
				let bars = navbar.querySelector('#navbar-bars')
				let x = navbar.querySelector('#navbar-x')
				let overlay = document.querySelector('#navbar-overlay')
				let toggleNav = (e) => {
					bars.classList.toggle('hidden')
					x.classList.toggle('hidden')
					overlay.classList.toggle('hidden')
				}
				bars.addEventListener('click', toggleNav)
				x.addEventListener('click', toggleNav)
				overlay.addEventListener('click', toggleNav)
			})()
		</script>

	`, title, subtext)
}

func MainElement(content ...string) string {
	contentStr := strings.Join(content, "\n")
	return fmt.Sprintf( /*html*/ `
		<main class='p-4 flex flex-col gap-4'>
			%s
		</main>
	`, contentStr)
}

func Article(content ...string) string {
	contentStr := strings.Join(content, "\n")
	return fmt.Sprintf( /*html*/ `
		<article class='flex flex-col gap-2'>
			%s
		</article>
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
