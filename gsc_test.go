package gsc

import (
	"net/http"
	"testing"

	"github.com/Phillip-England/ffh"
	"github.com/Phillip-England/vbf"
)

func homePage() string {
	fileTxt, _ := ffh.ReadFile("./gsc.go")
	code, _ := ffh.ExtractFuncByName(fileTxt, "HTMLDoc")
	return HTMLDoc(
		Head(CSSLink("/static/output.css")),
		Navbar("gsc", "go simple components", ""),
		MainElement(
			H1("HTMLDoc"),
			P("outputs the shell of an html document"),
			CodeBlock("go", code),
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
