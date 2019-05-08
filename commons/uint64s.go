
package commons

const EmptyUint64 uint64 =0

type Uint64Chain struct{
	value	[]uint64
}

func NewUint64Chain(value []uint64) *Uint64Chain {
	return &Uint64Chain{value:value}
}

func(c *Uint64Chain) Concate(given []uint64)  *Uint64Chain {
	value := make([]uint64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint64Chain) Drop(n int)  *Uint64Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint64Chain) Filter(fn func(int, uint64)bool)  *Uint64Chain {
	value := make([]uint64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint64Chain) First() uint64 {
	if len(c.value) <= 0 {
		return EmptyUint64
	} 
	return c.value[0]
}

func(c *Uint64Chain) Last() uint64 {
	if len(c.value) <= 0 {
		return EmptyUint64
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint64Chain) Map(fn func(int, uint64)) *Uint64Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint64Chain) Reduce(fn func(uint64, uint64, int) uint64,initial uint64) uint64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint64Chain) Reverse()  *Uint64Chain {
	value := make([]uint64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint64Chain) Unique()  *Uint64Chain{
	value := make([]uint64, 0, len(c.value))
	seen:=make(map[uint64]struct{})
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

func(c *Uint64Chain) Collect() []uint64{
	return c.value
}
