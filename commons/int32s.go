
package commons

const EmptyInt32 int32 =0

type Int32Chain struct{
	value	[]int32
}

func NewInt32Chain(value []int32) *Int32Chain {
	return &Int32Chain{value:value}
}

func(c *Int32Chain) Concate(given []int32)  *Int32Chain {
	value := make([]int32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int32Chain) Drop(n int)  *Int32Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int32Chain) Filter(fn func(int, int32)bool)  *Int32Chain {
	value := make([]int32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int32Chain) First() int32 {
	if len(c.value) <= 0 {
		return EmptyInt32
	} 
	return c.value[0]
}

func(c *Int32Chain) Last() int32 {
	if len(c.value) <= 0 {
		return EmptyInt32
	} 
	return c.value[len(c.value)-1]
}

func(c *Int32Chain) Map(fn func(int, int32)) *Int32Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int32Chain) Reduce(fn func(int32, int32, int) int32,initial int32) int32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int32Chain) Reverse()  *Int32Chain {
	value := make([]int32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int32Chain) Unique()  *Int32Chain{
	value := make([]int32, 0, len(c.value))
	seen:=make(map[int32]struct{})
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

func(c *Int32Chain) Collect() []int32{
	return c.value
}
