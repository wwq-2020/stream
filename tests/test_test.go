package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wwq1988/stream/outter"
)

func TestConcate(t *testing.T) {
	data1 := []*Some{&Some{A: "hello"}}
	data2 := []*Some{&Some{A: "world"}}
	c := SomePSlice(data1)
	r := c.Concat(data2).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{A: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestDrop(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	r := c.Drop(1).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestFilter(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	filter := func(idx int, some *Some) bool {
		return idx == 0
	}
	r := c.Filter(filter).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}}) {
		t.Fatal("mistach")
	}
}

func TestFilterByA(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	filter := func(idx int, val string) bool {
		return val == "hello"
	}
	r := c.FilterByA(filter).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}}) {
		t.Fatal("mistach")
	}
}

func TestFirst(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	first, err := c.First()
	if err != nil {
		t.Fatal("empty")
	}
	if !reflect.DeepEqual(first, &Some{A: "hello"}) {
		t.Fatal("mistach")
	}
}

func TestLast(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	last, err := c.Last()
	if err != nil {
		t.Fatal("empty")
	}
	if !reflect.DeepEqual(last, &Some{A: "world"}) {
		t.Fatal("mistach")
	}
}

func TestMap(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	mapFn := func(idx int, some *Some) *Some {
		return &Some{A: some.A + "_test"}
	}
	r := c.Map(mapFn).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello_test"}, &Some{A: "world_test"}}) {
		t.Fatal("mistach")
	}
}

func TestReduce(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	reduceFn := func(initial *Some, cur *Some, idx int) *Some {
		return &Some{A: initial.A + " " + cur.A}
	}
	r := c.Reduce(reduceFn, &Some{A: "initial"})
	if !reflect.DeepEqual(r, &Some{A: "initial hello world"}) {
		t.Fatal("mistach")
	}
}
func TestReverse(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := SomePSlice(data)
	r := c.Reverse().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{A: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestUnique(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}, &Some{A: "hello"}}
	c := SomePSlice(data)
	r := c.UniqueBy(func(one, another *Some) bool {
		return one.A == another.A && one.B == another.B
	}).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{A: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestAppend(t *testing.T) {
	data := []*Some{&Some{A: "hello"}}
	c := SomePSlice(data)
	r := c.Append(&Some{A: "world"}).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{A: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestLen(t *testing.T) {
	data := []*Some{&Some{A: "hello"}}
	c := SomePSlice(data)
	if r := c.Len(); r != len(data) {
		t.Fatal("mistach")
	}
}

func TestIsEmpty(t *testing.T) {
	var data []*Some
	c := SomePSlice(data)
	if !c.IsEmpty() {
		t.Fatal("mistach")
	}
}

func TestIsNotEmpty(t *testing.T) {
	data := []*Some{&Some{A: "hello"}}
	c := SomePSlice(data)
	if !c.IsNotEmpty() {
		t.Fatal("mistach")
	}
}

func TestSortBy(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}

	c := SomePSlice(data)
	r := c.SortBy(func(one, another *Some) bool {
		return one.A < another.A
	}).Collect()

	if len(r) != len(data) {
		t.Fatal("len mismatch")
	}
	if !reflect.DeepEqual(r[0], data[0]) || !reflect.DeepEqual(r[1], data[1]) {
		fmt.Println(r[0], r[1])
		t.Fatal("mistach")
	}
}

func TestAll(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "world"}}
	c := SomePSlice(data)
	if !c.All(func(i int, some *Some) bool {
		return some.A == "world"
	}) {
		t.Fatal("mistach")
	}
}

func TestAllByA(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "world"}}
	c := SomePSlice(data)
	if !c.AllByA(func(i int, val string) bool {
		return val == "world"
	}) {
		t.Fatal("mistach")
	}
}

func TestAnyByA(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	if !c.AnyByA(func(i int, val string) bool {
		return val == "hello"
	}) {
		t.Fatal("mistach")
	}
}

func TestAny(t *testing.T) {
	data := []*Some{&Some{A: "hello"}, &Some{A: "world"}}
	c := SomePSlice(data)
	if !c.Any(func(i int, some *Some) bool {
		return some.A == "world"
	}) {
		t.Fatal("mistach")
	}
}

func TestPaginate(t *testing.T) {
	data := []*Some{&Some{A: "hello1"}, &Some{A: "world1"}, &Some{A: "hello2"}}
	c := SomePSlice(data)
	pages := c.Paginate(2)
	if len(pages) != 2 {
		t.Fatal("mistach")
	}
	if !reflect.DeepEqual(pages[0], []*Some{&Some{A: "hello1"}, &Some{A: "world1"}}) {
		t.Fatal("mistach")
	}
	if !reflect.DeepEqual(pages[1], []*Some{&Some{A: "hello2"}}) {
		t.Fatal("mistach")
	}
}

