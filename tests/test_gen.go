package tests
					import (
						"sort"
						"math/rand"
						commons "github.com/wwq1988/stream/commons"
						"github.com/wwq1988/stream/outter"						
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
	
	
	func(c *SomeStream)  SortByA(less func(string,string)bool)  *SomeStream {
		sort.Slice(c.value, func(i,j int)bool{
			return less(c.value[i].A,c.value[j].A)
		})
		return c 
	}
	
	func(c *SomeStream)  SortByB(less func(string,string)bool)  *SomeStream {
		sort.Slice(c.value, func(i,j int)bool{
			return less(c.value[i].B,c.value[j].B)
		})
		return c 
	}
	
	func(c *SomeStream)  SortByC(less func(*Some,*Some)bool)  *SomeStream {
		sort.Slice(c.value, func(i,j int)bool{
			return less(c.value[i].C,c.value[j].C)
		})
		return c 
	}
	
	
	
	func(c *SomeStream)  UniqueByA(compare func(string,string)bool)  *SomeStream {
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
				if compare(inner.A,outter.A) {
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
	
	func(c *SomeStream)  UniqueByB(compare func(string,string)bool)  *SomeStream {
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
				if compare(inner.B,outter.B) {
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
	
	func(c *SomeStream)  UniqueByC(compare func(*Some,*Some)bool)  *SomeStream {
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
				if compare(inner.C,outter.C) {
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
	
	
	
	
	
	func(c *SomeStream)  CPStream()  *SomePStream {	
		value := make([]*Some, 0, len(c.value))	
		for _, each := range c.value {
			value = append(value, each.C)
		}
		newStream := PStreamOfSome(value)
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
	
	func(c *SomeStream)  Cs()  []*Some {	
		value := make([]*Some, 0, len(c.value))	
		for _, each := range c.value {
			value = append(value, each.C)
		}
		return value
	}
	
	func(c *SomeStream)  Ds()  []*outter.Some {	
		value := make([]*outter.Some, 0, len(c.value))	
		for _, each := range c.value {
			value = append(value, each.D)
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


func(c *SomePStream)  SortByA(less func(string,string)bool)  *SomePStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i].A,c.value[j].A)
	})
	return c 
}

func(c *SomePStream)  SortByB(less func(string,string)bool)  *SomePStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i].B,c.value[j].B)
	})
	return c 
}

func(c *SomePStream)  SortByC(less func(*Some,*Some)bool)  *SomePStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i].C,c.value[j].C)
	})
	return c 
}



