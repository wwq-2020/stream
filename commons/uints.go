
package commons

const EmptyUint uint =0

type UintChain struct{
	value	[]uint
}

func NewUintChain(value []uint) *UintChain {
	return &UintChain{value:value}
}

func(c *UintChain) Concate(given []uint)  *UintChain {
	value := make([]uint, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *UintChain) Drop(n int)  *UintChain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *UintChain) Filter(fn func(int, uint)bool)  *UintChain {
	value := make([]uint, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *UintChain) First() uint {
	if len(c.value) <= 0 {
		return EmptyUint
	} 
	return c.value[0]
}

func(c *UintChain) Last() uint {
	if len(c.value) <= 0 {
		return EmptyUint
	} 
	return c.value[len(c.value)-1]
}

func(c *UintChain) Map(fn func(int, uint)) *UintChain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *UintChain) Reduce(fn func(uint, uint, int) uint,initial uint) uint   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *UintChain) Reverse()  *UintChain {
	value := make([]uint, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *UintChain) Unique()  *UintChain{
	value := make([]uint, 0, len(c.value))
	seen:=make(map[uint]struct{})
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

func(c *UintChain) Collect() []uint{
	return c.value
}
