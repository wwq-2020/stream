
package commons

import (
	"sort"
	"math/rand"
)

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

func(c *Float64Chain) Append(given float64) *Float64Chain {
	c.value=append(c.value,given)
	return c
}

func(c *Float64Chain) Len() int {
	return len(c.value)
}

func(c *Float64Chain) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Float64Chain) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Float64Chain)  Sort()  *Float64Chain {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Float64Chain) All(fn func(int, float64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Float64Chain) Any(fn func(int, float64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Float64Chain) Paginate(size int)  [][]float64 {
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

func(c *Float64Chain) Pop() float64{
	if len(c.value) <= 0 {
		return EmptyFloat64 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Float64Chain) Prepend(given float64) *Float64Chain {
	c.value = append([]float64{given},c.value...)
	return c
}

func(c *Float64Chain) Max() float64{
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


func(c *Float64Chain) Min() float64{
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

func(c *Float64Chain) Random() float64{
	if len(c.value) <= 0 {
		return EmptyFloat64 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Float64Chain) Shuffle() *Float64Chain {
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

func(c *Float64Chain) Collect() []float64{
	return c.value
}
