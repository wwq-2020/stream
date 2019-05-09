
package tests

import (
	"sort"
	"math/rand"
)
type SomeCollection struct{
	value	[]*Some
}

func NewSomeCollection(value []*Some) *SomeCollection {
	return &SomeCollection{value:value}
}

func(c *SomeCollection) Concate(given []*Some)  *SomeCollection {
	value := make([]*Some, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *SomeCollection) Drop(n int)  *SomeCollection {
	l := len(c.value) - n
	if l <= 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *SomeCollection) Filter(fn func(int, *Some)bool)  *SomeCollection {
	value := make([]*Some, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *SomeCollection) First() *Some {
	if len(c.value) <= 0 {
		return nil
	} 
	return c.value[0]
}

func(c *SomeCollection) Last() *Some {
	if len(c.value) <= 0 {
		return nil
	} 
	return c.value[len(c.value)-1]
}

func(c *SomeCollection) Map(fn func(int, *Some)) *SomeCollection {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *SomeCollection) Reduce(fn func(*Some, *Some, int) *Some,initial *Some) *Some   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *SomeCollection) Reverse()  *SomeCollection {
	value := make([]*Some, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *SomeCollection) Unique()  *SomeCollection{
	value := make([]*Some, 0, len(c.value))
	seen:=make(map[int]struct{})
	for i, outter := range c.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range c.value {
			if i==j {
				continue
			}
			if inner.Compare(outter) == 0 {
				seen[j]=struct{}{}				
				dup=true
			}
		}
		if dup {
			seen[i]=struct{}{}
		}
		value=append(value,outter)			
	}
	c.value = value
	return c
}

func(c *SomeCollection) Append(given *Some) *SomeCollection {
	c.value=append(c.value,given)
	return c
}

func(c *SomeCollection) Len() int {
	return len(c.value)
}

func(c *SomeCollection) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *SomeCollection) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *SomeCollection)  Sort()  *SomeCollection {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i].Compare(c.value[j])<=0
	})
	return c 
}

func(c *SomeCollection) All(fn func(int, *Some)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *SomeCollection) Any(fn func(int, *Some)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *SomeCollection) Paginate(size int)  [][]*Some {
	var pages  [][]*Some
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

func(c *SomeCollection) Pop() *Some{
	if len(c.value) <= 0 {
		return nil
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value[lastIdx]=nil
	c.value=c.value[:lastIdx]
	return val
}

func(c *SomeCollection) Prepend(given *Some) *SomeCollection {
	c.value = append([]*Some{given},c.value...)
	return c
}

func(c *SomeCollection) Max() *Some{
	if len(c.value) <= 0 {
		return nil
	}
	var max *Some
	for _,each := range c.value {
		if max==nil{
			max=each
			continue
		}
		if max.Compare(each) <= 0 {
			max = each
		}
	}
	return max
}


func(c *SomeCollection) Min() *Some{
	if len(c.value) <= 0 {
		return nil
	}
	var min *Some
	for _,each := range c.value {
		if min==nil{
			min=each
			continue
		}
		if each.Compare(min) <= 0 {
			min = each
		}
	}
	return min
}

func(c *SomeCollection) Random() *Some{
	if len(c.value) <= 0 {
		return nil
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *SomeCollection) Shuffle() *SomeCollection {
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


func(c *SomeCollection)  SortByA()  *SomeCollection {
	
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i].A <= c.value[j].A
	})
	

	return c 
}

func(c *SomeCollection)  SortByB()  *SomeCollection {
	
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i].B <= c.value[j].B
	})
	

	return c 
}

func(c *SomeCollection)  SortByC()  *SomeCollection {
	
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i].C.Compare(c.value[j].C)<=0
	})
	

	return c 
}



func(c *SomeCollection)  UniqueByA()  *SomeCollection {
	value := make([]*Some, 0, len(c.value))
	seen:=make(map[interface{}]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each.A];exist{
			continue
		}		
		seen[each.A]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *SomeCollection)  UniqueByB()  *SomeCollection {
	value := make([]*Some, 0, len(c.value))
	seen:=make(map[interface{}]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each.B];exist{
			continue
		}		
		seen[each.B]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *SomeCollection)  UniqueByC()  *SomeCollection {
	value := make([]*Some, 0, len(c.value))
	seen:=make(map[interface{}]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each.C];exist{
			continue
		}		
		seen[each.C]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}


func(c *SomeCollection) Collect() []*Some{
	return c.value
}
