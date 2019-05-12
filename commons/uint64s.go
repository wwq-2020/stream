
package commons

import (
	"sort"
	"math/rand"
)

type Uint64Stream struct{
	value	[]uint64
	defaultReturn uint64
}

func StreamOfUint64(value []uint64) *Uint64Stream {
	return &Uint64Stream{value:value,defaultReturn:0}
}

func(c *Uint64Stream) OrElase(defaultReturn uint64)  *Uint64Stream {
	c.defaultReturn = defaultReturn
	return c
}


func(c *Uint64Stream) Concate(given []uint64)  *Uint64Stream {
	value := make([]uint64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint64Stream) Drop(n int)  *Uint64Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint64Stream) Filter(fn func(int, uint64)bool)  *Uint64Stream {
	value := make([]uint64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint64Stream) First() uint64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Uint64Stream) Last() uint64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint64Stream) Map(fn func(int, uint64)) *Uint64Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint64Stream) Reduce(fn func(uint64, uint64, int) uint64,initial uint64) uint64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint64Stream) Reverse()  *Uint64Stream {
	value := make([]uint64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint64Stream) Unique()  *Uint64Stream{
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

func(c *Uint64Stream) Append(given uint64) *Uint64Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Uint64Stream) Len() int {
	return len(c.value)
}

func(c *Uint64Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint64Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint64Stream)  Sort()  *Uint64Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] < c.value[j]
	})
	return c 
}

func(c *Uint64Stream) All(fn func(int, uint64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint64Stream) Any(fn func(int, uint64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint64Stream) Paginate(size int)  [][]uint64 {
	var pages  [][]uint64
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

func(c *Uint64Stream) Pop() uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint64Stream) Prepend(given uint64) *Uint64Stream {
	c.value = append([]uint64{given},c.value...)
	return c
}

func(c *Uint64Stream) Max() uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max uint64 = c.value[0]
	for _,each := range c.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(c *Uint64Stream) Min() uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min uint64 = c.value[0]
	for _,each := range c.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *Uint64Stream) Random() uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint64Stream) Shuffle() *Uint64Stream {
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

func(c *Uint64Stream) Collect() []uint64{
	return c.value
}


type Uint64PStream struct{
	value	[]*uint64
	defaultReturn *uint64
}

func PStreamOfUint64(value []*uint64) *Uint64PStream {
	return &Uint64PStream{value:value,defaultReturn:nil}
}

func(c *Uint64PStream) OrElse(defaultReturn *uint64)  *Uint64PStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *Uint64PStream) Concate(given []*uint64)  *Uint64PStream {
	value := make([]*uint64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Uint64PStream) Drop(n int)  *Uint64PStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Uint64PStream) Filter(fn func(int, *uint64)bool)  *Uint64PStream {
	value := make([]*uint64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Uint64PStream) First() *uint64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Uint64PStream) Last() *uint64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Uint64PStream) Map(fn func(int, *uint64)) *Uint64PStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Uint64PStream) Reduce(fn func(*uint64, *uint64, int) *uint64,initial *uint64) *uint64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Uint64PStream) Reverse()  *Uint64PStream {
	value := make([]*uint64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Uint64PStream) Unique()  *Uint64PStream{
	value := make([]*uint64, 0, len(c.value))
	seen:=make(map[*uint64]struct{})
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

func(c *Uint64PStream) Append(given *uint64) *Uint64PStream {
	c.value=append(c.value,given)
	return c
}

func(c *Uint64PStream) Len() int {
	return len(c.value)
}

func(c *Uint64PStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Uint64PStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Uint64PStream)  Sort(less func(*uint64,*uint64) bool )  *Uint64PStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Uint64PStream) All(fn func(int, *uint64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Uint64PStream) Any(fn func(int, *uint64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Uint64PStream) Paginate(size int)  [][]*uint64 {
	var pages  [][]*uint64
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

func(c *Uint64PStream) Pop() *uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Uint64PStream) Prepend(given *uint64) *Uint64PStream {
	c.value = append([]*uint64{given},c.value...)
	return c
}

func(c *Uint64PStream) Max() *uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *uint64 = c.value[0]
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


func(c *Uint64PStream) Min() *uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *uint64 = c.value[0]
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

func(c *Uint64PStream) Random() *uint64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Uint64PStream) Shuffle() *Uint64PStream {
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

func(c *Uint64PStream) Collect() []*uint64{
	return c.value
}
