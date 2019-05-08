
package commons

import (
	"sort"
	"math/rand"
)

const EmptyString string =""

type StringChain struct{
	value	[]string
}

func NewStringChain(value []string) *StringChain {
	return &StringChain{value:value}
}

func(c *StringChain) Concate(given []string)  *StringChain {
	value := make([]string, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *StringChain) Drop(n int)  *StringChain {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *StringChain) Filter(fn func(int, string)bool)  *StringChain {
	value := make([]string, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *StringChain) First() string {
	if len(c.value) <= 0 {
		return EmptyString
	} 
	return c.value[0]
}

func(c *StringChain) Last() string {
	if len(c.value) <= 0 {
		return EmptyString
	} 
	return c.value[len(c.value)-1]
}

func(c *StringChain) Map(fn func(int, string)) *StringChain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *StringChain) Reduce(fn func(string, string, int) string,initial string) string   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *StringChain) Reverse()  *StringChain {
	value := make([]string, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *StringChain) Unique()  *StringChain{
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

func(c *StringChain) Append(given string) *StringChain {
	c.value=append(c.value,given)
	return c
}

func(c *StringChain) Len() int {
	return len(c.value)
}

func(c *StringChain) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *StringChain) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *StringChain)  Sort()  *StringChain {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] <= (c.value[j])
	})
	return c 
}

func(c *StringChain) All(fn func(int, string)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *StringChain) Any(fn func(int, string)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *StringChain) Paginate(size int)  [][]string {
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

func(c *StringChain) Pop() string{
	if len(c.value) <= 0 {
		return EmptyString 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *StringChain) Prepend(given string) *StringChain {
	c.value = append([]string{given},c.value...)
	return c
}

func(c *StringChain) Max() string{
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


func(c *StringChain) Min() string{
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

func(c *StringChain) Random() string{
	if len(c.value) <= 0 {
		return EmptyString 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *StringChain) Shuffle() *StringChain {
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

func(c *StringChain) Collect() []string{
	return c.value
}
