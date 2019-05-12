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

	
	func(s *SomeStream) FilterByA(fn func(int,string)bool)  *SomeStream {
		value := make([]Some, 0, len(s.value))
		for i, each := range s.value {
			if fn(i,each.A){
				value = append(value,each)
			}
		}
		s.value = value
		return s
	}
	
	func(s *SomeStream) FilterByB(fn func(int,string)bool)  *SomeStream {
		value := make([]Some, 0, len(s.value))
		for i, each := range s.value {
			if fn(i,each.B){
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
	
	
	
	func(s *SomeStream)  UniqueByA()  *SomeStream {
		value := make([]Some, 0, len(s.value))
		seen:=make(map[string]struct{})
		for _, each := range s.value {
			if _,dup:=seen[each.A];dup{
				continue
			}
			value = append(value, each)
			seen[each.A]=struct{}{}	
		}
		s.value = value
		return s
	}
	
	
	
	func(s *SomeStream)  UniqueByB()  *SomeStream {
		value := make([]Some, 0, len(s.value))
		seen:=make(map[string]struct{})
		for _, each := range s.value {
			if _,dup:=seen[each.B];dup{
				continue
			}
			value = append(value, each)
			seen[each.B]=struct{}{}	
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
	
	
	
	func(s *SomeStream)  SortByA()  *SomeStream {
		sort.Slice(s.value, func(i,j int)bool{
			return s.value[i].A < s.value[j].A
		})
		return s 
	}
	
	
	
	func(s *SomeStream)  SortByB()  *SomeStream {
		sort.Slice(s.value, func(i,j int)bool{
			return s.value[i].B < s.value[j].B
		})
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


func(s *SomePStream) FilterByA(fn func(int,string)bool)  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each.A){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *SomePStream) FilterByB(fn func(int,string)bool)  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each.B){
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


func(s *SomePStream) AllByA(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each.A){
			return false
		}
	}
	return true
}

func(s *SomePStream) AllByB(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each.B){
			return false
		}
	}
	return true
}




func(s *SomeStream) AllByA(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each.A){
			return false
		}
	}
	return true
}

func(s *SomeStream) AllByB(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each.B){
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



func(s *SomePStream) AnyByA(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if fn(i,each.A){
			return true
		}
	}
	return false
}

func(s *SomePStream) AnyByB(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if fn(i,each.B){
			return true
		}
	}
	return false
}



func(s *SomeStream) AnyByA(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if fn(i,each.A){
			return true
		}
	}
	return false
}

func(s *SomeStream) AnyByB(fn func(int,string)bool)  bool {
	for i, each := range s.value {
		if fn(i,each.B){
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



func(s *SomePStream)  SortByA()  *SomePStream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i].A < s.value[j].A
	})
	return s 
}



func(s *SomePStream)  SortByB()  *SomePStream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i].B < s.value[j].B
	})
	return s 
}





func(s *SomePStream)  UniqueByA()  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[string]struct{})
	for _, each := range s.value {
		if _,dup:=seen[each.A];dup{
			continue
		}
		value = append(value, each)
		seen[each.A]=struct{}{}	
	}
	s.value = value
	return s
}



func(s *SomePStream)  UniqueByB()  *SomePStream {
	value := make([]*Some, 0, len(s.value))
	seen:=make(map[string]struct{})
	for _, each := range s.value {
		if _,dup:=seen[each.B];dup{
			continue
		}
		value = append(value, each)
		seen[each.B]=struct{}{}	
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

func(s *SomePStream) Collect() []*Some{
	return s.value
}
