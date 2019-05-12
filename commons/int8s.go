
package commons
import (
	"sort"
	"math/rand"
)
type Int8Stream struct{
	value	[]int8
	defaultReturn int8
}
func StreamOfInt8(value []int8) *Int8Stream {
	return &Int8Stream{value:value,defaultReturn:0}
}
func(s *Int8Stream) OrElase(defaultReturn int8)  *Int8Stream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Int8Stream) Concate(given []int8)  *Int8Stream {
	value := make([]int8, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Int8Stream) Drop(n int)  *Int8Stream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Int8Stream) Filter(fn func(int, int8)bool)  *Int8Stream {
	value := make([]int8, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Int8Stream) First() int8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Int8Stream) Last() int8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Int8Stream) Map(fn func(int, int8)) *Int8Stream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Int8Stream) Reduce(fn func(int8, int8, int) int8,initial int8) int8   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Int8Stream) Reverse()  *Int8Stream {
	value := make([]int8, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Int8Stream) Unique()  *Int8Stream{
	value := make([]int8, 0, len(s.value))
	seen:=make(map[int8]struct{})
	for _, each := range s.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	s.value = value
	return s
}
func(s *Int8Stream) Append(given int8) *Int8Stream {
	s.value=append(s.value,given)
	return s
}
func(s *Int8Stream) Len() int {
	return len(s.value)
}
func(s *Int8Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Int8Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Int8Stream)  Sort()  *Int8Stream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] < s.value[j]
	})
	return s 
}
func(s *Int8Stream) All(fn func(int, int8)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}
func(s *Int8Stream) Any(fn func(int, int8)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}
func(s *Int8Stream) Paginate(size int)  [][]int8 {
	var pages  [][]int8
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
func(s *Int8Stream) Pop() int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Int8Stream) Prepend(given int8) *Int8Stream {
	s.value = append([]int8{given},s.value...)
	return s
}
func(s *Int8Stream) Max() int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max int8 = s.value[0]
	for _,each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func(s *Int8Stream) Min() int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min int8 = s.value[0]
	for _,each := range s.value {
		if each  < min {
			min = each
		}
	}
	return min
}
func(s *Int8Stream) Random() int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Int8Stream) Shuffle() *Int8Stream {
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
func(s *Int8Stream) Collect() []int8{
	return s.value
}
type Int8PStream struct{
	value	[]*int8
	defaultReturn *int8
}
func PStreamOfInt8(value []*int8) *Int8PStream {
	return &Int8PStream{value:value,defaultReturn:nil}
}
func(s *Int8PStream) OrElse(defaultReturn *int8)  *Int8PStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Int8PStream) Concate(given []*int8)  *Int8PStream {
	value := make([]*int8, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Int8PStream) Drop(n int)  *Int8PStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Int8PStream) Filter(fn func(int, *int8)bool)  *Int8PStream {
	value := make([]*int8, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Int8PStream) First() *int8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Int8PStream) Last() *int8 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Int8PStream) Map(fn func(int, *int8)) *Int8PStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Int8PStream) Reduce(fn func(*int8, *int8, int) *int8,initial *int8) *int8   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Int8PStream) Reverse()  *Int8PStream {
	value := make([]*int8, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Int8PStream) Unique()  *Int8PStream{
	value := make([]*int8, 0, len(s.value))
	seen:=make(map[*int8]struct{})
	for _, each := range s.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	s.value = value
	return s
}
func(s *Int8PStream) Append(given *int8) *Int8PStream {
	s.value=append(s.value,given)
	return s
}
func(s *Int8PStream) Len() int {
	return len(s.value)
}
func(s *Int8PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Int8PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Int8PStream)  Sort(less func(*int8,*int8) bool )  *Int8PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}
func(s *Int8PStream) All(fn func(int, *int8)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}



func(s *Int8PStream) Any(fn func(int, *int8)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}


func(s *Int8PStream) Paginate(size int)  [][]*int8 {
	var pages  [][]*int8
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
func(s *Int8PStream) Pop() *int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Int8PStream) Prepend(given *int8) *Int8PStream {
	s.value = append([]*int8{given},s.value...)
	return s
}
func(s *Int8PStream) Max() *int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *int8 = s.value[0]
	for _,each := range s.value {
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
func(s *Int8PStream) Min() *int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *int8 = s.value[0]
	for _,each := range s.value {
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
func(s *Int8PStream) Random() *int8{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Int8PStream) Shuffle() *Int8PStream {
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
func(s *Int8PStream) Collect() []*int8{
	return s.value
}
