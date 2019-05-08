package tests

type Some struct {
	A string
}

func (s *Some) Compare(src *Some) int {
	if s.A == src.A {
		return 0
	}
	if s.A < src.A {
		return -1
	}
	return 1
}
