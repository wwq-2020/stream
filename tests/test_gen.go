package tests
				import (
					"sort"
					"math/rand"
				)
type SomeStream struct{
	value	[]*Some
}

func StreamOfSome(value []*Some) *SomeStream {
	return &SomeStream{value:value}
}

func(c *SomeStream) Concate(given []*Some)  *SomeStream {
	value := make([]*Some, len(c.value)+len(given))
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

func(c *SomeStream) Filter(fn func(int, *Some)bool)  *SomeStream {
	value := make([]*Some, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *SomeStream) First() *Some {
	if len(c.value) < 0 {
		return nil
	} 
	return c.value[0]
}

func(c *SomeStream) Last() *Some {
	if len(c.value) < 0 {
		return nil
	} 
	return c.value[len(c.value)-1]
}

func(c *SomeStream) Map(fn func(int, *Some)) *SomeStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *SomeStream) Reduce(fn func(*Some, *Some, int) *Some,initial *Some) *Some   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *SomeStream) Reverse()  *SomeStream {
	value := make([]*Some, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *SomeStream) UniqueBy(compare func(*Some,*Some)bool)  *SomeStream{
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

func(c *SomeStream) Append(given *Some) *SomeStream {
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

func(c *SomeStream)  SortBy(less func(*Some,*Some)bool)  *SomeStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *SomeStream) All(fn func(int, *Some)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *SomeStream) Any(fn func(int, *Some)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *SomeStream) Paginate(size int)  [][]*Some {
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

func(c *SomeStream) Pop() *Some{
	if len(c.value) < 0 {
		return nil
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value[lastIdx]=nil
	c.value=c.value[:lastIdx]
	return val
}

func(c *SomeStream) Prepend(given *Some) *SomeStream {
	c.value = append([]*Some{given},c.value...)
	return c
}

func(c *SomeStream) Max(bigger func(*Some,*Some)bool) *Some{
	if len(c.value) < 0 {
		return nil
	}
	var max *Some
	for _,each := range c.value {
		if max==nil{
			max=each
			continue
		}
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(c *SomeStream) Min(less func(*Some,*Some)bool) *Some{
	if len(c.value) < 0 {
		return nil
	}
	var min *Some
	for _,each := range c.value {
		if min==nil{
			min=each
			continue
		}
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(c *SomeStream) Random() *Some{
	if len(c.value) < 0 {
		return nil
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *SomeStream) Shuffle() *SomeStream {
	if len(c.value) < 0 {
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

func(c *SomeStream)  UniqueByB(compare func(string,string)bool)  *SomeStream {
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

func(c *SomeStream)  UniqueByC(compare func(*Some,*Some)bool)  *SomeStream {
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




func(c *SomeStream) Collect() []*Some{
	return c.value
}

type BStream struct{
	value	[]*B
}

func StreamOfB(value []*B) *BStream {
	return &BStream{value:value}
}

func(c *BStream) Concate(given []*B)  *BStream {
	value := make([]*B, len(c.value)+len(given))
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

func(c *BStream) Filter(fn func(int, *B)bool)  *BStream {
	value := make([]*B, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *BStream) First() *B {
	if len(c.value) < 0 {
		return nil
	} 
	return c.value[0]
}

func(c *BStream) Last() *B {
	if len(c.value) < 0 {
		return nil
	} 
	return c.value[len(c.value)-1]
}

func(c *BStream) Map(fn func(int, *B)) *BStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *BStream) Reduce(fn func(*B, *B, int) *B,initial *B) *B   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *BStream) Reverse()  *BStream {
	value := make([]*B, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *BStream) UniqueBy(compare func(*B,*B)bool)  *BStream{
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

func(c *BStream) Append(given *B) *BStream {
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

func(c *BStream)  SortBy(less func(*B,*B)bool)  *BStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *BStream) All(fn func(int, *B)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *BStream) Any(fn func(int, *B)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *BStream) Paginate(size int)  [][]*B {
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

func(c *BStream) Pop() *B{
	if len(c.value) < 0 {
		return nil
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value[lastIdx]=nil
	c.value=c.value[:lastIdx]
	return val
}

func(c *BStream) Prepend(given *B) *BStream {
	c.value = append([]*B{given},c.value...)
	return c
}

func(c *BStream) Max(bigger func(*B,*B)bool) *B{
	if len(c.value) < 0 {
		return nil
	}
	var max *B
	for _,each := range c.value {
		if max==nil{
			max=each
			continue
		}
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(c *BStream) Min(less func(*B,*B)bool) *B{
	if len(c.value) < 0 {
		return nil
	}
	var min *B
	for _,each := range c.value {
		if min==nil{
			min=each
			continue
		}
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(c *BStream) Random() *B{
	if len(c.value) < 0 {
		return nil
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *BStream) Shuffle() *BStream {
	if len(c.value) < 0 {
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







func(c *BStream) Collect() []*B{
	return c.value
}
