
package commons

import (
	"sort"
	"math/rand"
)

type Float64Stream struct{
	value	[]float64
	defaultReturn float64
}

func StreamOfFloat64(value []float64) *Float64Stream {
	return &Float64Stream{value:value,defaultReturn:0.0}
}

func(c *Float64Stream) OrElase(defaultReturn float64)  *Float64Stream {
	c.defaultReturn = defaultReturn
	return c
}


func(c *Float64Stream) Concate(given []float64)  *Float64Stream {
	value := make([]float64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float64Stream) Drop(n int)  *Float64Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float64Stream) Filter(fn func(int, float64)bool)  *Float64Stream {
	value := make([]float64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float64Stream) First() float64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Float64Stream) Last() float64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Float64Stream) Map(fn func(int, float64)) *Float64Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float64Stream) Reduce(fn func(float64, float64, int) float64,initial float64) float64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float64Stream) Reverse()  *Float64Stream {
	value := make([]float64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float64Stream) Unique()  *Float64Stream{
	value := make([]float64, 0, len(c.value))
	seen:=make(map[float64]struct{})
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

func(c *Float64Stream) Append(given float64) *Float64Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Float64Stream) Len() int {
	return len(c.value)
}

func(c *Float64Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Float64Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Float64Stream)  Sort()  *Float64Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] < c.value[j]
	})
	return c 
}

func(c *Float64Stream) All(fn func(int, float64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Float64Stream) Any(fn func(int, float64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Float64Stream) Paginate(size int)  [][]float64 {
	var pages  [][]float64
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

func(c *Float64Stream) Pop() float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Float64Stream) Prepend(given float64) *Float64Stream {
	c.value = append([]float64{given},c.value...)
	return c
}

func(c *Float64Stream) Max() float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max float64 = c.value[0]
	for _,each := range c.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(c *Float64Stream) Min() float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min float64 = c.value[0]
	for _,each := range c.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *Float64Stream) Random() float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Float64Stream) Shuffle() *Float64Stream {
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

func(c *Float64Stream) Collect() []float64{
	return c.value
}


type Float64PStream struct{
	value	[]*float64
	defaultReturn *float64
}

func PStreamOfFloat64(value []*float64) *Float64PStream {
	return &Float64PStream{value:value,defaultReturn:nil}
}

func(c *Float64PStream) OrElse(defaultReturn *float64)  *Float64PStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *Float64PStream) Concate(given []*float64)  *Float64PStream {
	value := make([]*float64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Float64PStream) Drop(n int)  *Float64PStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Float64PStream) Filter(fn func(int, *float64)bool)  *Float64PStream {
	value := make([]*float64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Float64PStream) First() *float64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Float64PStream) Last() *float64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Float64PStream) Map(fn func(int, *float64)) *Float64PStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Float64PStream) Reduce(fn func(*float64, *float64, int) *float64,initial *float64) *float64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Float64PStream) Reverse()  *Float64PStream {
	value := make([]*float64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Float64PStream) Unique()  *Float64PStream{
	value := make([]*float64, 0, len(c.value))
	seen:=make(map[*float64]struct{})
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

func(c *Float64PStream) Append(given *float64) *Float64PStream {
	c.value=append(c.value,given)
	return c
}

func(c *Float64PStream) Len() int {
	return len(c.value)
}

func(c *Float64PStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Float64PStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Float64PStream)  Sort(less func(*float64,*float64) bool )  *Float64PStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Float64PStream) All(fn func(int, *float64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Float64PStream) Any(fn func(int, *float64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Float64PStream) Paginate(size int)  [][]*float64 {
	var pages  [][]*float64
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

func(c *Float64PStream) Pop() *float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Float64PStream) Prepend(given *float64) *Float64PStream {
	c.value = append([]*float64{given},c.value...)
	return c
}

func(c *Float64PStream) Max() *float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *float64 = c.value[0]
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


func(c *Float64PStream) Min() *float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *float64 = c.value[0]
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

func(c *Float64PStream) Random() *float64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Float64PStream) Shuffle() *Float64PStream {
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

func(c *Float64PStream) Collect() []*float64{
	return c.value
}
