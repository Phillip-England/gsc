package gsc

import (
	"net/http"
	"testing"

	"github.com/Phillip-England/vbf"
)

func handler() string {
	return HTMLDoc(
		Head("", "go simple components",
			CSSLink("/static/output.css"),
			Meta(`name="viewport" content="width=device-width, initial-scale=1.0"`),
		),
		Body("",
			LayoutSideMenu("main-grid", "gsc", "go simple components",
				Ul(`class='p-2 flex flex-col gap-2 border-r h-full'`, Map(func(item string) string {
					return Li(`class='flex text-sm border rounded'`, item)
				},
					A(`href='/' class='p-4 w-full hover:bg-light-gray' active-link='bg-light-gray'`, "Home"),
					A(`href='/about' class='p-4 w-full hover:bg-light-gray' active-link='bg-light-gray'`, "About"),
					A(`href='/contact' class='p-4 w-full hover:bg-light-gray' active-link='bg-light-gray'`, "Contact"),
				)),
				Main(`class='p-8 md:py-12 md:px-32'`,
					Article(`class='flex flex-col gap-2'`,
						TitleAndText("Keep it Simple", "Stop overthinking web development. HTML, CSS, and Javascript. That's it."),
					),
				),
			),
			ScriptActiveLink(),
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
