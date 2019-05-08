
package commons

import (
	"sort"
	"math/rand"
)

const EmptyFloat64 float64 =0.0

type Float64Collection struct{
	value	[]float64
}

func NewFloat64Collection(value []float64) *Float64Collection {
	return &Float64Collection{value:value}
}

func(c *Float64Collection) Concate(given []float64)  *Float64Collection {
	value := make([]float64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float64Collection) Drop(n int)  *Float64Collection {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float64Collection) Filter(fn func(int, float64)bool)  *Float64Collection {
	value := make([]float64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float64Collection) First() float64 {
	if len(c.value) <= 0 {
		return EmptyFloat64
	} 
	return c.value[0]
}

func(c *Float64Collection) Last() float64 {
	if len(c.value) <= 0 {
		return EmptyFloat64
	} 
	return c.value[len(c.value)-1]
}

func(c *Float64Collection) Map(fn func(int, float64)) *Float64Collection {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float64Collection) Reduce(fn func(float64, float64, int) float64,initial float64) float64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float64Collection) Reverse()  *Float64Collection {
	value := make([]float64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float64Collection) Unique()  *Float64Collection{
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

func(c *Float64Collection) Append(given float64) *Float64Collection {
	c.value=append(c.value,given)
	return c
}

func(c *Float64Collection) Len() int {
	return len(c.value)
}

func(c *Float64Collection) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Float64Collection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Float64Collection)  Sort()  *Float64Collection {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Float64Collection) All(fn func(int, float64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Float64Collection) Any(fn func(int, float64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Float64Collection) Paginate(size int)  [][]float64 {
	var pages  [][]float64
	prev := -1
	for i := range c.value {
		if (i-prev) <= size-1 && i != (len(c.value)-1) {
			continue
		}
		pages=append(pages,c.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(c *Float64Collection) Pop() float64{
	if len(c.value) <= 0 {
		return EmptyFloat64 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Float64Collection) Prepend(given float64) *Float64Collection {
	c.value = append([]float64{given},c.value...)
	return c
}

func(c *Float64Collection) Max() float64{
	if len(c.value) <= 0 {
		return EmptyFloat64 
	}
	var max float64
	for idx,each := range c.value {
		if idx==0{
			max=each
			continue
		}
		if max <= each {
			max = each
		}
	}
	return max
}


func(c *Float64Collection) Min() float64{
	if len(c.value) <= 0 {
		return EmptyFloat64 
	}
	var min float64
	for idx,each := range c.value {
		if idx==0{
			min=each
			continue
		}
		if each  <= min {
			min = each
		}
	}
	return min
}

func(c *Float64Collection) Random() float64{
	if len(c.value) <= 0 {
		return EmptyFloat64 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Float64Collection) Shuffle() *Float64Collection {
	if len(c.value) <= 0 {
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

func(c *Float64Collection) Collect() []float64{
	return c.value
}
