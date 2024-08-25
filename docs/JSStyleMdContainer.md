```go
func JSStyleMdContainer(containerSelector string) string {
	return Script("", `
		(() => {
			let container = document.querySelector("`+containerSelector+`")
			container.classList.add('flex', 'flex-col')
			container.querySelectorAll("h1").forEach(el => {
				el.classList.add('text-3xl', 'font-bold', 'mb-16')
			})
			container.querySelectorAll("h2").forEach(el => {
				el.classList.add('text-xl', 'font-bold', "mb-2")
			})
			let h3 = container.querySelectorAll("h3")
			let h4 = container.querySelectorAll("h4")
			let h5 = container.querySelectorAll("h5")
			let h6 = container.querySelectorAll("h6")
		 	container.querySelectorAll('p').forEach(el => {
				el.classList.add('text-sm', 'mb-12')
			})
			container.querySelectorAll('ol').forEach(el => {
				el.classList.add('text-sm', 'flex', 'flex-col', 'mb-12', 'gap-1', 'ml-4')
				let count = 1
				let listItems = el.querySelectorAll('li')
				listItems.forEach(li => {
					li.textContent = count + ". " + li.textContent
					count++
				})
			})
		
		})()
		
	`)
}
```