
package commons

const EmptyInt int =0

type IntChain struct{
	value	[]int
}

func NewIntChain(value []int) *IntChain {
	return &IntChain{value:value}
}

func(c *IntChain) Concate(given []int)  *IntChain {
	value := make([]int, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *IntChain) Drop(n int)  *IntChain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *IntChain) Filter(fn func(int, int)bool)  *IntChain {
	value := make([]int, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *IntChain) First() int {
	if len(c.value) <= 0 {
		return EmptyInt
	} 
	return c.value[0]
}

func(c *IntChain) Last() int {
	if len(c.value) <= 0 {
		return EmptyInt
	} 
	return c.value[len(c.value)-1]
}

func(c *IntChain) Map(fn func(int, int)) *IntChain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *IntChain) Reduce(fn func(int, int, int) int,initial int) int   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *IntChain) Reverse()  *IntChain {
	value := make([]int, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *IntChain) Unique()  *IntChain{
	value := make([]int, 0, len(c.value))
	seen:=make(map[int]struct{})
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

func(c *IntChain) Collect() []int{
	return c.value
}
