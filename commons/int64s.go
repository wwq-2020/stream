
package commons

const EmptyInt64 int64 =0

type Int64Chain struct{
	value	[]int64
}

func NewInt64Chain(value []int64) *Int64Chain {
	return &Int64Chain{value:value}
}

func(c *Int64Chain) Concate(given []int64)  *Int64Chain {
	value := make([]int64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int64Chain) Drop(n int)  *Int64Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int64Chain) Filter(fn func(int, int64)bool)  *Int64Chain {
	value := make([]int64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int64Chain) First() int64 {
	if len(c.value) <= 0 {
		return EmptyInt64
	} 
	return c.value[0]
}

func(c *Int64Chain) Last() int64 {
	if len(c.value) <= 0 {
		return EmptyInt64
	} 
	return c.value[len(c.value)-1]
}

func(c *Int64Chain) Map(fn func(int, int64)) *Int64Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int64Chain) Reduce(fn func(int64, int64, int) int64,initial int64) int64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int64Chain) Reverse()  *Int64Chain {
	value := make([]int64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int64Chain) Unique()  *Int64Chain{
	value := make([]int64, 0, len(c.value))
	seen:=make(map[int64]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *Int64Chain) Collect() []int64{
	return c.value
}
