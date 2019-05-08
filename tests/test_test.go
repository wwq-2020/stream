package tests

import (
	"reflect"
	"testing"
)

func TestConcate(t *testing.T) {
	data1 := []*Some{&Some{"hello"}}
	data2 := []*Some{&Some{"world"}}
	c := NewSomeChain(data1)
	r := c.Concate(data2).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestDrop(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeChain(data)
	r := c.Drop(1).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestFilter(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeChain(data)
	filter := func(idx int, some *Some) bool {
		return idx == 0
	}
	r := c.Filter(filter).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{"hello"}}) {
		t.Fatal("mistach")
	}
}

func TestFirst(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeChain(data)
	r := c.First()
	if !reflect.DeepEqual(r, &Some{"hello"}) {
		t.Fatal("mistach")
	}
}

func TestLast(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeChain(data)
	r := c.Last()
	if !reflect.DeepEqual(r, &Some{"world"}) {
		t.Fatal("mistach")
	}
}

func TestMap(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeChain(data)
	mapFn := func(idx int, some *Some) {
		some.A += "_test"
	}
	r := c.Map(mapFn).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{"hello_test"}, &Some{"world_test"}}) {
		t.Fatal("mistach")
	}
}

func TestReduce(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeChain(data)
	reduceFn := func(initial *Some, cur *Some, idx int) *Some {
		return &Some{initial.A + " " + cur.A}
	}
	r := c.Reduce(reduceFn, &Some{"initial"})
	if !reflect.DeepEqual(r, &Some{"initial hello world"}) {
		t.Fatal("mistach")
	}
}
func TestReverse(t *testing.T) {
	data := []*Some{&Some{"world"}, &Some{"hello"}}
	c := NewSomeChain(data)
	r := c.Reverse().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestUnique(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{"world"}, &Some{"hello"}}
	c := NewSomeChain(data)
	r := c.Unique().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}
