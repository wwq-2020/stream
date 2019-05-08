
package commons

import (
	"sort"
	"math/rand"
)

const EmptyUint8 uint8 =0

type Uint8Collection struct{
	value	[]uint8
}

func NewUint8Collection(value []uint8) *Uint8Collection {
	return &Uint8Collection{value:value}
}

func(c *Uint8Collection) Concate(given []uint8)  *Uint8Collection {
	value := make([]uint8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint8Collection) Drop(n int)  *Uint8Collection {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint8Collection) Filter(fn func(int, uint8)bool)  *Uint8Collection {
	value := make([]uint8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint8Collection) First() uint8 {
	if len(c.value) <= 0 {
		return EmptyUint8
	} 
	return c.value[0]
}

func(c *Uint8Collection) Last() uint8 {
	if len(c.value) <= 0 {
		return EmptyUint8
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint8Collection) Map(fn func(int, uint8)) *Uint8Collection {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint8Collection) Reduce(fn func(uint8, uint8, int) uint8,initial uint8) uint8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint8Collection) Reverse()  *Uint8Collection {
	value := make([]uint8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint8Collection) Unique()  *Uint8Collection{
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

func(c *Uint8Collection) Append(given uint8) *Uint8Collection {
	c.value=append(c.value,given)
	return c
}

func(c *Uint8Collection) Len() int {
	return len(c.value)
}

func(c *Uint8Collection) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint8Collection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint8Collection)  Sort()  *Uint8Collection {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *Uint8Collection) All(fn func(int, uint8)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint8Collection) Any(fn func(int, uint8)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint8Collection) Paginate(size int)  [][]uint8 {
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

func(c *Uint8Collection) Pop() uint8{
	if len(c.value) <= 0 {
		return EmptyUint8 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint8Collection) Prepend(given uint8) *Uint8Collection {
	c.value = append([]uint8{given},c.value...)
	return c
}

func(c *Uint8Collection) Max() uint8{
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


func(c *Uint8Collection) Min() uint8{
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

func(c *Uint8Collection) Random() uint8{
	if len(c.value) <= 0 {
		return EmptyUint8 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint8Collection) Shuffle() *Uint8Collection {
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

func(c *Uint8Collection) Collect() []uint8{
	return c.value
}
