package gsc

import "fmt"

func IconX(attr string) string {
	return fmt.Sprintf( /*html*/ `
			<div %s>
				<svg class="w-8 h-8" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
					<path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18 17.94 6M18 18 6.06 6"/>
				</svg>
			</div>
	`, attr)
}
