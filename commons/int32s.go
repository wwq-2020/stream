
package commons

import (
	"sort"
	"math/rand"
)

const EmptyInt32 int32 =0

type Int32Stream struct{
	value	[]int32
}

func StreamOfInt32(value []int32) *Int32Stream {
	return &Int32Stream{value:value}
}

func(c *Int32Stream) Concate(given []int32)  *Int32Stream {
	value := make([]int32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int32Stream) Drop(n int)  *Int32Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int32Stream) Filter(fn func(int, int32)bool)  *Int32Stream {
	value := make([]int32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int32Stream) First() int32 {
	if len(c.value) < 0 {
		return EmptyInt32
	} 
	return c.value[0]
}

func(c *Int32Stream) Last() int32 {
	if len(c.value) < 0 {
		return EmptyInt32
	} 
	return c.value[len(c.value)-1]
}

func(c *Int32Stream) Map(fn func(int, int32)) *Int32Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int32Stream) Reduce(fn func(int32, int32, int) int32,initial int32) int32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int32Stream) Reverse()  *Int32Stream {
	value := make([]int32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int32Stream) Unique()  *Int32Stream{
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

func(c *Int32Stream) Append(given int32) *Int32Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Int32Stream) Len() int {
	return len(c.value)
}

func(c *Int32Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int32Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int32Stream)  SortBy(less func(int32,int32) bool )  *Int32Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Int32Stream) All(fn func(int, int32)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int32Stream) Any(fn func(int, int32)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int32Stream) Paginate(size int)  [][]int32 {
	var pages  [][]int32
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

func(c *Int32Stream) Pop() int32{
	if len(c.value) < 0 {
		return EmptyInt32 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int32Stream) Prepend(given int32) *Int32Stream {
	c.value = append([]int32{given},c.value...)
	return c
}

func(c *Int32Stream) Max() int32{
	if len(c.value) < 0 {
		return EmptyInt32 
	}
	var max int32
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


func(c *Int32Stream) Min() int32{
	if len(c.value) < 0 {
		return EmptyInt32 
	}
	var min int32
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

func(c *Int32Stream) Random() int32{
	if len(c.value) < 0 {
		return EmptyInt32 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int32Stream) Shuffle() *Int32Stream {
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

func(c *Int32Stream) Collect() []int32{
	return c.value
}
