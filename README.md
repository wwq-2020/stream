# stream

##  quick start
```
gen:

  go install collections
  collections -dir tests
 
 tests should be replaced with your package dir witch contains your struct
 
use:

  data1 := []*tests.Some{&tests.Some{"hello"}}
  data2 := []*tests.Some{&tests.Some{"world"}}
  s := tests.StreamOfSome(data1)
  r := s.Concate(data2).Collect()
  fmt.Println(reflect.DeepEqual(r, []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{"world"}}))

  output should be true
  
more:
  see tests or commons
```


## doc
support:
  Random pick a random element from slice 
```
    	data := []*tests.Some{&tests.Some{A: "world"}, &tests.Some{A: "hello"}}
	    s := PSomeSlice(data)
	    r := c.Random()
```
  Shuffle shuffle slice
```
      	    data := []*tests.Some{&tests.Some{A: "world"}, &tests.Some{A: "hello"}}
	    s := PSomeSlice(data)
	    r := c.Shuffle().Collect()
```
  Min pick the min element from slice
```
	    data := []*tests.Some{&tests.Some{A: "world"}, &tests.Some{A: "hello"}}
	    s := PSomeSlice(data)
	    r := c.Min()

```
  Max pick the max element from slice
```
            data := []*tests.Some{&tests.Some{A: "world"}, &tests.Some{A: "hello"}}
	    s := PSomeSlice(data)
	    r := c.Max()
```
  Prepend add the element to the head of the slice
```
	    data := []*tests.Some{&tests.Some{A: "world"}}
	    c := PSomeSlice(data)
    	    r := c.Prepend(&tests.Some{A: "hello"}).Collect()
```
  Pop pick a element from the end of the slice 
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      r := c.Pop()
```
  Paginate split elements to pages by given size
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      pages := c.Paginate(1)
```
  Any if any element match the condition
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      r := c.Any(func(i int, some *tests.Some) bool {
        return some.A == "world"
      })

```
  All if all element match the condition
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      r := c.All(func(i int, some *tests.Some) bool {
        return some.A == "world"
      })
```
  SortBy sort by given compare func
```
      data := []*tests.Some{&tests.Some{A: "world", B: "hello"}, &tests.Some{A: "hello", B: "world"}}
      c := PSomeSlice(data)
      r := c.SortByA(func(one, another string) bool {
        return one < another
      }).Collect()
```
  IsNotEmpty check  not empty
```
      data := []*tests.Some{&tests.Some{A: "hello"}}
      c := PSomeSlice(data)
      r := c.IsNotEmpty() 
``` 
  IsEmpty check empty
```
      data := []*tests.Some{&tests.Some{A: "hello"}}
      c := PSomeSlice(data)
      r := c.IsEmpty() 
``` 
  Len get cur slice len
```
      data := []*tests.Some{&tests.Some{A: "hello"}}
      c := PSomeSlice(data)
      r := c.Len() 
```  
  Append add a element to the end of slice
```
      data := []*tests.Some{&tests.Some{A: "hello"}}
      c := PSomeSlice(data)
      r := c.Append(&tests.Some{A: "world"}) 
```
  UniqueBy deduplication the elements by git unique func
```
      data := []*tests.Some{&tests.Some{A: "world", B: "hello"}, &tests.Some{A: "world", B: "world"}}
      c := PSomeSlice(data)
      r := c.UniqueBy(func(one, another *tests.Some) bool {
        return one.A == another.A
      }).Collect()
```  
  Reverse reverse the slice 
```
      data := []*tests.Some{&tests.Some{A: "world"}, &tests.Some{A: "hello"}}
      c := PSomeSlice(data)
      r := c.Reverse().Collect()
```  
  Reduce iter the slice and call the given func with cur element and a initial element or previous result
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      reduceFn := func(initial *tests.Some, cur *tests.Some, idx int) *tests.Some {
        return &tests.Some{A: initial.A + " " + cur.A}
      }
      r := c.Reduce(reduceFn, &tests.Some{A: "initial"})
```  
  Map iter the slice and call the given func with cur element
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      mapFn := func(idx int, some *tests.Some) {
        some.A += "_test"
      }
      r := c.Map(mapFn).Collect()
```  
  Last the last element of the slice 
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      r := c.Last()
``` 
  First the first element of the slice 
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      r := c.Last()
```  
  Filter filter the slice element with  the given func 
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      filter := func(idx int, some *tests.Some) bool {
        return idx == 0
      }
      r := c.Filter(filter).Collect()
```
  Drop drop the head n elements
```
      data := []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{A: "world"}}
      c := PSomeSlice(data)
      r := c.Drop(1).Collect()
```
  Concate concate two slice
```
      data1 := []*tests.Some{&tests.Some{A: "hello"}}
      data2 := []*tests.Some{&tests.Some{A: "world"}}
      c := PSomeSlice(data1)
      r := c.Concate(data2).Collect()
```  
  OrElse set the defaultreturn for slice empty
```
      var data []*tests.Some
      defaultReturn:=&tests.Some{A: "hello"}
      c := PSomeSlice(data)
      r := c.OrElse(defaultReturn).First()
```
SortByField(such as SortByA)

```
      data := []*tests.Some{&tests.Some{A: "world", B: "hello"}, &tests.Some{A: "hello", B: "world"}}
      c := PSomeSlice(data)
      r := c.SortByA(func(one, another string) bool {
        return one < another
      }).Collect()

```
Fields(such as As)
```
      data := []*tests.Some{&tests.Some{A: "hello", B: "world", D: &outter.Some{A: "world", B: "hello"}}, &tests.Some{A: "hello", B: "world", D: &outter.Some{A: "hello", B: "world"}}}
      c := PSomeSlice(data)
      r := c.As()
```
FieldStream(such as DPStream)
```
      data := []*tests.Some{&tests.Some{A: "hello", B: "world", D: &outter.Some{A: "world", B: "hello"}}, &tests.Some{A: "hello", B: "world", D: &outter.Some{A: "hello", B: "world"}}}
      c := PSomeSlice(data)
      r := c.DPStream().First()
```
