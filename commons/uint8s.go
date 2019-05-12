
package commons

import (
	"sort"
	"math/rand"
)

type Uint8Stream struct{
	value	[]uint8
	defaultReturn uint8
}

func StreamOfUint8(value []uint8) *Uint8Stream {
	return &Uint8Stream{value:value,defaultReturn:0}
}

func(c *Uint8Stream) OrElase(defaultReturn uint8)  *Uint8Stream {
	c.defaultReturn = defaultReturn
	return c
}


func(c *Uint8Stream) Concate(given []uint8)  *Uint8Stream {
	value := make([]uint8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint8Stream) Drop(n int)  *Uint8Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint8Stream) Filter(fn func(int, uint8)bool)  *Uint8Stream {
	value := make([]uint8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint8Stream) First() uint8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Uint8Stream) Last() uint8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint8Stream) Map(fn func(int, uint8)) *Uint8Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint8Stream) Reduce(fn func(uint8, uint8, int) uint8,initial uint8) uint8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint8Stream) Reverse()  *Uint8Stream {
	value := make([]uint8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint8Stream) Unique()  *Uint8Stream{
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

func(c *Uint8Stream) Append(given uint8) *Uint8Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Uint8Stream) Len() int {
	return len(c.value)
}

func(c *Uint8Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint8Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint8Stream)  Sort()  *Uint8Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] < c.value[j]
	})
	return c 
}

func(c *Uint8Stream) All(fn func(int, uint8)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint8Stream) Any(fn func(int, uint8)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint8Stream) Paginate(size int)  [][]uint8 {
	var pages  [][]uint8
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

func(c *Uint8Stream) Pop() uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint8Stream) Prepend(given uint8) *Uint8Stream {
	c.value = append([]uint8{given},c.value...)
	return c
}

func(c *Uint8Stream) Max() uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max uint8 = c.value[0]
	for _,each := range c.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(c *Uint8Stream) Min() uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min uint8 = c.value[0]
	for _,each := range c.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *Uint8Stream) Random() uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint8Stream) Shuffle() *Uint8Stream {
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

func(c *Uint8Stream) Collect() []uint8{
	return c.value
}


type Uint8PStream struct{
	value	[]*uint8
	defaultReturn *uint8
}

func PStreamOfUint8(value []*uint8) *Uint8PStream {
	return &Uint8PStream{value:value,defaultReturn:nil}
}

func(c *Uint8PStream) OrElse(defaultReturn *uint8)  *Uint8PStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *Uint8PStream) Concate(given []*uint8)  *Uint8PStream {
	value := make([]*uint8, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint8PStream) Drop(n int)  *Uint8PStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint8PStream) Filter(fn func(int, *uint8)bool)  *Uint8PStream {
	value := make([]*uint8, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint8PStream) First() *uint8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Uint8PStream) Last() *uint8 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint8PStream) Map(fn func(int, *uint8)) *Uint8PStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint8PStream) Reduce(fn func(*uint8, *uint8, int) *uint8,initial *uint8) *uint8   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint8PStream) Reverse()  *Uint8PStream {
	value := make([]*uint8, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint8PStream) Unique()  *Uint8PStream{
	value := make([]*uint8, 0, len(c.value))
	seen:=make(map[*uint8]struct{})
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

func(c *Uint8PStream) Append(given *uint8) *Uint8PStream {
	c.value=append(c.value,given)
	return c
}

func(c *Uint8PStream) Len() int {
	return len(c.value)
}

func(c *Uint8PStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint8PStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint8PStream)  Sort(less func(*uint8,*uint8) bool )  *Uint8PStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Uint8PStream) All(fn func(int, *uint8)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint8PStream) Any(fn func(int, *uint8)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint8PStream) Paginate(size int)  [][]*uint8 {
	var pages  [][]*uint8
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

func(c *Uint8PStream) Pop() *uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint8PStream) Prepend(given *uint8) *Uint8PStream {
	c.value = append([]*uint8{given},c.value...)
	return c
}

func(c *Uint8PStream) Max() *uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *uint8 = c.value[0]
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


func(c *Uint8PStream) Min() *uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *uint8 = c.value[0]
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

func(c *Uint8PStream) Random() *uint8{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint8PStream) Shuffle() *Uint8PStream {
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

func(c *Uint8PStream) Collect() []*uint8{
	return c.value
}
