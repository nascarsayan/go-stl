package util

// Reverse takes a string as input and returns its reverse.
// Reverse takes a string as input and returns its reverse.
func Reverse(str string) string {
	rev := []rune(str)
	for i := 0; i < len(rev)/2; i++ {
		rev[i], rev[len(rev)-i-1] = rev[len(rev)-i-1], rev[i]
	}
	return string(rev)
}
