package gsc

func Map(f func(item string) string, items ...string) string {
	output := ""
	for _, item := range items {
		output = output + f(item) + "\n"
	}
	return output
}
