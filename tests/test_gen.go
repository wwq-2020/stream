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

	func(s *SomeStream) OrElse(defaultReturn Some)  *SomeStream {
		s.defaultReturn = defaultReturn
		return s
	}	

	func(s *SomeStream) Concate(given []Some)  *SomeStream {
		value := make([]Some, len(s.value)+len(given))
		copy(value,s.value)
		copy(value[len(s.value):],given)
		s.value = value
		return s
	}
	
	func(s *SomeStream) Drop(n int)  *SomeStream {
		l := len(s.value) - n
		if l < 0 {
			l = 0
		}
		s.value = s.value[len(s.value)-l:]
		return s
	}
	
	func(s *SomeStream) Filter(fn func(int, Some)bool)  *SomeStream {
		value := make([]Some, 0, len(s.value))
		for i, each := range s.value {
			if fn(i,each){
				value = append(value,each)
			}
		}
		s.value = value
		return s
	}
	
	func(s *SomeStream) First() Some {
		if len(s.value) <= 0 {
			return s.defaultReturn
		} 
		return s.value[0]
	}
	
	func(s *SomeStream) Last() Some {
		if len(s.value) <= 0 {
			return s.defaultReturn
		} 
		return s.value[len(s.value)-1]
	}
	
	func(s *SomeStream) Map(fn func(int, Some)) *SomeStream {
		for i, each := range s.value {
			fn(i,each)
		}
		return s
	}
	
	func(s *SomeStream) Reduce(fn func(Some, Some, int) Some,initial Some) Some   {
		final := initial
		for i, each := range s.value {
			final = fn(final,each,i)
		}
		return final
	}
	
	func(s *SomeStream) Reverse()  *SomeStream {
		value := make([]Some, len(s.value))
		for i, each := range s.value {
			value[len(s.value)-1-i] = each
		}
		s.value = value
		return s
	}
	
	func(s *SomeStream) UniqueBy(compare func(Some,Some)bool)  *SomeStream{
		value := make([]Some, 0, len(s.value))
		seen:=make(map[int]struct{})
		for i, outter := range s.value {
			dup:=false
			if _,exist:=seen[i];exist{
				continue
			}		
			for j,inner :=range s.value {
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
		s.value = value
		return s
	}
	
	func(s *SomeStream) Append(given Some) *SomeStream {
		s.value=append(s.value,given)
		return s
	}
	
	func(s *SomeStream) Len() int {
		return len(s.value)
	}
	
	func(s *SomeStream) IsEmpty() bool {
		return len(s.value) == 0
	}
	
	func(s *SomeStream) IsNotEmpty() bool {
		return len(s.value) != 0
	}
	
	func(s *SomeStream)  SortBy(less func(Some,Some)bool)  *SomeStream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i],s.value[j])
		})
		return s 
	}
	
	func(s *SomeStream) All(fn func(int, Some)bool)  bool {
		for i, each := range s.value {
			if !fn(i,each){
				return false
			}
		}
		return true
	}
	
	func(s *SomeStream) Any(fn func(int, Some)bool)  bool {
		for i, each := range s.value {
			if fn(i,each){
				return true
			}
		}
		return false
	}
	
	func(s *SomeStream) Paginate(size int)  [][]Some {
		var pages  [][]Some
		prev := -1
		for i := range s.value {
			if (i-prev) < size-1 && i != (len(s.value)-1) {
				continue
			}
			pages=append(pages,s.value[prev+1:i+1])
			prev=i
		}
		return pages
	}
	
	func(s *SomeStream) Pop() Some{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		lastIdx := len(s.value)-1
		val:=s.value[lastIdx]
		s.value[lastIdx]=s.defaultReturn
		s.value=s.value[:lastIdx]
		return val
	}
	
	func(s *SomeStream) Prepend(given Some) *SomeStream {
		s.value = append([]Some{given},s.value...)
		return s
	}
	
	func(s *SomeStream) Max(bigger func(Some,Some)bool) Some{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		var max Some = s.value[0]
		for _,each := range s.value {
			if bigger(each, max) {
				max = each
			}
		}
		return max
	}
	
	
	func(s *SomeStream) Min(less func(Some,Some)bool) Some{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		var min Some = s.value[0]
		for _,each := range s.value {
			if less(each, min) {
				min = each
			}
		}
		return min
	}
	
	func(s *SomeStream) Random() Some{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		n := rand.Intn(len(s.value))
		return s.value[n]
	}
	
	func(s *SomeStream) Shuffle() *SomeStream {
		if len(s.value) <= 0 {
			return s
		}
		indexes := make([]int, len(s.value))
		for i := range s.value {
			indexes[i] = i
		}
		
		rand.Shuffle(len(s.value), func(i, j int) {
			s.value[i], s.value[j] = 	s.value[j], s.value[i] 
		})
		
		return s
	}
	
	
	func(s *SomeStream)  SortByA(less func(string,string)bool)  *SomeStream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i].A,s.value[j].A)
		})
		return s 
	}
	
	func(s *SomeStream)  SortByB(less func(string,string)bool)  *SomeStream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i].B,s.value[j].B)
		})
		return s 
	}
	
	func(s *SomeStream)  SortByC(less func(*Some,*Some)bool)  *SomeStream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i].C,s.value[j].C)
		})
		return s 
	}
	
	
	
	func(s *SomeStream)  UniqueByA(compare func(string,string)bool)  *SomeStream {
		value := make([]Some, 0, len(s.value))
		seen:=make(map[int]struct{})
		for i, outter := range s.value {
			dup:=false
			if _,exist:=seen[i];exist{
				continue
			}		
			for j,inner :=range s.value {
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
		s.value = value
		
		return s
	}
	
	func(s *SomeStream)  UniqueByB(compare func(string,string)bool)  *SomeStream {
		value := make([]Some, 0, len(s.value))
		seen:=make(map[int]struct{})
		for i, outter := range s.value {
			dup:=false
			if _,exist:=seen[i];exist{
				continue
			}		
			for j,inner :=range s.value {
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
		s.value = value
		
		return s
	}
	
	func(s *SomeStream)  UniqueByC(compare func(*Some,*Some)bool)  *SomeStream {
		value := make([]Some, 0, len(s.value))
		seen:=make(map[int]struct{})
		for i, outter := range s.value {
			dup:=false
			if _,exist:=seen[i];exist{
				continue
			}		
			for j,inner :=range s.value {
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
		s.value = value
		
		return s
	}
	
	
	
	
	
	func(s *SomeStream)  AStream()  *commons.StringStream {	
		value := make([]string, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.A)
		}
		newStream := commons.StreamOfString(value)
		return newStream
	}
	
	
	
	
	
	func(s *SomeStream)  BStream()  *commons.StringStream {	
		value := make([]string, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.B)
		}
		newStream := commons.StreamOfString(value)
		return newStream
	}
	
	
	
	
	
	func(s *SomeStream)  CPStream()  *SomePStream {	
		value := make([]*Some, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.C)
		}
		newStream := PStreamOfSome(value)
		return newStream
	}
	
	
	
	
	
	func(s *SomeStream)  DPStream()  *outter.SomePStream {	
		value := make([]*outter.Some, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.D)
		}
		newStream := outter.PStreamOfSome(value)
		return newStream
	}
	
	
	
	
	
	func(s *SomeStream)  As()  []string {	
		value := make([]string, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.A)
		}
		return value
	}
	
	func(s *SomeStream)  Bs()  []string {	
		value := make([]string, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.B)
		}
		return value
	}
	
	func(s *SomeStream)  Cs()  []*Some {	
		value := make([]*Some, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.C)
		}
		return value
	}
	
	func(s *SomeStream)  Ds()  []*outter.Some {	
		value := make([]*outter.Some, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.D)
		}
		return value
	}
	
	
	func(s *SomeStream) Collect() []Some{
		return s.value
	}

