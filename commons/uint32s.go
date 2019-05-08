
package commons

const EmptyUint32 uint32 =0

type Uint32Chain struct{
	value	[]uint32
}

func NewUint32Chain(value []uint32) *Uint32Chain {
	return &Uint32Chain{value:value}
}

func(c *Uint32Chain) Concate(given []uint32)  *Uint32Chain {
	value := make([]uint32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint32Chain) Drop(n int)  *Uint32Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint32Chain) Filter(fn func(int, uint32)bool)  *Uint32Chain {
	value := make([]uint32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint32Chain) First() uint32 {
	if len(c.value) <= 0 {
		return EmptyUint32
	} 
	return c.value[0]
}

func(c *Uint32Chain) Last() uint32 {
	if len(c.value) <= 0 {
		return EmptyUint32
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint32Chain) Map(fn func(int, uint32)) *Uint32Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint32Chain) Reduce(fn func(uint32, uint32, int) uint32,initial uint32) uint32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint32Chain) Reverse()  *Uint32Chain {
	value := make([]uint32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint32Chain) Unique()  *Uint32Chain{
	value := make([]uint32, 0, len(c.value))
	seen:=make(map[uint32]struct{})
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

func(c *Uint32Chain) Collect() []uint32{
	return c.value
}