func TestPrepend(t *testing.T) {
	data := []*Some{&Some{A: "world"}}
	c := SomePSlice(data)
	r := c.Preappend(&Some{A: "hello"}).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello"}, &Some{A: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestMax(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := SomePSlice(data)
	max, err := c.Max(func(one, another *Some) bool {
		return one.A > another.A
	})
	if err != nil {
		t.Fatal("empty")
	}
	if !reflect.DeepEqual(max, &Some{A: "world"}) {
		t.Fatal("mistach")
	}
}
func TestMin(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := SomePSlice(data)
	min, err := c.Min(func(one, another *Some) bool {
		return one.A < another.A
	})
	if err != nil {
		t.Fatal("empty")
	}
	if !reflect.DeepEqual(min, &Some{A: "hello"}) {
		t.Fatal("mistach")
	}
}

func TestRandom(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := SomePSlice(data)
	random, err := c.Random()
	if err != nil {
		t.Fatal("empty")
	}

	if !reflect.DeepEqual(random, &Some{A: "hello"}) && !reflect.DeepEqual(random, &Some{A: "world"}) {
		t.Fatal("mistach")
	}
}

func TestShuffle(t *testing.T) {
	data := []*Some{&Some{A: "world"}, &Some{A: "hello"}}
	c := SomePSlice(data)
	r := c.Shuffle().Collect()
	if len(r) != 2 {
		t.Fatal("mistach")
	}
	if reflect.DeepEqual(r[0], r[1]) {
		t.Fatal("mistach")
	}
}

func TestSortByA(t *testing.T) {
	data := []*Some{&Some{A: "world", B: "hello"}, &Some{A: "hello", B: "world"}}
	c := SomePSlice(data)
	r := c.SortByA().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello", B: "world"}, &Some{A: "world", B: "hello"}}) {
		t.Fatal("mistach")
	}
}

func TestSortByB(t *testing.T) {
	data := []*Some{&Some{A: "world", B: "hello"}, &Some{A: "hello", B: "world"}}
	c := SomePSlice(data)
	r := c.SortByB().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "world", B: "hello"}, &Some{A: "hello", B: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestSortByC(t *testing.T) {
	data := []*Some{&Some{A: "world", B: "hello", C: &Some{A: "world", B: "hello"}}, &Some{A: "hello", B: "world", C: &Some{A: "hello", B: "world"}}}
	c := SomePSlice(data)
	r := c.SortByC(func(one, another *Some) bool {
		return one.A < another.A
	}).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello", B: "world", C: &Some{A: "hello", B: "world"}}, &Some{A: "world", B: "hello", C: &Some{A: "world", B: "hello"}}}) {
		t.Fatal("mistach")
	}
}

func TestUniqueByA(t *testing.T) {
	data := []*Some{&Some{A: "world", B: "hello"}, &Some{A: "world", B: "world"}}
	c := SomePSlice(data)
	r := c.UniqueByA().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "world", B: "hello"}}) {
		t.Fatal("mistach")
	}
}

func TestUniqueByB(t *testing.T) {
	data := []*Some{&Some{A: "hello", B: "world"}, &Some{A: "world", B: "world"}}
	c := SomePSlice(data)
	r := c.UniqueByB().Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello", B: "world"}}) {
		t.Fatal("mistach")
	}
}

func TestUniqueByC(t *testing.T) {
	data := []*Some{&Some{A: "hello", B: "world", C: &Some{A: "world", B: "hello"}}}
	c := SomePSlice(data)
	r := c.UniqueByC(func(one, another *Some) bool {
		return one.A == another.A && one.B == another.B
	}).Collect()
	if !reflect.DeepEqual(r, []*Some{&Some{A: "hello", B: "world", C: &Some{A: "world", B: "hello"}}}) {
		t.Fatal("mistach")
	}
}

func TestFields(t *testing.T) {
	data := []*Some{&Some{A: "hello", B: "world", D: &outter.Some{A: "world", B: "hello"}}, &Some{A: "hello", B: "world", D: &outter.Some{A: "hello", B: "world"}}}
	c := SomePSlice(data)
	r := c.As()
	if !reflect.DeepEqual(r, []string{"hello", "hello"}) {
		t.Fatal("mistach")
	}
}

func TestFieldStream(t *testing.T) {
	data := []*Some{&Some{A: "hello", B: "world", D: &outter.Some{A: "world", B: "hello"}}, &Some{A: "hello", B: "world", D: &outter.Some{A: "hello", B: "world"}}}
	c := SomePSlice(data)
	first, err := c.DPSlice().First()
	if err != nil {
		t.Fatal("empty")
	}
	if !reflect.DeepEqual(first, &outter.Some{A: "world", B: "hello"}) {
		t.Fatal("mistach")
	}
}