type SomePStream struct{
	value	[]*Some
	defaultReturn *Some
}

func PStreamOfSome(value []*Some) *SomePStream {
	return &SomePStream{value:value,defaultReturn:nil}
}
func(s *SomePStream) OrElse(defaultReturn *Some)  *SomePStream {
	s.defaultReturn = defaultReturn
	return s
}

func(s *SomePStream) Concate(given []*Some)  *SomePStream {
	value := make([]*Some, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}

func(s *SomePStream) Drop(n int)  *SomePStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}

func(s *SomePStream) Filter(fn func(int, *Some)bool)  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *SomePStream) First() *Some {
	if len(s.value) <= 0 {
		return s.defaultReturn 
	} 
	return s.value[0]
}

func(s *SomePStream) Last() *Some {
	if len(s.value) <= 0 {
		return s.defaultReturn 
	} 
	return s.value[len(s.value)-1]
}

func(s *SomePStream) Map(fn func(int, *Some)) *SomePStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}

func(s *SomePStream) Reduce(fn func(*Some, *Some, int) *Some,initial *Some) *Some   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}

func(s *SomePStream) Reverse()  *SomePStream {
	value := make([]*Some, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

func(s *SomePStream) UniqueBy(compare func(*Some,*Some)bool)  *SomePStream{
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range s.value {
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
	s.value = value
	return s
}

func(s *SomePStream) Append(given *Some) *SomePStream {
	s.value=append(s.value,given)
	return s
}

func(s *SomePStream) Len() int {
	return len(s.value)
}

func(s *SomePStream) IsEmpty() bool {
	return len(s.value) == 0
}

func(s *SomePStream) IsNotEmpty() bool {
	return len(s.value) != 0
}

func(s *SomePStream)  SortBy(less func(*Some,*Some)bool)  *SomePStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}

func(s *SomePStream) All(fn func(int, *Some)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(s *SomePStream) Any(fn func(int, *Some)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(s *SomePStream) Paginate(size int)  [][]*Some {
	var pages  [][]*Some
	prev := -1
	for i := range s.value {
		if (i-prev) < size-1 && i != (len(s.value)-1) {
			continue
		}
		pages=append(pages,s.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(s *SomePStream) Pop() *Some{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value[lastIdx]=s.defaultReturn
	s.value=s.value[:lastIdx]
	return val
}

func(s *SomePStream) Prepend(given *Some) *SomePStream {
	s.value = append([]*Some{given},s.value...)
	return s
}

func(s *SomePStream) Max(bigger func(*Some,*Some)bool) *Some{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *Some  = s.value[0]
	for _,each := range s.value {
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(s *SomePStream) Min(less func(*Some,*Some)bool) *Some{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *Some = s.value[0]
	for _,each := range s.value {
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(s *SomePStream) Random() *Some{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}

func(s *SomePStream) Shuffle() *SomePStream {
	if len(s.value) <= 0 {
		return s
	}
	indexes := make([]int, len(s.value))
	for i := range s.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = 	s.value[j], s.value[i] 
	})
	
	return s
}


func(s *SomePStream)  SortByA(less func(string,string)bool)  *SomePStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i].A,s.value[j].A)
	})
	return s 
}

func(s *SomePStream)  SortByB(less func(string,string)bool)  *SomePStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i].B,s.value[j].B)
	})
	return s 
}

