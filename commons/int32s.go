
package commons

import (
	"sort"
	"math/rand"
)

const EmptyInt32 int32 =0

type Int32Chain struct{
	value	[]int32
}

func NewInt32Chain(value []int32) *Int32Chain {
	return &Int32Chain{value:value}
}

func(c *Int32Chain) Concate(given []int32)  *Int32Chain {
	value := make([]int32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int32Chain) Drop(n int)  *Int32Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int32Chain) Filter(fn func(int, int32)bool)  *Int32Chain {
	value := make([]int32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int32Chain) First() int32 {
	if len(c.value) <= 0 {
		return EmptyInt32
	} 
	return c.value[0]
}

func(c *Int32Chain) Last() int32 {
	if len(c.value) <= 0 {
		return EmptyInt32
	} 
	return c.value[len(c.value)-1]
}

func(c *Int32Chain) Map(fn func(int, int32)) *Int32Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int32Chain) Reduce(fn func(int32, int32, int) int32,initial int32) int32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int32Chain) Reverse()  *Int32Chain {
	value := make([]int32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int32Chain) Unique()  *Int32Chain{
	value := make([]int32, 0, len(c.value))
	seen:=make(map[int32]struct{})
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

func(c *Int32Chain) Append(given int32) *Int32Chain {
	c.value=append(c.value,given)
	return c
}

func(c *Int32Chain) Len() int {
	return len(c.value)
}

func(c *Int32Chain) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int32Chain) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int32Chain)  Sort()  *Int32Chain {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Int32Chain) All(fn func(int, int32)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int32Chain) Any(fn func(int, int32)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int32Chain) Paginate(size int)  [][]int32 {
	var pages  [][]int32
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

func(c *Int32Chain) Pop() int32{
	if len(c.value) <= 0 {
		return EmptyInt32 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int32Chain) Prepend(given int32) *Int32Chain {
	c.value = append([]int32{given},c.value...)
	return c
}

func(c *Int32Chain) Max() int32{
	if len(c.value) <= 0 {
		return EmptyInt32 
	}
	var max int32
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


func(c *Int32Chain) Min() int32{
	if len(c.value) <= 0 {
		return EmptyInt32 
	}
	var min int32
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

func(c *Int32Chain) Random() int32{
	if len(c.value) <= 0 {
		return EmptyInt32 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int32Chain) Shuffle() *Int32Chain {
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

func(c *Int32Chain) Collect() []int32{
	return c.value
}
