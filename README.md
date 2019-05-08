# collections

##  quick start
```
gen:

  go install collections
  collections -dir tests
 
 tests should be replaced with your package dir witch contains your struct
 
use:

  data1 := []*tests.Some{&tests.Some{"hello"}}
  data2 := []*tests.Some{&tests.Some{"world"}}
  c := tests.NewSomeChain(data1)
  r := c.Concate(data2).Collect()
  fmt.Println(reflect.DeepEqual(r, []*tests.Some{&tests.Some{A: "hello"}, &tests.Some{"world"}}))

  output should be true
  
more:
  see tests or commons
```
