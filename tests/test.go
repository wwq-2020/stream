package tests

type Some struct {
	A string `collections:"sort,unique"`
	B string `collections:"sort,unique"`
	C *Some  `collections:"sort,unique"`
}

type B struct {
}