func(s *SomePStream)  SortByC(less func(*Some,*Some)bool)  *SomePStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i].C,s.value[j].C)
	})
	return s 
}



func(s *SomePStream)  UniqueByA(compare func(string,string)bool)  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range s.value {
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
	s.value = value
	
	return s
}

func(s *SomePStream)  UniqueByB(compare func(string,string)bool)  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range s.value {
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
	s.value = value
	
	return s
}

func(s *SomePStream)  UniqueByC(compare func(*Some,*Some)bool)  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range s.value {
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
	s.value = value
	
	return s
}





func(s *SomePStream)  AStream()  *commons.StringStream {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.A)
	}
	newStream := commons.StreamOfString(value)
	return newStream
}





func(s *SomePStream)  BStream()  *commons.StringStream {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.B)
	}
	newStream := commons.StreamOfString(value)
	return newStream
}





func(s *SomePStream)  CPStream()  *SomePStream {	
	value := make([]*Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.C)
	}
	newStream := PStreamOfSome(value)
	return newStream
}





func(s *SomePStream)  DPStream()  *outter.SomePStream {	
	value := make([]*outter.Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.D)
	}
	newStream := outter.PStreamOfSome(value)
	return newStream
}





func(s *SomePStream)  As()  []string {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.A)
	}
	return value
}

