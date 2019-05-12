package tests

import "github.com/wwq1988/stream/outter"

type Some struct {
	A string `stream:"sort,unique"`
	B string `stream:"sort,unique"`
	C *Some  `stream:"sort,unique"`
	D *outter.Some
}

type B struct {
}
