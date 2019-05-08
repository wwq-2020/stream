
package commons

import (
	"sort"
	"math/rand"
)

const EmptyUint64 uint64 =0

type Uint64Chain struct{
	value	[]uint64
}

func NewUint64Chain(value []uint64) *Uint64Chain {
	return &Uint64Chain{value:value}
}

func(c *Uint64Chain) Concate(given []uint64)  *Uint64Chain {
	value := make([]uint64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint64Chain) Drop(n int)  *Uint64Chain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint64Chain) Filter(fn func(int, uint64)bool)  *Uint64Chain {
	value := make([]uint64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint64Chain) First() uint64 {
	if len(c.value) <= 0 {
		return EmptyUint64
	} 
	return c.value[0]
}

func(c *Uint64Chain) Last() uint64 {
	if len(c.value) <= 0 {
		return EmptyUint64
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint64Chain) Map(fn func(int, uint64)) *Uint64Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint64Chain) Reduce(fn func(uint64, uint64, int) uint64,initial uint64) uint64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint64Chain) Reverse()  *Uint64Chain {
	value := make([]uint64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint64Chain) Unique()  *Uint64Chain{
	value := make([]uint64, 0, len(c.value))
	seen:=make(map[uint64]struct{})
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

func(c *Uint64Chain) Append(given uint64) *Uint64Chain {
	c.value=append(c.value,given)
	return c
}

func(c *Uint64Chain) Len() int {
	return len(c.value)
}

func(c *Uint64Chain) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint64Chain) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint64Chain)  Sort()  *Uint64Chain {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Uint64Chain) All(fn func(int, uint64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint64Chain) Any(fn func(int, uint64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint64Chain) Paginate(size int)  [][]uint64 {
	var pages  [][]uint64
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

func(c *Uint64Chain) Pop() uint64{
	if len(c.value) <= 0 {
		return EmptyUint64 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint64Chain) Prepend(given uint64) *Uint64Chain {
	c.value = append([]uint64{given},c.value...)
	return c
}

func(c *Uint64Chain) Max() uint64{
	if len(c.value) <= 0 {
		return EmptyUint64 
	}
	var max uint64
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


func(c *Uint64Chain) Min() uint64{
	if len(c.value) <= 0 {
		return EmptyUint64 
	}
	var min uint64
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

func(c *Uint64Chain) Random() uint64{
	if len(c.value) <= 0 {
		return EmptyUint64 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint64Chain) Shuffle() *Uint64Chain {
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

func(c *Uint64Chain) Collect() []uint64{
	return c.value
}
