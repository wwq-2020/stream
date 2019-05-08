
package commons

import (
	"sort"
	"math/rand"
)

const EmptyString string =""

type StringCollection struct{
	value	[]string
}

func NewStringCollection(value []string) *StringCollection {
	return &StringCollection{value:value}
}

func(c *StringCollection) Concate(given []string)  *StringCollection {
	value := make([]string, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *StringCollection) Drop(n int)  *StringCollection {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *StringCollection) Filter(fn func(int, string)bool)  *StringCollection {
	value := make([]string, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *StringCollection) First() string {
	if len(c.value) <= 0 {
		return EmptyString
	} 
	return c.value[0]
}

func(c *StringCollection) Last() string {
	if len(c.value) <= 0 {
		return EmptyString
	} 
	return c.value[len(c.value)-1]
}

func(c *StringCollection) Map(fn func(int, string)) *StringCollection {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *StringCollection) Reduce(fn func(string, string, int) string,initial string) string   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *StringCollection) Reverse()  *StringCollection {
	value := make([]string, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *StringCollection) Unique()  *StringCollection{
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

func(c *StringCollection) Append(given string) *StringCollection {
	c.value=append(c.value,given)
	return c
}

func(c *StringCollection) Len() int {
	return len(c.value)
}

func(c *StringCollection) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *StringCollection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *StringCollection)  Sort()  *StringCollection {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *StringCollection) All(fn func(int, string)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *StringCollection) Any(fn func(int, string)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *StringCollection) Paginate(size int)  [][]string {
	var pages  [][]string
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

func(c *StringCollection) Pop() string{
	if len(c.value) <= 0 {
		return EmptyString 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *StringCollection) Prepend(given string) *StringCollection {
	c.value = append([]string{given},c.value...)
	return c
}

func(c *StringCollection) Max() string{
	if len(c.value) <= 0 {
		return EmptyString 
	}
	var max string
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


func(c *StringCollection) Min() string{
	if len(c.value) <= 0 {
		return EmptyString 
	}
	var min string
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

func(c *StringCollection) Random() string{
	if len(c.value) <= 0 {
		return EmptyString 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *StringCollection) Shuffle() *StringCollection {
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

func(c *StringCollection) Collect() []string{
	return c.value
}
