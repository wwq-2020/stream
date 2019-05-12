package outter
			import (
				"sort"
				"math/rand"
						commons "github.com/wwq1988/stream/commons"						
					
				)
	type SomeStream struct{
		value	[]Some
		defaultReturn Some
	}
	
	func StreamOfSome(value []Some) *SomeStream {
		return &SomeStream{value:value, defaultReturn:Some{}}
	}

	func(c *SomeStream) OrElse(defaultReturn Some)  *SomeStream {
		c.defaultReturn = defaultReturn
		return c
	}	

	func(c *SomeStream) Concate(given []Some)  *SomeStream {
		value := make([]Some, len(c.value)+len(given))
		copy(value,c.value)
		copy(value[len(c.value):],given)
		c.value = value
		return c
	}
	
	func(c *SomeStream) Drop(n int)  *SomeStream {
		l := len(c.value) - n
		if l < 0 {
			l = 0
		}
		c.value = c.value[len(c.value)-l:]
		return c
	}
	
	func(c *SomeStream) Filter(fn func(int, Some)bool)  *SomeStream {
		value := make([]Some, 0, len(c.value))
		for i, each := range c.value {
			if fn(i,each){
				value = append(value,each)
			}
		}
		c.value = value
		return c
	}
	
	func(c *SomeStream) First() Some {
		if len(c.value) <= 0 {
			return c.defaultReturn
		} 
		return c.value[0]
	}
	
	func(c *SomeStream) Last() Some {
		if len(c.value) <= 0 {
			return c.defaultReturn
		} 
		return c.value[len(c.value)-1]
	}
	
	func(c *SomeStream) Map(fn func(int, Some)) *SomeStream {
		for i, each := range c.value {
			fn(i,each)
		}
		return c
	}
	
	func(c *SomeStream) Reduce(fn func(Some, Some, int) Some,initial Some) Some   {
		final := initial
		for i, each := range c.value {
			final = fn(final,each,i)
		}
		return final
	}
	
	func(c *SomeStream) Reverse()  *SomeStream {
		value := make([]Some, len(c.value))
		for i, each := range c.value {
			value[len(c.value)-1-i] = each
		}
		c.value = value
		return c
	}
	
	func(c *SomeStream) UniqueBy(compare func(Some,Some)bool)  *SomeStream{
		value := make([]Some, 0, len(c.value))
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
				if compare(inner,outter) {
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
	
	func(c *SomeStream) Append(given Some) *SomeStream {
		c.value=append(c.value,given)
		return c
	}
	
	func(c *SomeStream) Len() int {
		return len(c.value)
	}
	
	func(c *SomeStream) IsEmpty() bool {
		return len(c.value) == 0
	}
	
	func(c *SomeStream) IsNotEmpty() bool {
		return len(c.value) != 0
	}
	
	func(c *SomeStream)  SortBy(less func(Some,Some)bool)  *SomeStream {
		sort.Slice(c.value, func(i,j int)bool{
			return less(c.value[i],c.value[j])
		})
		return c 
	}
	
	func(c *SomeStream) All(fn func(int, Some)bool)  bool {
		for i, each := range c.value {
			if !fn(i,each){
				return false
			}
		}
		return true
	}
	
	func(c *SomeStream) Any(fn func(int, Some)bool)  bool {
		for i, each := range c.value {
			if fn(i,each){
				return true
			}
		}
		return false
	}
	
	func(c *SomeStream) Paginate(size int)  [][]Some {
		var pages  [][]Some
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
	
	func(c *SomeStream) Pop() Some{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		lastIdx := len(c.value)-1
		val:=c.value[lastIdx]
		c.value[lastIdx]=c.defaultReturn
		c.value=c.value[:lastIdx]
		return val
	}
	
	func(c *SomeStream) Prepend(given Some) *SomeStream {
		c.value = append([]Some{given},c.value...)
		return c
	}
	
	func(c *SomeStream) Max(bigger func(Some,Some)bool) Some{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		var max Some = c.value[0]
		for _,each := range c.value {
			if bigger(each, max) {
				max = each
			}
		}
		return max
	}
	
	
	func(c *SomeStream) Min(less func(Some,Some)bool) Some{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		var min Some = c.value[0]
		for _,each := range c.value {
			if less(each, min) {
				min = each
			}
		}
		return min
	}
	
	func(c *SomeStream) Random() Some{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		n := rand.Intn(len(c.value))
		return c.value[n]
	}
	
	func(c *SomeStream) Shuffle() *SomeStream {
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
	
	
	
	
	
	
	
	
	func(c *SomeStream)  AStream()  *commons.StringStream {	
		value := make([]string, 0, len(c.value))	
		for _, each := range c.value {
			value = append(value, each.A)
		}
		newStream := commons.StreamOfString(value)
		return newStream
	}
	
	
	
	
	
	func(c *SomeStream)  BStream()  *commons.StringStream {	
		value := make([]string, 0, len(c.value))	
		for _, each := range c.value {
			value = append(value, each.B)
		}
		newStream := commons.StreamOfString(value)
		return newStream
	}
	
	
	
	
	
	func(c *SomeStream)  As()  []string {	
		value := make([]string, 0, len(c.value))	
		for _, each := range c.value {
			value = append(value, each.A)
		}
		return value
	}
	
	func(c *SomeStream)  Bs()  []string {	
		value := make([]string, 0, len(c.value))	
		for _, each := range c.value {
			value = append(value, each.B)
		}
		return value
	}
	
	
	func(c *SomeStream) Collect() []Some{
		return c.value
	}

type SomePStream struct{
	value	[]*Some
	defaultReturn *Some
}

func PStreamOfSome(value []*Some) *SomePStream {
	return &SomePStream{value:value,defaultReturn:nil}
}
func(c *SomePStream) OrElse(defaultReturn *Some)  *SomePStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *SomePStream) Concate(given []*Some)  *SomePStream {
	value := make([]*Some, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *SomePStream) Drop(n int)  *SomePStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *SomePStream) Filter(fn func(int, *Some)bool)  *SomePStream {
	value := make([]*Some, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *SomePStream) First() *Some {
	if len(c.value) <= 0 {
		return c.defaultReturn 
	} 
	return c.value[0]
}

func(c *SomePStream) Last() *Some {
	if len(c.value) <= 0 {
		return c.defaultReturn 
	} 
	return c.value[len(c.value)-1]
}

func(c *SomePStream) Map(fn func(int, *Some)) *SomePStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *SomePStream) Reduce(fn func(*Some, *Some, int) *Some,initial *Some) *Some   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *SomePStream) Reverse()  *SomePStream {
	value := make([]*Some, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *SomePStream) UniqueBy(compare func(*Some,*Some)bool)  *SomePStream{
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
			if compare(inner,outter) {
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

func(c *SomePStream) Append(given *Some) *SomePStream {
	c.value=append(c.value,given)
	return c
}

func(c *SomePStream) Len() int {
	return len(c.value)
}

func(c *SomePStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *SomePStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *SomePStream)  SortBy(less func(*Some,*Some)bool)  *SomePStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *SomePStream) All(fn func(int, *Some)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *SomePStream) Any(fn func(int, *Some)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *SomePStream) Paginate(size int)  [][]*Some {
	var pages  [][]*Some
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

func(c *SomePStream) Pop() *Some{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value[lastIdx]=c.defaultReturn
	c.value=c.value[:lastIdx]
	return val
}

func(c *SomePStream) Prepend(given *Some) *SomePStream {
	c.value = append([]*Some{given},c.value...)
	return c
}

func(c *SomePStream) Max(bigger func(*Some,*Some)bool) *Some{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *Some  = c.value[0]
	for _,each := range c.value {
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(c *SomePStream) Min(less func(*Some,*Some)bool) *Some{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *Some = c.value[0]
	for _,each := range c.value {
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(c *SomePStream) Random() *Some{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *SomePStream) Shuffle() *SomePStream {
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








func(c *SomePStream)  AStream()  *commons.StringStream {	
	value := make([]string, 0, len(c.value))	
	for _, each := range c.value {
		value = append(value, each.A)
	}
	newStream := commons.StreamOfString(value)
	return newStream
}





func(c *SomePStream)  BStream()  *commons.StringStream {	
	value := make([]string, 0, len(c.value))	
	for _, each := range c.value {
		value = append(value, each.B)
	}
	newStream := commons.StreamOfString(value)
	return newStream
}





func(c *SomePStream)  As()  []string {	
	value := make([]string, 0, len(c.value))	
	for _, each := range c.value {
		value = append(value, each.A)
	}
	return value
}

func(c *SomePStream)  Bs()  []string {	
	value := make([]string, 0, len(c.value))	
	for _, each := range c.value {
		value = append(value, each.B)
	}
	return value
}


func(c *SomePStream) Collect() []*Some{
	return c.value
}
