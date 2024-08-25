package gsc

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Phillip-England/ffh"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/renderer/html"
)

func Markdown(filepath string) (string, error) {
	mdContent, err := ffh.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	md := goldmark.New(
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
	var buf bytes.Buffer
	if err := md.Convert([]byte(mdContent), &buf); err != nil {
		return "", err
	}
	html := buf.String()
	// giving all <pre> a class so I can access them via JS
	preClassName, err := UtilRandStr(32)
	if err != nil {
		return "", err
	}
	html = strings.ReplaceAll(html, "<pre", fmt.Sprintf(`<pre class="%s"`, preClassName))
	html = strings.ReplaceAll(html, fmt.Sprintf(`<pre class="%s" style="`, preClassName), fmt.Sprintf(`<pre class="%s text-sm flex flex-col relative p-2" style="overflow-x:auto;`, preClassName))
	copyButton := `<div id='copy-button' class='flex items-center cursor-pointer absolute right-0'><svg class="w-6 h-6 text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M18 3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1V9a4 4 0 0 0-4-4h-3a1.99 1.99 0 0 0-1 .267V5a2 2 0 0 1 2-2h7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 7.054V11H4.2a2 2 0 0 1 .281-.432l2.46-2.87A2 2 0 0 1 8 7.054ZM10 7v4a2 2 0 0 1-2 2H4v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3Z" clip-rule="evenodd"/></svg></div>`
	copyIndicator := `<div id='copy-indicator' class='hidden flex items-center absolute right-0'><svg class="w-6 h-6 text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M9 2a1 1 0 0 0-1 1H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2a1 1 0 0 0-1-1H9Zm1 2h4v2h1a1 1 0 1 1 0 2H9a1 1 0 0 1 0-2h1V4Zm5.707 8.707a1 1 0 0 0-1.414-1.414L11 14.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4Z" clip-rule="evenodd"/></svg></div>`
	html = strings.ReplaceAll(html, "<code>", `<div class='flex flex-row items-center gap-2 p-2 relative'>`+copyIndicator+copyButton+`</div><code class='whitespace-pre-wrap'>`)
	html = html + Script("", `
		(() => {
			let blocks = document.querySelectorAll('.`+preClassName+`')
			for (let i = 0; i < blocks.length; i++) {
				let block = blocks[i]
				let copyButton = block.querySelector('#copy-button')
				let copyIndicator = block.querySelector('#copy-indicator')
				let lines = block.textContent.split('\n')
				let newLines = []
				for (let i2 = 0; i2 < lines.length; i2++) {
					let line = lines[i2]
					for (let i3 = 0; i3 < line.length; i3++) {
						let char = line[i3]
						let num = parseInt(char, 10)
						if (Number.isInteger(num)) {
							continue
						} 
						newLines.push(line.slice(i2))
						break
					}
				}
				let htmlData = newLines.join("\n")
				let firstTenChars = htmlData.slice(0, 10)
				let parts = firstTenChars.split(".")
				let firstPart = parts[0]
				let firstPartFirstChar = parts[0][0]
				let isNum = parseInt(firstPartFirstChar, 10)
				if (firstPart && Number.isInteger(isNum)) {
					htmlData = htmlData.slice(htmlData.indexOf(firstPartFirstChar)+1, htmlData.length)
				}
				block.addEventListener('click', (e) => {
					copyButton.classList.add('hidden')
					copyIndicator.classList.remove('hidden')
					navigator.clipboard.writeText(htmlData)
					setTimeout(() => {
						copyButton.classList.remove('hidden')
						copyIndicator.classList.add('hidden')
					}, 1000)
				})
			} 
		})()
	`)
	return html, nil
}
