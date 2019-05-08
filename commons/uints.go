
package commons

import (
	"sort"
	"math/rand"
)

const EmptyUint uint =0

type UintCollection struct{
	value	[]uint
}

func NewUintCollection(value []uint) *UintCollection {
	return &UintCollection{value:value}
}

func(c *UintCollection) Concate(given []uint)  *UintCollection {
	value := make([]uint, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *UintCollection) Drop(n int)  *UintCollection {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *UintCollection) Filter(fn func(int, uint)bool)  *UintCollection {
	value := make([]uint, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *UintCollection) First() uint {
	if len(c.value) <= 0 {
		return EmptyUint
	} 
	return c.value[0]
}

func(c *UintCollection) Last() uint {
	if len(c.value) <= 0 {
		return EmptyUint
	} 
	return c.value[len(c.value)-1]
}

func(c *UintCollection) Map(fn func(int, uint)) *UintCollection {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *UintCollection) Reduce(fn func(uint, uint, int) uint,initial uint) uint   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *UintCollection) Reverse()  *UintCollection {
	value := make([]uint, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *UintCollection) Unique()  *UintCollection{
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

func(c *UintCollection) Append(given uint) *UintCollection {
	c.value=append(c.value,given)
	return c
}

func(c *UintCollection) Len() int {
	return len(c.value)
}

func(c *UintCollection) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *UintCollection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *UintCollection)  Sort()  *UintCollection {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *UintCollection) All(fn func(int, uint)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *UintCollection) Any(fn func(int, uint)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *UintCollection) Paginate(size int)  [][]uint {
	var pages  [][]uint
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

func(c *UintCollection) Pop() uint{
	if len(c.value) <= 0 {
		return EmptyUint 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *UintCollection) Prepend(given uint) *UintCollection {
	c.value = append([]uint{given},c.value...)
	return c
}

func(c *UintCollection) Max() uint{
	if len(c.value) <= 0 {
		return EmptyUint 
	}
	var max uint
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


func(c *UintCollection) Min() uint{
	if len(c.value) <= 0 {
		return EmptyUint 
	}
	var min uint
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

func(c *UintCollection) Random() uint{
	if len(c.value) <= 0 {
		return EmptyUint 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *UintCollection) Shuffle() *UintCollection {
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

func(c *UintCollection) Collect() []uint{
	return c.value
}
