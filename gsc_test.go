package gsc

import (
	"net/http"
	"testing"

	"github.com/Phillip-England/vbf"
)

func handler() string {
	bulkLinks := ""
	for i := 0; i < 10; i++ {
		bulkLinks = bulkLinks + "<a href='/'>home</a>"
	}
	return HTMLDoc(
		Head("", "go simple components",
			CSSLink("/static/output.css"),
			Meta(`name="viewport" content="width=device-width, initial-scale=1.0"`),
		),
		Body("",
			GridOneCol("main-grid", "gsc", "go simple components", bulkLinks, "<p class='h-[3000px]'>yooo</p>"),
		),
	)
}

func Test_Gsc(t *testing.T) {

	mux, gCtx := vbf.VeryBestFramework()

	vbf.AddRoute("GET /", mux, gCtx, func(w http.ResponseWriter, r *http.Request) {
		vbf.WriteHTML(w, handler())
	}, vbf.MwLogger)

	_ = vbf.Serve(mux, "8080")

}
