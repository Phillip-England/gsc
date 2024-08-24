package gsc

import (
	"bytes"
	"fmt"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/renderer/html"
)

func CodeBlock(language string, filepathMarker string, content string) string {
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
	title := buf2.String()
	randStr, _ := GetRandomString(32)
	title = strings.Replace(title, "<pre", fmt.Sprintf("<pre id='%s'", randStr), 1)
	title = strings.Replace(title, "style=\"", `class="text-sm flex flex-col relative pb-2 pr-4"  style="overflow-x:auto;`, 1)
	filenameHTML := ``
	if filepathMarker != "" {
		filenameHTML = `<div class='absolute left-6'><p class='text-sm'>` + filepathMarker + `</p></div>`
	}
	copyIndicator := `<div id='copy-indicator' class='hidden flex items-center absolute right-0'><svg class="w-6 h-6 text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M9 2a1 1 0 0 0-1 1H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2a1 1 0 0 0-1-1H9Zm1 2h4v2h1a1 1 0 1 1 0 2H9a1 1 0 0 1 0-2h1V4Zm5.707 8.707a1 1 0 0 0-1.414-1.414L11 14.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4Z" clip-rule="evenodd"/></svg></div>`
	copyBtn := `<div id='copy-btn' class='flex items-center cursor-pointer absolute right-0'><svg class="w-6 h-6 text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M18 3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1V9a4 4 0 0 0-4-4h-3a1.99 1.99 0 0 0-1 .267V5a2 2 0 0 1 2-2h7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 7.054V11H4.2a2 2 0 0 1 .281-.432l2.46-2.87A2 2 0 0 1 8 7.054ZM10 7v4a2 2 0 0 1-2 2H4v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3Z" clip-rule="evenodd"/></svg></div>`
	title = strings.Replace(title, "<code>", `<div class='flex flex-row items-center gap-2 p-6 relative'>`+filenameHTML+copyIndicator+copyBtn+`</div><code>`, 1)

	title = title + fmt.Sprintf( /*html*/ `
		<script>
			(() => {
				let block = document.querySelector("#%s")
				let copyBtn = block.querySelector('#copy-btn')
				let copyIndicator = block.querySelector('#copy-indicator')
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
					}
				}
				let htmlData = newLines.join("\n")
				let firstTenChars = htmlData.slice(0, 10)
				let parts = firstTenChars.split(" ")
				let firstPart = parts[0]
				let secondPartFirstChar = parts[1][0]
				let isNum = parseInt(secondPartFirstChar, 10)
				if (firstPart && Number.isInteger(isNum)) {
					htmlData = htmlData.slice(htmlData.indexOf(secondPartFirstChar)+1, htmlData.length)
				}
				block.addEventListener('click', (e) => {
					copyBtn.classList.add('hidden')
					copyIndicator.classList.remove('hidden')
					navigator.clipboard.writeText(htmlData)
					setTimeout(() => {
						copyBtn.classList.remove('hidden')
						copyIndicator.classList.add('hidden')
					}, 1000)
				})
			})()
			
		</script>
	`, randStr)
	return title
}
