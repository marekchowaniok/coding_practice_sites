package strain

type Ints []int

func (src Ints) Keep(predicate func(int) bool) (output Ints) {
	//if pred.
	for _, item := range src {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return
}

func (src Ints) Discard(predicate func(int) bool) (output Ints) {
	for _, item := range src {
		if !predicate(item) {
			output = append(output, item)
		}
	}
	return
}

type Strings []string

func (src Strings) Keep(predicate func(s string) bool) (output Strings) {
	for _, item := range src {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return
}

type Lists [][]int

func (src Lists) Keep(predicate func(l []int) bool) (output Lists) {
	for _, item := range src {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return
}
