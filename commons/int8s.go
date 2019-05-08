
package commons

const EmptyInt8 int8 =0

type Int8Chain struct{
	value	[]int8
}

func NewInt8Chain(value []int8) *Int8Chain {
	return &Int8Chain{value:value}
}

func(c *Int8Chain) Concate(given []int8)  *Int8Chain {
	value := make([]int8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int8Chain) Drop(n int)  *Int8Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int8Chain) Filter(fn func(int, int8)bool)  *Int8Chain {
	value := make([]int8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int8Chain) First() int8 {
	if len(c.value) <= 0 {
		return EmptyInt8
	} 
	return c.value[0]
}

func(c *Int8Chain) Last() int8 {
	if len(c.value) <= 0 {
		return EmptyInt8
	} 
	return c.value[len(c.value)-1]
}

func(c *Int8Chain) Map(fn func(int, int8)) *Int8Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int8Chain) Reduce(fn func(int8, int8, int) int8,initial int8) int8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int8Chain) Reverse()  *Int8Chain {
	value := make([]int8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int8Chain) Unique()  *Int8Chain{
	value := make([]int8, 0, len(c.value))
	seen:=make(map[int8]struct{})
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

func(c *Int8Chain) Collect() []int8{
	return c.value
}
