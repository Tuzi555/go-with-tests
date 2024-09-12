package iteration

func Repeat(input string, repeats int) string {
	var repeated string
	for range repeats {
		repeated += input
	}
	return repeated
}
