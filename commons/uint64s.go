
package commons
import (
	"sort"
	"math/rand"
)
type Uint64Stream struct{
	value	[]uint64
	defaultReturn uint64
}
func StreamOfUint64(value []uint64) *Uint64Stream {
	return &Uint64Stream{value:value,defaultReturn:0}
}
func(s *Uint64Stream) OrElase(defaultReturn uint64)  *Uint64Stream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Uint64Stream) Concate(given []uint64)  *Uint64Stream {
	value := make([]uint64, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Uint64Stream) Drop(n int)  *Uint64Stream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Uint64Stream) Filter(fn func(int, uint64)bool)  *Uint64Stream {
	value := make([]uint64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Uint64Stream) First() uint64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Uint64Stream) Last() uint64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Uint64Stream) Map(fn func(int, uint64)) *Uint64Stream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Uint64Stream) Reduce(fn func(uint64, uint64, int) uint64,initial uint64) uint64   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Uint64Stream) Reverse()  *Uint64Stream {
	value := make([]uint64, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Uint64Stream) Unique()  *Uint64Stream{
	value := make([]uint64, 0, len(s.value))
	seen:=make(map[uint64]struct{})
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
func(s *Uint64Stream) Append(given uint64) *Uint64Stream {
	s.value=append(s.value,given)
	return s
}
func(s *Uint64Stream) Len() int {
	return len(s.value)
}
func(s *Uint64Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Uint64Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Uint64Stream)  Sort()  *Uint64Stream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] < s.value[j]
	})
	return s 
}
func(s *Uint64Stream) All(fn func(int, uint64)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}
func(s *Uint64Stream) Any(fn func(int, uint64)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}
func(s *Uint64Stream) Paginate(size int)  [][]uint64 {
	var pages  [][]uint64
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
func(s *Uint64Stream) Pop() uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Uint64Stream) Prepend(given uint64) *Uint64Stream {
	s.value = append([]uint64{given},s.value...)
	return s
}
func(s *Uint64Stream) Max() uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max uint64 = s.value[0]
	for _,each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func(s *Uint64Stream) Min() uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min uint64 = s.value[0]
	for _,each := range s.value {
		if each  < min {
			min = each
		}
	}
	return min
}
func(s *Uint64Stream) Random() uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Uint64Stream) Shuffle() *Uint64Stream {
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
func(s *Uint64Stream) Collect() []uint64{
	return s.value
}
type Uint64PStream struct{
	value	[]*uint64
	defaultReturn *uint64
}
func PStreamOfUint64(value []*uint64) *Uint64PStream {
	return &Uint64PStream{value:value,defaultReturn:nil}
}
func(s *Uint64PStream) OrElse(defaultReturn *uint64)  *Uint64PStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Uint64PStream) Concate(given []*uint64)  *Uint64PStream {
	value := make([]*uint64, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Uint64PStream) Drop(n int)  *Uint64PStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Uint64PStream) Filter(fn func(int, *uint64)bool)  *Uint64PStream {
	value := make([]*uint64, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Uint64PStream) First() *uint64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Uint64PStream) Last() *uint64 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Uint64PStream) Map(fn func(int, *uint64)) *Uint64PStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Uint64PStream) Reduce(fn func(*uint64, *uint64, int) *uint64,initial *uint64) *uint64   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Uint64PStream) Reverse()  *Uint64PStream {
	value := make([]*uint64, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Uint64PStream) Unique()  *Uint64PStream{
	value := make([]*uint64, 0, len(s.value))
	seen:=make(map[*uint64]struct{})
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
func(s *Uint64PStream) Append(given *uint64) *Uint64PStream {
	s.value=append(s.value,given)
	return s
}
func(s *Uint64PStream) Len() int {
	return len(s.value)
}
func(s *Uint64PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Uint64PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Uint64PStream)  Sort(less func(*uint64,*uint64) bool )  *Uint64PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}
func(s *Uint64PStream) All(fn func(int, *uint64)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}



func(s *Uint64PStream) Any(fn func(int, *uint64)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}



func(s *Uint64PStream) Paginate(size int)  [][]*uint64 {
	var pages  [][]*uint64
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
func(s *Uint64PStream) Pop() *uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Uint64PStream) Prepend(given *uint64) *Uint64PStream {
	s.value = append([]*uint64{given},s.value...)
	return s
}
func(s *Uint64PStream) Max() *uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *uint64 = s.value[0]
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
func(s *Uint64PStream) Min() *uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *uint64 = s.value[0]
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
func(s *Uint64PStream) Random() *uint64{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Uint64PStream) Shuffle() *Uint64PStream {
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
func(s *Uint64PStream) Collect() []*uint64{
	return s.value
}
