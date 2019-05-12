
package commons

import (
	"sort"
	"math/rand"
)

type Int8Stream struct{
	value	[]int8
	defaultReturn int8
}

func StreamOfInt8(value []int8) *Int8Stream {
	return &Int8Stream{value:value,defaultReturn:0}
}

func(c *Int8Stream) OrElase(defaultReturn int8)  *Int8Stream {
	c.defaultReturn = defaultReturn
	return c
}


func(c *Int8Stream) Concate(given []int8)  *Int8Stream {
	value := make([]int8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int8Stream) Drop(n int)  *Int8Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int8Stream) Filter(fn func(int, int8)bool)  *Int8Stream {
	value := make([]int8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int8Stream) First() int8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Int8Stream) Last() int8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Int8Stream) Map(fn func(int, int8)) *Int8Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int8Stream) Reduce(fn func(int8, int8, int) int8,initial int8) int8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int8Stream) Reverse()  *Int8Stream {
	value := make([]int8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int8Stream) Unique()  *Int8Stream{
	value := make([]int8, 0, len(c.value))
	seen:=make(map[int8]struct{})
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

func(c *Int8Stream) Append(given int8) *Int8Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Int8Stream) Len() int {
	return len(c.value)
}

func(c *Int8Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int8Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int8Stream)  Sort()  *Int8Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] < c.value[j]
	})
	return c 
}

func(c *Int8Stream) All(fn func(int, int8)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int8Stream) Any(fn func(int, int8)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int8Stream) Paginate(size int)  [][]int8 {
	var pages  [][]int8
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

func(c *Int8Stream) Pop() int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int8Stream) Prepend(given int8) *Int8Stream {
	c.value = append([]int8{given},c.value...)
	return c
}

func(c *Int8Stream) Max() int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max int8 = c.value[0]
	for _,each := range c.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(c *Int8Stream) Min() int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min int8 = c.value[0]
	for _,each := range c.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *Int8Stream) Random() int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int8Stream) Shuffle() *Int8Stream {
	if len(c.value) <= 0 {
		return c
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

func(c *Int8Stream) Collect() []int8{
	return c.value
}


type Int8PStream struct{
	value	[]*int8
	defaultReturn *int8
}

func PStreamOfInt8(value []*int8) *Int8PStream {
	return &Int8PStream{value:value,defaultReturn:nil}
}

func(c *Int8PStream) OrElse(defaultReturn *int8)  *Int8PStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *Int8PStream) Concate(given []*int8)  *Int8PStream {
	value := make([]*int8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int8PStream) Drop(n int)  *Int8PStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int8PStream) Filter(fn func(int, *int8)bool)  *Int8PStream {
	value := make([]*int8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int8PStream) First() *int8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Int8PStream) Last() *int8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Int8PStream) Map(fn func(int, *int8)) *Int8PStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int8PStream) Reduce(fn func(*int8, *int8, int) *int8,initial *int8) *int8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int8PStream) Reverse()  *Int8PStream {
	value := make([]*int8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int8PStream) Unique()  *Int8PStream{
	value := make([]*int8, 0, len(c.value))
	seen:=make(map[*int8]struct{})
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

func(c *Int8PStream) Append(given *int8) *Int8PStream {
	c.value=append(c.value,given)
	return c
}

func(c *Int8PStream) Len() int {
	return len(c.value)
}

func(c *Int8PStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int8PStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int8PStream)  Sort(less func(*int8,*int8) bool )  *Int8PStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Int8PStream) All(fn func(int, *int8)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int8PStream) Any(fn func(int, *int8)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int8PStream) Paginate(size int)  [][]*int8 {
	var pages  [][]*int8
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

func(c *Int8PStream) Pop() *int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int8PStream) Prepend(given *int8) *Int8PStream {
	c.value = append([]*int8{given},c.value...)
	return c
}

func(c *Int8PStream) Max() *int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *int8 = c.value[0]
	for _,each := range c.value {
		if max == nil{
			max = each
			continue
		}
		if each != nil && *max <= *each {
			max = each
		}
	}
	return max
}


func(c *Int8PStream) Min() *int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *int8 = c.value[0]
	for _,each := range c.value {
		if min == nil{
			min = each
			continue
		}
		if  each != nil && *each  <= *min {
			min = each
		}
	}
	return min
}

func(c *Int8PStream) Random() *int8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int8PStream) Shuffle() *Int8PStream {
	if len(c.value) <= 0 {
		return c
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

func(c *Int8PStream) Collect() []*int8{
	return c.value
}
