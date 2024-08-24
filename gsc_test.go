package gsc

import (
	"net/http"
	"testing"

	"github.com/Phillip-England/ffh"
	"github.com/Phillip-England/vbf"
)

func homePage() string {
	fileTxt, _ := ffh.ReadFile("./containers.go")
	code, _ := ffh.ExtractFuncByName(fileTxt, "HTMLDoc")
	return HTMLDoc(
		Head("go simple components",
			CSSLink("/static/output.css"),
			Meta(`name="viewport" content="width=device-width, initial-scale=1.0"`),
		),
		Navbar("gsc", "go simple components", Navmenu(
			NavItem("Home", "/"),
			NavItem("About", "/about"),
		)),
		MainElement(
			H1("Navigation"),
			Article(
				H3("HTMLDoc"),
				P("outputs the shell of an html document"),
				CodeBlock("go", "gsc.go", code),
			),
			Article(
				H3("MainElement"),
				P("a flex container which can wrap the main content for a good content-based layout"),
				CodeBlock("go", "gsc.go", code),
			),
		),
	)
}

func Test_VBF(t *testing.T) {

	mux, gCtx := vbf.VeryBestFramework()

	vbf.AddRoute("GET /", mux, gCtx, func(w http.ResponseWriter, r *http.Request) {
		vbf.WriteHTML(w, homePage())
	}, vbf.MwLogger)

	err := vbf.Serve(mux, "8080")
	if err != nil {
		panic(err)
	}
}
