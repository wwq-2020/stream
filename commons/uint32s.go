
package commons

import (
	"sort"
	"math/rand"
)

const EmptyUint32 uint32 =0

type Uint32Collection struct{
	value	[]uint32
}

func NewUint32Collection(value []uint32) *Uint32Collection {
	return &Uint32Collection{value:value}
}

func(c *Uint32Collection) Concate(given []uint32)  *Uint32Collection {
	value := make([]uint32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint32Collection) Drop(n int)  *Uint32Collection {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint32Collection) Filter(fn func(int, uint32)bool)  *Uint32Collection {
	value := make([]uint32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint32Collection) First() uint32 {
	if len(c.value) <= 0 {
		return EmptyUint32
	} 
	return c.value[0]
}

func(c *Uint32Collection) Last() uint32 {
	if len(c.value) <= 0 {
		return EmptyUint32
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint32Collection) Map(fn func(int, uint32)) *Uint32Collection {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint32Collection) Reduce(fn func(uint32, uint32, int) uint32,initial uint32) uint32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint32Collection) Reverse()  *Uint32Collection {
	value := make([]uint32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint32Collection) Unique()  *Uint32Collection{
	value := make([]uint32, 0, len(c.value))
	seen:=make(map[uint32]struct{})
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

func(c *Uint32Collection) Append(given uint32) *Uint32Collection {
	c.value=append(c.value,given)
	return c
}

func(c *Uint32Collection) Len() int {
	return len(c.value)
}

func(c *Uint32Collection) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint32Collection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint32Collection)  Sort()  *Uint32Collection {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Uint32Collection) All(fn func(int, uint32)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint32Collection) Any(fn func(int, uint32)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint32Collection) Paginate(size int)  [][]uint32 {
	var pages  [][]uint32
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

func(c *Uint32Collection) Pop() uint32{
	if len(c.value) <= 0 {
		return EmptyUint32 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint32Collection) Prepend(given uint32) *Uint32Collection {
	c.value = append([]uint32{given},c.value...)
	return c
}

func(c *Uint32Collection) Max() uint32{
	if len(c.value) <= 0 {
		return EmptyUint32 
	}
	var max uint32
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


func(c *Uint32Collection) Min() uint32{
	if len(c.value) <= 0 {
		return EmptyUint32 
	}
	var min uint32
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

func(c *Uint32Collection) Random() uint32{
	if len(c.value) <= 0 {
		return EmptyUint32 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint32Collection) Shuffle() *Uint32Collection {
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

func(c *Uint32Collection) Collect() []uint32{
	return c.value
}
