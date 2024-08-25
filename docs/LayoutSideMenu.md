```go
func LayoutSideMenu(id string, title string, subtext string, navMenu string, content string) string {
	return Group(
		Style(
			fmt.Sprintf(`
			#%s {
				grid-template-rows: auto auto;
				grid-template-columns: auto;
				grid-template-areas: 
					"nav"
					"content"
				;
			}
			`, id),
		),
		Div(`id=`+id+` class='grid relative height-[100vh]'`,
			Header(`id='header' style='grid-area:nav;' class='flex p-4 bg-white z-30 fixed w-full h-[85px] border-b justify-between items-center'`,
				Div(`class='flex flex-col gap-1'`,
					H1(`class='text-xl font-bold'`, title),
					P(`class='text-sm'`, subtext),
				),
				IconBars(`id='bars' class='flex items-center lg:hidden'`),
				IconX(`id='x' class='flex items-center hidden lg:hidden'`),
			),
			Nav(`id='menu' class='fixed z-30 top-[85px] w-[300px] md:w-[350px] bg-white h-[100%] hidden flex-col overflow-auto lg:flex'`, navMenu),
			Div(`id='overlay' class='z-20 opacity-50 hidden bg-black top-[85px] h-[100%] w-full fixed lg:hidden'`, ``),
			Div(`style='grid-area:content;' id='content' class='relative top-[85px] lg:pl-[350px]'`, content),
			Script(``, `
				(() => {
					let grid = document.querySelector('#`+id+`')
					let header = grid.querySelector('#header')
					let bars = grid.querySelector('#bars')
					let x = grid.querySelector('#x')
					let menu = grid.querySelector('#menu')
					let overlay = grid.querySelector('#overlay')
					let toggleMenu = () => {
						overlay.classList.toggle('hidden')
						menu.classList.toggle('hidden')
						menu.classList.toggle('flex')
						bars.classList.toggle('hidden')
						x.classList.toggle('hidden')
					}
					bars.addEventListener('click', toggleMenu)
					x.addEventListener('click', toggleMenu)
					overlay.addEventListener('click', toggleMenu)
				})()
			`),
		),
	)
}
```