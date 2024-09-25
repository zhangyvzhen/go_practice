package iteration

func Repeat(c string, n int) (s string) {
	for i := 0; i < n; i++ {
		s += c
	}
	return
}