func(c *SomePStream)  UniqueByA(compare func(string,string)bool)  *SomePStream {
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
			if compare(inner.A,outter.A) {
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

func(c *SomePStream)  UniqueByB(compare func(string,string)bool)  *SomePStream {
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
			if compare(inner.B,outter.B) {
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

func(c *SomePStream)  UniqueByC(compare func(*Some,*Some)bool)  *SomePStream {
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
			if compare(inner.C,outter.C) {
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





func(c *SomePStream)  CPStream()  *SomePStream {	
	value := make([]*Some, 0, len(c.value))	
	for _, each := range c.value {
		value = append(value, each.C)
	}
	newStream := PStreamOfSome(value)
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

func(c *SomePStream)  Cs()  []*Some {	
	value := make([]*Some, 0, len(c.value))	
	for _, each := range c.value {
		value = append(value, each.C)
	}
	return value
}

func(c *SomePStream)  Ds()  []*outter.Some {	
	value := make([]*outter.Some, 0, len(c.value))	
	for _, each := range c.value {
		value = append(value, each.D)
	}
	return value
}


func(c *SomePStream) Collect() []*Some{
	return c.value
}

	type BStream struct{
		value	[]B
		defaultReturn B
	}
	
	func StreamOfB(value []B) *BStream {
		return &BStream{value:value, defaultReturn:B{}}
	}

	func(c *BStream) OrElse(defaultReturn B)  *BStream {
		c.defaultReturn = defaultReturn
		return c
	}	

	func(c *BStream) Concate(given []B)  *BStream {
		value := make([]B, len(c.value)+len(given))
		copy(value,c.value)
		copy(value[len(c.value):],given)
		c.value = value
		return c
	}
	
	func(c *BStream) Drop(n int)  *BStream {
		l := len(c.value) - n
		if l < 0 {
			l = 0
		}
		c.value = c.value[len(c.value)-l:]
		return c
	}
	
	func(c *BStream) Filter(fn func(int, B)bool)  *BStream {
		value := make([]B, 0, len(c.value))
		for i, each := range c.value {
			if fn(i,each){
				value = append(value,each)
			}
		}
		c.value = value
		return c
	}
	
	func(c *BStream) First() B {
		if len(c.value) <= 0 {
			return c.defaultReturn
		} 
		return c.value[0]
	}
	
	func(c *BStream) Last() B {
		if len(c.value) <= 0 {
			return c.defaultReturn
		} 
		return c.value[len(c.value)-1]
	}
	
	func(c *BStream) Map(fn func(int, B)) *BStream {
		for i, each := range c.value {
			fn(i,each)
		}
		return c
	}
	
	func(c *BStream) Reduce(fn func(B, B, int) B,initial B) B   {
		final := initial
		for i, each := range c.value {
			final = fn(final,each,i)
		}
		return final
	}
	
	func(c *BStream) Reverse()  *BStream {
		value := make([]B, len(c.value))
		for i, each := range c.value {
			value[len(c.value)-1-i] = each
		}
		c.value = value
		return c
	}
	
	func(c *BStream) UniqueBy(compare func(B,B)bool)  *BStream{
		value := make([]B, 0, len(c.value))
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
	
	func(c *BStream) Append(given B) *BStream {
		c.value=append(c.value,given)
		return c
	}
	
	func(c *BStream) Len() int {
		return len(c.value)
	}
	
	func(c *BStream) IsEmpty() bool {
		return len(c.value) == 0
	}
	
	func(c *BStream) IsNotEmpty() bool {
		return len(c.value) != 0
	}
	
	func(c *BStream)  SortBy(less func(B,B)bool)  *BStream {
		sort.Slice(c.value, func(i,j int)bool{
			return less(c.value[i],c.value[j])
		})
		return c 
	}
	
	func(c *BStream) All(fn func(int, B)bool)  bool {
		for i, each := range c.value {
			if !fn(i,each){
				return false
			}
		}
		return true
	}
	
	func(c *BStream) Any(fn func(int, B)bool)  bool {
		for i, each := range c.value {
			if fn(i,each){
				return true
			}
		}
		return false
	}
	
	func(c *BStream) Paginate(size int)  [][]B {
		var pages  [][]B
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
	
	func(c *BStream) Pop() B{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		lastIdx := len(c.value)-1
		val:=c.value[lastIdx]
		c.value[lastIdx]=c.defaultReturn
		c.value=c.value[:lastIdx]
		return val
	}
	
	func(c *BStream) Prepend(given B) *BStream {
		c.value = append([]B{given},c.value...)
		return c
	}
	
	func(c *BStream) Max(bigger func(B,B)bool) B{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		var max B = c.value[0]
		for _,each := range c.value {
			if bigger(each, max) {
				max = each
			}
		}
		return max
	}
	
	
	func(c *BStream) Min(less func(B,B)bool) B{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		var min B = c.value[0]
		for _,each := range c.value {
			if less(each, min) {
				min = each
			}
		}
		return min
	}
	
	func(c *BStream) Random() B{
		if len(c.value) <= 0 {
			return c.defaultReturn
		}
		n := rand.Intn(len(c.value))
		return c.value[n]
	}
	
	func(c *BStream) Shuffle() *BStream {
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
	
	
	
	
	
	
	
	
	
	func(c *BStream) Collect() []B{
		return c.value
	}

type BPStream struct{
	value	[]*B
	defaultReturn *B
}

func PStreamOfB(value []*B) *BPStream {
	return &BPStream{value:value,defaultReturn:nil}
}
func(c *BPStream) OrElse(defaultReturn *B)  *BPStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *BPStream) Concate(given []*B)  *BPStream {
	value := make([]*B, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *BPStream) Drop(n int)  *BPStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *BPStream) Filter(fn func(int, *B)bool)  *BPStream {
	value := make([]*B, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *BPStream) First() *B {
	if len(c.value) <= 0 {
		return c.defaultReturn 
	} 
	return c.value[0]
}

func(c *BPStream) Last() *B {
	if len(c.value) <= 0 {
		return c.defaultReturn 
	} 
	return c.value[len(c.value)-1]
}

func(c *BPStream) Map(fn func(int, *B)) *BPStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *BPStream) Reduce(fn func(*B, *B, int) *B,initial *B) *B   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *BPStream) Reverse()  *BPStream {
	value := make([]*B, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *BPStream) UniqueBy(compare func(*B,*B)bool)  *BPStream{
	value := make([]*B, 0, len(c.value))
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

func(c *BPStream) Append(given *B) *BPStream {
	c.value=append(c.value,given)
	return c
}

func(c *BPStream) Len() int {
	return len(c.value)
}

func(c *BPStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *BPStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *BPStream)  SortBy(less func(*B,*B)bool)  *BPStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *BPStream) All(fn func(int, *B)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *BPStream) Any(fn func(int, *B)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *BPStream) Paginate(size int)  [][]*B {
	var pages  [][]*B
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

func(c *BPStream) Pop() *B{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value[lastIdx]=c.defaultReturn
	c.value=c.value[:lastIdx]
	return val
}

func(c *BPStream) Prepend(given *B) *BPStream {
	c.value = append([]*B{given},c.value...)
	return c
}

func(c *BPStream) Max(bigger func(*B,*B)bool) *B{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *B  = c.value[0]
	for _,each := range c.value {
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(c *BPStream) Min(less func(*B,*B)bool) *B{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *B = c.value[0]
	for _,each := range c.value {
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(c *BPStream) Random() *B{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *BPStream) Shuffle() *BPStream {
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









func(c *BPStream) Collect() []*B{
	return c.value
}
