package tests

type Some struct {
	A string
}

func (s *Some) Compare(src *Some) bool {
	return s.A == src.A
}
