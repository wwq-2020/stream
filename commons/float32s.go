
package commons

import (
	"sort"
	"math/rand"
)

const EmptyFloat32 float32 =0.0

type Float32Stream struct{
	value	[]float32
}

func StreamOfFloat32(value []float32) *Float32Stream {
	return &Float32Stream{value:value}
}

func(c *Float32Stream) Concate(given []float32)  *Float32Stream {
	value := make([]float32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float32Stream) Drop(n int)  *Float32Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float32Stream) Filter(fn func(int, float32)bool)  *Float32Stream {
	value := make([]float32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float32Stream) First() float32 {
	if len(c.value) < 0 {
		return EmptyFloat32
	} 
	return c.value[0]
}

func(c *Float32Stream) Last() float32 {
	if len(c.value) < 0 {
		return EmptyFloat32
	} 
	return c.value[len(c.value)-1]
}

func(c *Float32Stream) Map(fn func(int, float32)) *Float32Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float32Stream) Reduce(fn func(float32, float32, int) float32,initial float32) float32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float32Stream) Reverse()  *Float32Stream {
	value := make([]float32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float32Stream) Unique()  *Float32Stream{
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

func(c *Float32Stream) Append(given float32) *Float32Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Float32Stream) Len() int {
	return len(c.value)
}

func(c *Float32Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Float32Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Float32Stream)  SortBy(less func(float32,float32) bool )  *Float32Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Float32Stream) All(fn func(int, float32)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Float32Stream) Any(fn func(int, float32)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Float32Stream) Paginate(size int)  [][]float32 {
	var pages  [][]float32
	prev := -1
	for i := range c.value {
		if (i-prev) < size-1 && i != (len(c.value)-1) {
			continue
		}
		pages=append(pages,c.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(c *Float32Stream) Pop() float32{
	if len(c.value) < 0 {
		return EmptyFloat32 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Float32Stream) Prepend(given float32) *Float32Stream {
	c.value = append([]float32{given},c.value...)
	return c
}

func(c *Float32Stream) Max() float32{
	if len(c.value) < 0 {
		return EmptyFloat32 
	}
	var max float32
	for idx,each := range c.value {
		if idx==0{
			max=each
			continue
		}
		if max < each {
			max = each
		}
	}
	return max
}


func(c *Float32Stream) Min() float32{
	if len(c.value) < 0 {
		return EmptyFloat32 
	}
	var min float32
	for idx,each := range c.value {
		if idx==0{
			min=each
			continue
		}
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *Float32Stream) Random() float32{
	if len(c.value) < 0 {
		return EmptyFloat32 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Float32Stream) Shuffle() *Float32Stream {
	if len(c.value) < 0 {
		return nil
	}
	indexes := make([]int, len(c.value))
	for i := range c.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(c.value), func(i, j int) {
		c.value[i], c.value[j] = 	c.value[j], c.value[i] 
	})
	
	return c
}

func(c *Float32Stream) Collect() []float32{
	return c.value
}
