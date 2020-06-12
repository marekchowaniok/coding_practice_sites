package reverse

// Reverse of string
func Reverse(input string) string {
	out := ""
	for _, ch := range input {
		out = string(ch) + out
	}
	return out
}
