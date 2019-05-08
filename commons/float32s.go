
package commons

const EmptyFloat32 float32 =0.0

type Float32Chain struct{
	value	[]float32
}

func NewFloat32Chain(value []float32) *Float32Chain {
	return &Float32Chain{value:value}
}

func(c *Float32Chain) Concate(given []float32)  *Float32Chain {
	value := make([]float32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float32Chain) Drop(n int)  *Float32Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float32Chain) Filter(fn func(int, float32)bool)  *Float32Chain {
	value := make([]float32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float32Chain) First() float32 {
	if len(c.value) <= 0 {
		return EmptyFloat32
	} 
	return c.value[0]
}

func(c *Float32Chain) Last() float32 {
	if len(c.value) <= 0 {
		return EmptyFloat32
	} 
	return c.value[len(c.value)-1]
}

func(c *Float32Chain) Map(fn func(int, float32)) *Float32Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float32Chain) Reduce(fn func(float32, float32, int) float32,initial float32) float32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float32Chain) Reverse()  *Float32Chain {
	value := make([]float32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float32Chain) Unique()  *Float32Chain{
	value := make([]float32, 0, len(c.value))
	seen:=make(map[float32]struct{})
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

func(c *Float32Chain) Collect() []float32{
	return c.value
}
