
package tests
type SomeChain struct{
	value	[]*Some
}

func NewSomeChain(value []*Some) *SomeChain {
	return &SomeChain{value:value}
}

func(c *SomeChain) Concate(given []*Some)  *SomeChain {
	value := make([]*Some, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *SomeChain) Drop(n int)  *SomeChain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *SomeChain) Filter(fn func(int, *Some)bool)  *SomeChain {
	value := make([]*Some, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *SomeChain) First() *Some {
	if len(c.value) <= 0 {
		return nil
	} 
	return c.value[0]
}

func(c *SomeChain) Last() *Some {
	if len(c.value) <= 0 {
		return nil
	} 
	return c.value[len(c.value)-1]
}

func(c *SomeChain) Map(fn func(int, *Some)) *SomeChain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *SomeChain) Reduce(fn func(*Some, *Some, int) *Some,initial *Some) *Some   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *SomeChain) Reverse()  *SomeChain {
	value := make([]*Some, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *SomeChain) Unique()  *SomeChain{
	value := make([]*Some, 0, len(c.value))
	seen:=make(map[int]struct{})
	for i, outter := range c.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range c.value {
			if i==j {
				continue
			}
			if inner.Compare(outter){
				seen[j]=struct{}{}				
				dup=true
			}
		}
		if dup {
			seen[i]=struct{}{}
		}
		value=append(value,outter)			
	}
	c.value = value
	return c
}

func(c *SomeChain) Collect() []*Some{
	return c.value
}
