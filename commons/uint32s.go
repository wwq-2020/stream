
package commons

import (
	"sort"
	"math/rand"
)

const EmptyUint32 uint32 =0

type Uint32Stream struct{
	value	[]uint32
}

func StreamOfUint32(value []uint32) *Uint32Stream {
	return &Uint32Stream{value:value}
}

func(c *Uint32Stream) Concate(given []uint32)  *Uint32Stream {
	value := make([]uint32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint32Stream) Drop(n int)  *Uint32Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint32Stream) Filter(fn func(int, uint32)bool)  *Uint32Stream {
	value := make([]uint32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint32Stream) First() uint32 {
	if len(c.value) < 0 {
		return EmptyUint32
	} 
	return c.value[0]
}

func(c *Uint32Stream) Last() uint32 {
	if len(c.value) < 0 {
		return EmptyUint32
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint32Stream) Map(fn func(int, uint32)) *Uint32Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint32Stream) Reduce(fn func(uint32, uint32, int) uint32,initial uint32) uint32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint32Stream) Reverse()  *Uint32Stream {
	value := make([]uint32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint32Stream) Unique()  *Uint32Stream{
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

func(c *Uint32Stream) Append(given uint32) *Uint32Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Uint32Stream) Len() int {
	return len(c.value)
}

func(c *Uint32Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint32Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint32Stream)  SortBy(less func(uint32,uint32) bool )  *Uint32Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Uint32Stream) All(fn func(int, uint32)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint32Stream) Any(fn func(int, uint32)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint32Stream) Paginate(size int)  [][]uint32 {
	var pages  [][]uint32
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

func(c *Uint32Stream) Pop() uint32{
	if len(c.value) < 0 {
		return EmptyUint32 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint32Stream) Prepend(given uint32) *Uint32Stream {
	c.value = append([]uint32{given},c.value...)
	return c
}

func(c *Uint32Stream) Max() uint32{
	if len(c.value) < 0 {
		return EmptyUint32 
	}
	var max uint32
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


func(c *Uint32Stream) Min() uint32{
	if len(c.value) < 0 {
		return EmptyUint32 
	}
	var min uint32
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

func(c *Uint32Stream) Random() uint32{
	if len(c.value) < 0 {
		return EmptyUint32 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint32Stream) Shuffle() *Uint32Stream {
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

func(c *Uint32Stream) Collect() []uint32{
	return c.value
}
