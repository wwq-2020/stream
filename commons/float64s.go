
package commons

const EmptyFloat64 float64 =0.0

type Float64Chain struct{
	value	[]float64
}

func NewFloat64Chain(value []float64) *Float64Chain {
	return &Float64Chain{value:value}
}

func(c *Float64Chain) Concate(given []float64)  *Float64Chain {
	value := make([]float64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float64Chain) Drop(n int)  *Float64Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float64Chain) Filter(fn func(int, float64)bool)  *Float64Chain {
	value := make([]float64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float64Chain) First() float64 {
	if len(c.value) <= 0 {
		return EmptyFloat64
	} 
	return c.value[0]
}

func(c *Float64Chain) Last() float64 {
	if len(c.value) <= 0 {
		return EmptyFloat64
	} 
	return c.value[len(c.value)-1]
}

func(c *Float64Chain) Map(fn func(int, float64)) *Float64Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float64Chain) Reduce(fn func(float64, float64, int) float64,initial float64) float64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float64Chain) Reverse()  *Float64Chain {
	value := make([]float64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float64Chain) Unique()  *Float64Chain{
	value := make([]float64, 0, len(c.value))
	seen:=make(map[float64]struct{})
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

func(c *Float64Chain) Collect() []float64{
	return c.value
}
