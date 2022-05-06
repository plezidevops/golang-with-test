package iteration

func Repeat(text string) string {
	//return strings.Repeat(text, 5)
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + text
	}
	return repeated
}
