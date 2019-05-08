package tests

import (
	"reflect"
	"testing"
)

func TestConcate(t *testing.T) {
	data1 := []*Some{&Some{"hello"}}
	data2 := []*Some{&Some{"world"}}
	c := NewSomeCollection(data1)
	r := c.Concate(data2).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestDrop(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeCollection(data)
	r := c.Drop(1).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestFilter(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeCollection(data)
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
	c := NewSomeCollection(data)
	r := c.First()
	if !reflect.DeepEqual(r, &Some{"hello"}) {
		t.Fatal("mistach")
	}
}

func TestLast(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeCollection(data)
	r := c.Last()
	if !reflect.DeepEqual(r, &Some{"world"}) {
		t.Fatal("mistach")
	}
}

func TestMap(t *testing.T) {
	data := []*Some{&Some{"hello"}, &Some{"world"}}
	c := NewSomeCollection(data)
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
	c := NewSomeCollection(data)
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
	c := NewSomeCollection(data)
	r := c.Reverse().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestUnique(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{"world"}, &Some{"hello"}}
	c := NewSomeCollection(data)
	r := c.Unique().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestAppend(t *testing.T) {
	data := []*Some{&Some{A: "hello"}}
	c := NewSomeCollection(data)
	r := c.Append(&Some{A: "world"}).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestLen(t *testing.T) {
	data := []*Some{&Some{A: "hello"}}
	c := NewSomeCollection(data)
	if r := c.Len(); r != len(data) {
		t.Fatal("mistach")
	}
}

func TestIsEmpty(t *testing.T) {
	var data []*Some
	c := NewSomeCollection(data)
	if !c.IsEmpty() {
		t.Fatal("mistach")
	}
}

func TestIsNotEmpty(t *testing.T) {
	data := []*Some{&Some{A: "hello"}}
	c := NewSomeCollection(data)
	if !c.IsNotEmpty() {
		t.Fatal("mistach")
	}
}

func TestSort(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := NewSomeCollection(data)
	r := c.Sort().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestAll(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "world"}}
	c := NewSomeCollection(data)
	if !c.All(func(i int, some *Some) bool {
		return some.A == "world"
	}) {
		t.Fatal("mistach")
	}
}

func TestAny(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := NewSomeCollection(data)
	if !c.Any(func(i int, some *Some) bool {
		return some.A == "world"
	}) {
		t.Fatal("mistach")
	}
}

func TestPaginate(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := NewSomeCollection(data)
	pages := c.Paginate(1)
	if len(pages) != 2 {
		t.Fatal("mistach")
	}
	if !reflect.DeepEqual(pages[0], []*Some{&Some{A: "hello"}}) {
		t.Fatal("mistach")
	}
	if !reflect.DeepEqual(pages[1], []*Some{&Some{A: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestPop(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := NewSomeCollection(data)
	r := c.Pop()
	if !reflect.DeepEqual(r, &Some{A: "world"}) {
		t.Fatal("mistach")
	}
}

func TestPrepend(t *testing.T) {
	data := []*Some{&Some{A: "world"}}
	c := NewSomeCollection(data)
	r := c.Prepend(&Some{A: "hello"}).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{"world"}}) {
		t.Fatal("mistach")
	}
}

func TestMax(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := NewSomeCollection(data)
	r := c.Max()
	if !reflect.DeepEqual(r, &Some{"world"}) {
		t.Fatal("mistach")
	}
}
func TestMin(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := NewSomeCollection(data)
	r := c.Min()
	if !reflect.DeepEqual(r, &Some{"hello"}) {
		t.Fatal("mistach")
	}
}

func TestRandom(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := NewSomeCollection(data)
	if r := c.Random(); r == nil {
		t.Fatal("mistach")
	}
}

func TestShuffle(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := NewSomeCollection(data)
	r := c.Shuffle().Collect()
	if len(r) != 2 {
		t.Fatal("mistach")
	}
	if reflect.DeepEqual(r[0], r[1]) {
		t.Fatal("mistach")
	}
}
