
package commons

import (
	"sort"
	"math/rand"
)

type StringStream struct{
	value	[]string
	defaultReturn string
}

func StreamOfString(value []string) *StringStream {
	return &StringStream{value:value,defaultReturn:""}
}

func(c *StringStream) OrElase(defaultReturn string)  *StringStream {
	c.defaultReturn = defaultReturn
	return c
}


func(c *StringStream) Concate(given []string)  *StringStream {
	value := make([]string, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *StringStream) Drop(n int)  *StringStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *StringStream) Filter(fn func(int, string)bool)  *StringStream {
	value := make([]string, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *StringStream) First() string {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *StringStream) Last() string {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *StringStream) Map(fn func(int, string)) *StringStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *StringStream) Reduce(fn func(string, string, int) string,initial string) string   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *StringStream) Reverse()  *StringStream {
	value := make([]string, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *StringStream) Unique()  *StringStream{
	value := make([]string, 0, len(c.value))
	seen:=make(map[string]struct{})
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

func(c *StringStream) Append(given string) *StringStream {
	c.value=append(c.value,given)
	return c
}

func(c *StringStream) Len() int {
	return len(c.value)
}

func(c *StringStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *StringStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *StringStream)  Sort()  *StringStream {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] < c.value[j]
	})
	return c 
}

func(c *StringStream) All(fn func(int, string)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *StringStream) Any(fn func(int, string)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *StringStream) Paginate(size int)  [][]string {
	var pages  [][]string
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

func(c *StringStream) Pop() string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *StringStream) Prepend(given string) *StringStream {
	c.value = append([]string{given},c.value...)
	return c
}

func(c *StringStream) Max() string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max string = c.value[0]
	for _,each := range c.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(c *StringStream) Min() string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min string = c.value[0]
	for _,each := range c.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *StringStream) Random() string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *StringStream) Shuffle() *StringStream {
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

func(c *StringStream) Collect() []string{
	return c.value
}


type StringPStream struct{
	value	[]*string
	defaultReturn *string
}

func PStreamOfString(value []*string) *StringPStream {
	return &StringPStream{value:value,defaultReturn:nil}
}

func(c *StringPStream) OrElse(defaultReturn *string)  *StringPStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *StringPStream) Concate(given []*string)  *StringPStream {
	value := make([]*string, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *StringPStream) Drop(n int)  *StringPStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *StringPStream) Filter(fn func(int, *string)bool)  *StringPStream {
	value := make([]*string, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *StringPStream) First() *string {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *StringPStream) Last() *string {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *StringPStream) Map(fn func(int, *string)) *StringPStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *StringPStream) Reduce(fn func(*string, *string, int) *string,initial *string) *string   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *StringPStream) Reverse()  *StringPStream {
	value := make([]*string, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *StringPStream) Unique()  *StringPStream{
	value := make([]*string, 0, len(c.value))
	seen:=make(map[*string]struct{})
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

func(c *StringPStream) Append(given *string) *StringPStream {
	c.value=append(c.value,given)
	return c
}

func(c *StringPStream) Len() int {
	return len(c.value)
}

func(c *StringPStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *StringPStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *StringPStream)  Sort(less func(*string,*string) bool )  *StringPStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *StringPStream) All(fn func(int, *string)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *StringPStream) Any(fn func(int, *string)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *StringPStream) Paginate(size int)  [][]*string {
	var pages  [][]*string
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

func(c *StringPStream) Pop() *string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *StringPStream) Prepend(given *string) *StringPStream {
	c.value = append([]*string{given},c.value...)
	return c
}

func(c *StringPStream) Max() *string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *string = c.value[0]
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


func(c *StringPStream) Min() *string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *string = c.value[0]
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

func(c *StringPStream) Random() *string{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *StringPStream) Shuffle() *StringPStream {
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

func(c *StringPStream) Collect() []*string{
	return c.value
}
