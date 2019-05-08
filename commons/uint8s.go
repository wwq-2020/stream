
package commons

import (
	"sort"
	"math/rand"
)

const EmptyUint8 uint8 =0

type Uint8Chain struct{
	value	[]uint8
}

func NewUint8Chain(value []uint8) *Uint8Chain {
	return &Uint8Chain{value:value}
}

func(c *Uint8Chain) Concate(given []uint8)  *Uint8Chain {
	value := make([]uint8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint8Chain) Drop(n int)  *Uint8Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint8Chain) Filter(fn func(int, uint8)bool)  *Uint8Chain {
	value := make([]uint8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint8Chain) First() uint8 {
	if len(c.value) <= 0 {
		return EmptyUint8
	} 
	return c.value[0]
}

func(c *Uint8Chain) Last() uint8 {
	if len(c.value) <= 0 {
		return EmptyUint8
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint8Chain) Map(fn func(int, uint8)) *Uint8Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint8Chain) Reduce(fn func(uint8, uint8, int) uint8,initial uint8) uint8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint8Chain) Reverse()  *Uint8Chain {
	value := make([]uint8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint8Chain) Unique()  *Uint8Chain{
	value := make([]uint8, 0, len(c.value))
	seen:=make(map[uint8]struct{})
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

func(c *Uint8Chain) Append(given uint8) *Uint8Chain {
	c.value=append(c.value,given)
	return c
}

func(c *Uint8Chain) Len() int {
	return len(c.value)
}

func(c *Uint8Chain) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint8Chain) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint8Chain)  Sort()  *Uint8Chain {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Uint8Chain) All(fn func(int, uint8)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint8Chain) Any(fn func(int, uint8)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint8Chain) Paginate(size int)  [][]uint8 {
	var pages  [][]uint8
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

func(c *Uint8Chain) Pop() uint8{
	if len(c.value) <= 0 {
		return EmptyUint8 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint8Chain) Prepend(given uint8) *Uint8Chain {
	c.value = append([]uint8{given},c.value...)
	return c
}

func(c *Uint8Chain) Max() uint8{
	if len(c.value) <= 0 {
		return EmptyUint8 
	}
	var max uint8
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


func(c *Uint8Chain) Min() uint8{
	if len(c.value) <= 0 {
		return EmptyUint8 
	}
	var min uint8
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

func(c *Uint8Chain) Random() uint8{
	if len(c.value) <= 0 {
		return EmptyUint8 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint8Chain) Shuffle() *Uint8Chain {
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

func(c *Uint8Chain) Collect() []uint8{
	return c.value
}
