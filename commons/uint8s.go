
package commons

const EmptyUint8 uint8 =0

type Uint8Chain struct{
	value	[]uint8
}

func NewUint8Chain(value []uint8) *Uint8Chain {
	return &Uint8Chain{value:value}
}

func(c *Uint8Chain) Concate(given []uint8)  *Uint8Chain {
	value := make([]uint8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint8Chain) Drop(n int)  *Uint8Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint8Chain) Filter(fn func(int, uint8)bool)  *Uint8Chain {
	value := make([]uint8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint8Chain) First() uint8 {
	if len(c.value) <= 0 {
		return EmptyUint8
	} 
	return c.value[0]
}

func(c *Uint8Chain) Last() uint8 {
	if len(c.value) <= 0 {
		return EmptyUint8
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint8Chain) Map(fn func(int, uint8)) *Uint8Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint8Chain) Reduce(fn func(uint8, uint8, int) uint8,initial uint8) uint8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint8Chain) Reverse()  *Uint8Chain {
	value := make([]uint8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint8Chain) Unique()  *Uint8Chain{
	value := make([]uint8, 0, len(c.value))
	seen:=make(map[uint8]struct{})
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

func(c *Uint8Chain) Collect() []uint8{
	return c.value
}
