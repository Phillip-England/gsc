package gsc

import (
	"fmt"
	"strings"
)

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

func Body(attr string, content string) string {
	return fmt.Sprintf( /*html*/ `<body %s>%s</body>`, attr, content)
}

func Head(title string, content ...string) string {
	output := ""
	for _, str := range content {
		output = output + str
	}
	return fmt.Sprintf(`<head>%s<title>%s</title></head>`, output, title)
}

func MainElement(content ...string) string {
	contentStr := strings.Join(content, "\n")
	return fmt.Sprintf( /*html*/ `
		<main class='p-4 flex flex-col gap-12'>
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

func TwoColumnGrid() string {

}