func(s *SomePStream)  Bs()  []string {	
	value := make([]string, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.B)
	}
	return value
}

func(s *SomePStream)  Cs()  []*Some {	
	value := make([]*Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.C)
	}
	return value
}

func(s *SomePStream)  Ds()  []*outter.Some {	
	value := make([]*outter.Some, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.D)
	}
	return value
}


func(s *SomePStream) Collect() []*Some{
	return s.value
}

	type BStream struct{
		value	[]B
		defaultReturn B
	}
	
	func StreamOfB(value []B) *BStream {
		return &BStream{value:value, defaultReturn:B{}}
	}

	func(s *BStream) OrElse(defaultReturn B)  *BStream {
		s.defaultReturn = defaultReturn
		return s
	}	

	func(s *BStream) Concate(given []B)  *BStream {
		value := make([]B, len(s.value)+len(given))
		copy(value,s.value)
		copy(value[len(s.value):],given)
		s.value = value
		return s
	}
	
	func(s *BStream) Drop(n int)  *BStream {
		l := len(s.value) - n
		if l < 0 {
			l = 0
		}
		s.value = s.value[len(s.value)-l:]
		return s
	}
	
	func(s *BStream) Filter(fn func(int, B)bool)  *BStream {
		value := make([]B, 0, len(s.value))
		for i, each := range s.value {
			if fn(i,each){
				value = append(value,each)
			}
		}
		s.value = value
		return s
	}
	
	func(s *BStream) First() B {
		if len(s.value) <= 0 {
			return s.defaultReturn
		} 
		return s.value[0]
	}
	
	func(s *BStream) Last() B {
		if len(s.value) <= 0 {
			return s.defaultReturn
		} 
		return s.value[len(s.value)-1]
	}
	
	func(s *BStream) Map(fn func(int, B)) *BStream {
		for i, each := range s.value {
			fn(i,each)
		}
		return s
	}
	
	func(s *BStream) Reduce(fn func(B, B, int) B,initial B) B   {
		final := initial
		for i, each := range s.value {
			final = fn(final,each,i)
		}
		return final
	}
	
	func(s *BStream) Reverse()  *BStream {
		value := make([]B, len(s.value))
		for i, each := range s.value {
			value[len(s.value)-1-i] = each
		}
		s.value = value
		return s
	}
	
	func(s *BStream) UniqueBy(compare func(B,B)bool)  *BStream{
		value := make([]B, 0, len(s.value))
		seen:=make(map[int]struct{})
		for i, outter := range s.value {
			dup:=false
			if _,exist:=seen[i];exist{
				continue
			}		
			for j,inner :=range s.value {
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
		s.value = value
		return s
	}
	
	func(s *BStream) Append(given B) *BStream {
		s.value=append(s.value,given)
		return s
	}
	
	func(s *BStream) Len() int {
		return len(s.value)
	}
	
	func(s *BStream) IsEmpty() bool {
		return len(s.value) == 0
	}
	
	func(s *BStream) IsNotEmpty() bool {
		return len(s.value) != 0
	}
	
	func(s *BStream)  SortBy(less func(B,B)bool)  *BStream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i],s.value[j])
		})
		return s 
	}
	
	func(s *BStream) All(fn func(int, B)bool)  bool {
		for i, each := range s.value {
			if !fn(i,each){
				return false
			}
		}
		return true
	}
	
	func(s *BStream) Any(fn func(int, B)bool)  bool {
		for i, each := range s.value {
			if fn(i,each){
				return true
			}
		}
		return false
	}
	
	func(s *BStream) Paginate(size int)  [][]B {
		var pages  [][]B
		prev := -1
		for i := range s.value {
			if (i-prev) < size-1 && i != (len(s.value)-1) {
				continue
			}
			pages=append(pages,s.value[prev+1:i+1])
			prev=i
		}
		return pages
	}
	
	func(s *BStream) Pop() B{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		lastIdx := len(s.value)-1
		val:=s.value[lastIdx]
		s.value[lastIdx]=s.defaultReturn
		s.value=s.value[:lastIdx]
		return val
	}
	
	func(s *BStream) Prepend(given B) *BStream {
		s.value = append([]B{given},s.value...)
		return s
	}
	
	func(s *BStream) Max(bigger func(B,B)bool) B{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		var max B = s.value[0]
		for _,each := range s.value {
			if bigger(each, max) {
				max = each
			}
		}
		return max
	}
	
	
	func(s *BStream) Min(less func(B,B)bool) B{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		var min B = s.value[0]
		for _,each := range s.value {
			if less(each, min) {
				min = each
			}
		}
		return min
	}
	
	func(s *BStream) Random() B{
		if len(s.value) <= 0 {
			return s.defaultReturn
		}
		n := rand.Intn(len(s.value))
		return s.value[n]
	}
	
	func(s *BStream) Shuffle() *BStream {
		if len(s.value) <= 0 {
			return s
		}
		indexes := make([]int, len(s.value))
		for i := range s.value {
			indexes[i] = i
		}
		
		rand.Shuffle(len(s.value), func(i, j int) {
			s.value[i], s.value[j] = 	s.value[j], s.value[i] 
		})
		
		return s
	}
	
	
	
	
	
	
	
	
	
	func(s *BStream) Collect() []B{
		return s.value
	}

