package gsc

import "fmt"

func NavItem(title string, href string) string {
	id, _ := UtilRandStr(32)
	return fmt.Sprintf( /*html*/ `
		<li class='flex w-full rounded border hover:bg-gray-100'>
			<a id='%s' href='%s' class='p-4 w-full text-sm'>%s</a>
		</li>
		<script>
			(() => {
				let navItem = document.querySelector('#%s')
				let href = navItem.getAttribute('href')
				let path = window.location.pathname
				if (path == href) {
					navItem.classList.add('bg-gray-100')
				}
			})()
		</script>
	`, id, href, title, id)
}
