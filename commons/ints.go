
package commons

import (
	"sort"
	"math/rand"
)

const EmptyInt int =0

type IntChain struct{
	value	[]int
}

func NewIntChain(value []int) *IntChain {
	return &IntChain{value:value}
}

func(c *IntChain) Concate(given []int)  *IntChain {
	value := make([]int, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *IntChain) Drop(n int)  *IntChain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *IntChain) Filter(fn func(int, int)bool)  *IntChain {
	value := make([]int, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *IntChain) First() int {
	if len(c.value) <= 0 {
		return EmptyInt
	} 
	return c.value[0]
}

func(c *IntChain) Last() int {
	if len(c.value) <= 0 {
		return EmptyInt
	} 
	return c.value[len(c.value)-1]
}

func(c *IntChain) Map(fn func(int, int)) *IntChain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *IntChain) Reduce(fn func(int, int, int) int,initial int) int   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *IntChain) Reverse()  *IntChain {
	value := make([]int, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *IntChain) Unique()  *IntChain{
	value := make([]int, 0, len(c.value))
	seen:=make(map[int]struct{})
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

func(c *IntChain) Append(given int) *IntChain {
	c.value=append(c.value,given)
	return c
}

func(c *IntChain) Len() int {
	return len(c.value)
}

func(c *IntChain) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *IntChain) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *IntChain)  Sort()  *IntChain {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *IntChain) All(fn func(int, int)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *IntChain) Any(fn func(int, int)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *IntChain) Paginate(size int)  [][]int {
	var pages  [][]int
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

func(c *IntChain) Pop() int{
	if len(c.value) <= 0 {
		return EmptyInt 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *IntChain) Prepend(given int) *IntChain {
	c.value = append([]int{given},c.value...)
	return c
}

func(c *IntChain) Max() int{
	if len(c.value) <= 0 {
		return EmptyInt 
	}
	var max int
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


func(c *IntChain) Min() int{
	if len(c.value) <= 0 {
		return EmptyInt 
	}
	var min int
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

func(c *IntChain) Random() int{
	if len(c.value) <= 0 {
		return EmptyInt 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *IntChain) Shuffle() *IntChain {
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

func(c *IntChain) Collect() []int{
	return c.value
}
