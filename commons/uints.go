
package commons

import (
	"sort"
	"math/rand"
)

const EmptyUint uint =0

type UintStream struct{
	value	[]uint
}

func StreamOfUint(value []uint) *UintStream {
	return &UintStream{value:value}
}

func(c *UintStream) Concate(given []uint)  *UintStream {
	value := make([]uint, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *UintStream) Drop(n int)  *UintStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *UintStream) Filter(fn func(int, uint)bool)  *UintStream {
	value := make([]uint, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *UintStream) First() uint {
	if len(c.value) < 0 {
		return EmptyUint
	} 
	return c.value[0]
}

func(c *UintStream) Last() uint {
	if len(c.value) < 0 {
		return EmptyUint
	} 
	return c.value[len(c.value)-1]
}

func(c *UintStream) Map(fn func(int, uint)) *UintStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *UintStream) Reduce(fn func(uint, uint, int) uint,initial uint) uint   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *UintStream) Reverse()  *UintStream {
	value := make([]uint, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *UintStream) Unique()  *UintStream{
	value := make([]uint, 0, len(c.value))
	seen:=make(map[uint]struct{})
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

func(c *UintStream) Append(given uint) *UintStream {
	c.value=append(c.value,given)
	return c
}

func(c *UintStream) Len() int {
	return len(c.value)
}

func(c *UintStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *UintStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *UintStream)  SortBy(less func(uint,uint) bool )  *UintStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *UintStream) All(fn func(int, uint)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *UintStream) Any(fn func(int, uint)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *UintStream) Paginate(size int)  [][]uint {
	var pages  [][]uint
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

func(c *UintStream) Pop() uint{
	if len(c.value) < 0 {
		return EmptyUint 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *UintStream) Prepend(given uint) *UintStream {
	c.value = append([]uint{given},c.value...)
	return c
}

func(c *UintStream) Max() uint{
	if len(c.value) < 0 {
		return EmptyUint 
	}
	var max uint
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


func(c *UintStream) Min() uint{
	if len(c.value) < 0 {
		return EmptyUint 
	}
	var min uint
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

func(c *UintStream) Random() uint{
	if len(c.value) < 0 {
		return EmptyUint 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *UintStream) Shuffle() *UintStream {
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

func(c *UintStream) Collect() []uint{
	return c.value
}
