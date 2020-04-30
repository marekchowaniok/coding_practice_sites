package accumulate

// Words used as main type
type Words []string

// Accumulate function
func Accumulate(input Words, function func(string) string) (output Words) {

	for _, word := range input {
		output = append(output, function(word))
	}

	return output
}
