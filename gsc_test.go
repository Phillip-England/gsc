package gsc

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Phillip-England/ffh"
	"github.com/Phillip-England/vbf"
)

func handler(mdPath string, navMenu string) string {
	mdContent, _ := Markdown(mdPath)
	return HTMLDoc(
		Head("", "go simple components",
			CSSLink("/static/output.css"),
			JSSignal(),
			Meta(`name="viewport" content="width=device-width, initial-scale=1.0"`),
		),
		Body("",
			LayoutSideMenu("main-grid", "gsc layouts", "layouts made with go simple components",
				navMenu,
				Main(`class='p-8 md:py-12 md:px-16'`,
					Article(`id='main-article'`, mdContent),
				),
			),
			JSActiveLink(),
			JSStyleMdContainer("#main-article"),
		),
	)
}

func Test_Gsc(t *testing.T) {

	mux, gCtx := vbf.VeryBestFramework()

	anchors := []string{}
	mdPaths := []string{}
	funcNames := []string{}

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".go") && !strings.Contains(path, "_test") {
			parts := strings.Split(path, ".")
			funcName := parts[0]
			fileContent, err := ffh.ReadFile(path)
			if err != nil {
				return err
			}
			componentFunc, err := ffh.ExtractFuncByName(fileContent, funcName)
			if err != nil {
				return err
			}
			mdPath := "./docs/" + funcName + ".md"
			err = os.WriteFile(mdPath, []byte("```go\n"+componentFunc+"```"), 0644)
			if err != nil {
				return err
			}
			mdPaths = append(mdPaths, mdPath)
			funcNames = append(funcNames, funcName)
			anchors = append(anchors, A(`href='/`+funcName+`' class='p-4 w-full hover:bg-light-gray' active-link='bg-light-gray'`, funcName))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	navMenu := Ul(`class='p-2 flex flex-col gap-2 border-r h-full'`, Map(func(anchor string) string {
		return Li(`class='flex text-sm border rounded'`, anchor)
	}, anchors...))

	for i, mdPath := range mdPaths {
		vbf.AddRoute("GET /"+funcNames[i], mux, gCtx, func(w http.ResponseWriter, r *http.Request) {
			vbf.WriteHTML(w, handler(mdPath, navMenu))
		}, vbf.MwLogger)
	}

	_ = vbf.Serve(mux, "8080")

}
