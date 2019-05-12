package tests

import "github.com/wwq1988/stream/outter"

type Some struct {
	A string
	B string
	C *Some
	D *outter.Some
	E *string
	F *string
}

type B struct {
}