type BPStream struct{
	value	[]*B
	defaultReturn *B
}

func PStreamOfB(value []*B) *BPStream {
	return &BPStream{value:value,defaultReturn:nil}
}
func(s *BPStream) OrElse(defaultReturn *B)  *BPStream {
	s.defaultReturn = defaultReturn
	return s
}

func(s *BPStream) Concate(given []*B)  *BPStream {
	value := make([]*B, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}

func(s *BPStream) Drop(n int)  *BPStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}

func(s *BPStream) Filter(fn func(int, *B)bool)  *BPStream {
	value := make([]*B, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *BPStream) First() *B {
	if len(s.value) <= 0 {
		return s.defaultReturn 
	} 
	return s.value[0]
}

func(s *BPStream) Last() *B {
	if len(s.value) <= 0 {
		return s.defaultReturn 
	} 
	return s.value[len(s.value)-1]
}

func(s *BPStream) Map(fn func(int, *B)) *BPStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}

func(s *BPStream) Reduce(fn func(*B, *B, int) *B,initial *B) *B   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}

func(s *BPStream) Reverse()  *BPStream {
	value := make([]*B, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

func(s *BPStream) UniqueBy(compare func(*B,*B)bool)  *BPStream{
	value := make([]*B, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range s.value {
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
	s.value = value
	return s
}

func(s *BPStream) Append(given *B) *BPStream {
	s.value=append(s.value,given)
	return s
}

func(s *BPStream) Len() int {
	return len(s.value)
}

func(s *BPStream) IsEmpty() bool {
	return len(s.value) == 0
}

func(s *BPStream) IsNotEmpty() bool {
	return len(s.value) != 0
}

func(s *BPStream)  SortBy(less func(*B,*B)bool)  *BPStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}

func(s *BPStream) All(fn func(int, *B)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(s *BPStream) Any(fn func(int, *B)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(s *BPStream) Paginate(size int)  [][]*B {
	var pages  [][]*B
	prev := -1
	for i := range s.value {
		if (i-prev) < size-1 && i != (len(s.value)-1) {
			continue
		}
		pages=append(pages,s.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(s *BPStream) Pop() *B{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value[lastIdx]=s.defaultReturn
	s.value=s.value[:lastIdx]
	return val
}

func(s *BPStream) Prepend(given *B) *BPStream {
	s.value = append([]*B{given},s.value...)
	return s
}

func(s *BPStream) Max(bigger func(*B,*B)bool) *B{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *B  = s.value[0]
	for _,each := range s.value {
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(s *BPStream) Min(less func(*B,*B)bool) *B{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *B = s.value[0]
	for _,each := range s.value {
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(s *BPStream) Random() *B{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}

func(s *BPStream) Shuffle() *BPStream {
	if len(s.value) <= 0 {
		return s
	}
	indexes := make([]int, len(s.value))
	for i := range s.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = 	s.value[j], s.value[i] 
	})
	
	return s
}









func(s *BPStream) Collect() []*B{
	return s.value
}
