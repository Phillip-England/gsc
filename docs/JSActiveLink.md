```go
func JSActiveLink() string {
	return Script(``, `
		(() => {
			let links = document.querySelectorAll('[active-link]')
			let path = window.location.pathname
			for (let i=0; i < links.length; i++) {
				let link = links[i]
				let href = link.getAttribute('href')
				let activeStyle = link.getAttribute('active-link')
				let parts = activeStyle.split(' ')
				if (href==path) {
					for (let i2=0; i2 < parts.length; i2++) {
						link.classList.add(parts[i2])
					}
				}
			}
		})()
	`)
}
```