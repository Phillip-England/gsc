```go
func Navbar(title string, subtext string, navItems ...string) string {
	return fmt.Sprintf( /*html*/ `
		<nav id='navbar' class='p-4 border-b flex flex-row justify-between z-30 bg-white relative h-[90px]'>
			<div>
				<h1 class='text-lg font-bold'>%s</h1>
				<p class='text-sm'>%s</p>
			</div>
			<div id='navbar-bars' class='flex items-center'>
				<svg class="w-8 h-8" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
					<path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 7h14M5 12h14M5 17h14"/>
				</svg>
			</div>
			<div id='navbar-x' class='flex items-center hidden'>
				<svg class="w-8 h-8" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
			  		<path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18 17.94 6M18 18 6.06 6"/>
				</svg>
			</div>
		</nav>
		<div id='navbar-overlay' class='absolute top-0 left-0 h-full w-full opacity-50 bg-black z-20 hidden'></div>
		<div id='navbar-menu' class='hidden relative z-30'>
			<div class='bg-white w-4/5 h-full fixed top-[90px]'>
				<div class=''></div>
				<ul class='flex flex-col gap-2 p-2'>%s</ul>
			</div>	
		</div>

		<script>
			(() => {
				let navbar = document.querySelector("#navbar")
				let bars = navbar.querySelector('#navbar-bars')
				let x = navbar.querySelector('#navbar-x')
				let overlay = document.querySelector('#navbar-overlay')
				let menu = document.querySelector('#navbar-menu')
				let toggleNav = (e) => {
					bars.classList.toggle('hidden')
					x.classList.toggle('hidden')
					overlay.classList.toggle('hidden')
					menu.classList.toggle('hidden')
				}
				bars.addEventListener('click', toggleNav)
				x.addEventListener('click', toggleNav)
				overlay.addEventListener('click', toggleNav)
			})()
		</script>

	`, title, subtext, strings.Join(navItems, "\n"))
}
```