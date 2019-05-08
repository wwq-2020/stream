
package commons

import (
	"sort"
	"math/rand"
)

const EmptyInt64 int64 =0

type Int64Chain struct{
	value	[]int64
}

func NewInt64Chain(value []int64) *Int64Chain {
	return &Int64Chain{value:value}
}

func(c *Int64Chain) Concate(given []int64)  *Int64Chain {
	value := make([]int64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int64Chain) Drop(n int)  *Int64Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int64Chain) Filter(fn func(int, int64)bool)  *Int64Chain {
	value := make([]int64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int64Chain) First() int64 {
	if len(c.value) <= 0 {
		return EmptyInt64
	} 
	return c.value[0]
}

func(c *Int64Chain) Last() int64 {
	if len(c.value) <= 0 {
		return EmptyInt64
	} 
	return c.value[len(c.value)-1]
}

func(c *Int64Chain) Map(fn func(int, int64)) *Int64Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int64Chain) Reduce(fn func(int64, int64, int) int64,initial int64) int64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int64Chain) Reverse()  *Int64Chain {
	value := make([]int64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int64Chain) Unique()  *Int64Chain{
	value := make([]int64, 0, len(c.value))
	seen:=make(map[int64]struct{})
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

func(c *Int64Chain) Append(given int64) *Int64Chain {
	c.value=append(c.value,given)
	return c
}

func(c *Int64Chain) Len() int {
	return len(c.value)
}

func(c *Int64Chain) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int64Chain) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int64Chain)  Sort()  *Int64Chain {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Int64Chain) All(fn func(int, int64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int64Chain) Any(fn func(int, int64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int64Chain) Paginate(size int)  [][]int64 {
	var pages  [][]int64
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

func(c *Int64Chain) Pop() int64{
	if len(c.value) <= 0 {
		return EmptyInt64 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int64Chain) Prepend(given int64) *Int64Chain {
	c.value = append([]int64{given},c.value...)
	return c
}

func(c *Int64Chain) Max() int64{
	if len(c.value) <= 0 {
		return EmptyInt64 
	}
	var max int64
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


func(c *Int64Chain) Min() int64{
	if len(c.value) <= 0 {
		return EmptyInt64 
	}
	var min int64
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

func(c *Int64Chain) Random() int64{
	if len(c.value) <= 0 {
		return EmptyInt64 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int64Chain) Shuffle() *Int64Chain {
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

func(c *Int64Chain) Collect() []int64{
	return c.value
}
