package gsc

func TitleAndText(title string, text string) string {
	return Div(`class='flex flex-col gap-2'`,
		H2(`class='font-bold text-xl'`, title),
		P(`class='text-sm'`, text),
	)
}
