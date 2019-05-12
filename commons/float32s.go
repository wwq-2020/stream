
package commons

import (
	"sort"
	"math/rand"
)

type Float32Stream struct{
	value	[]float32
	defaultReturn float32
}

func StreamOfFloat32(value []float32) *Float32Stream {
	return &Float32Stream{value:value,defaultReturn:0.0}
}

func(c *Float32Stream) OrElase(defaultReturn float32)  *Float32Stream {
	c.defaultReturn = defaultReturn
	return c
}


func(c *Float32Stream) Concate(given []float32)  *Float32Stream {
	value := make([]float32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float32Stream) Drop(n int)  *Float32Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float32Stream) Filter(fn func(int, float32)bool)  *Float32Stream {
	value := make([]float32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float32Stream) First() float32 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Float32Stream) Last() float32 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Float32Stream) Map(fn func(int, float32)) *Float32Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float32Stream) Reduce(fn func(float32, float32, int) float32,initial float32) float32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float32Stream) Reverse()  *Float32Stream {
	value := make([]float32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float32Stream) Unique()  *Float32Stream{
	value := make([]float32, 0, len(c.value))
	seen:=make(map[float32]struct{})
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

func(c *Float32Stream) Append(given float32) *Float32Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Float32Stream) Len() int {
	return len(c.value)
}

func(c *Float32Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Float32Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Float32Stream)  Sort()  *Float32Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] < c.value[j]
	})
	return c 
}

func(c *Float32Stream) All(fn func(int, float32)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Float32Stream) Any(fn func(int, float32)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Float32Stream) Paginate(size int)  [][]float32 {
	var pages  [][]float32
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

func(c *Float32Stream) Pop() float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Float32Stream) Prepend(given float32) *Float32Stream {
	c.value = append([]float32{given},c.value...)
	return c
}

func(c *Float32Stream) Max() float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max float32 = c.value[0]
	for _,each := range c.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(c *Float32Stream) Min() float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min float32 = c.value[0]
	for _,each := range c.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *Float32Stream) Random() float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Float32Stream) Shuffle() *Float32Stream {
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

func(c *Float32Stream) Collect() []float32{
	return c.value
}


type Float32PStream struct{
	value	[]*float32
	defaultReturn *float32
}

func PStreamOfFloat32(value []*float32) *Float32PStream {
	return &Float32PStream{value:value,defaultReturn:nil}
}

func(c *Float32PStream) OrElse(defaultReturn *float32)  *Float32PStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *Float32PStream) Concate(given []*float32)  *Float32PStream {
	value := make([]*float32, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float32PStream) Drop(n int)  *Float32PStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float32PStream) Filter(fn func(int, *float32)bool)  *Float32PStream {
	value := make([]*float32, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float32PStream) First() *float32 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Float32PStream) Last() *float32 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Float32PStream) Map(fn func(int, *float32)) *Float32PStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float32PStream) Reduce(fn func(*float32, *float32, int) *float32,initial *float32) *float32   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float32PStream) Reverse()  *Float32PStream {
	value := make([]*float32, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float32PStream) Unique()  *Float32PStream{
	value := make([]*float32, 0, len(c.value))
	seen:=make(map[*float32]struct{})
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

func(c *Float32PStream) Append(given *float32) *Float32PStream {
	c.value=append(c.value,given)
	return c
}

func(c *Float32PStream) Len() int {
	return len(c.value)
}

func(c *Float32PStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Float32PStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Float32PStream)  Sort(less func(*float32,*float32) bool )  *Float32PStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Float32PStream) All(fn func(int, *float32)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Float32PStream) Any(fn func(int, *float32)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Float32PStream) Paginate(size int)  [][]*float32 {
	var pages  [][]*float32
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

func(c *Float32PStream) Pop() *float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Float32PStream) Prepend(given *float32) *Float32PStream {
	c.value = append([]*float32{given},c.value...)
	return c
}

func(c *Float32PStream) Max() *float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *float32 = c.value[0]
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


func(c *Float32PStream) Min() *float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *float32 = c.value[0]
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

func(c *Float32PStream) Random() *float32{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Float32PStream) Shuffle() *Float32PStream {
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

func(c *Float32PStream) Collect() []*float32{
	return c.value
}
