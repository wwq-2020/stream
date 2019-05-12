
package commons
import (
	"sort"
	"math/rand"
)
type Uint32Stream struct{
	value	[]uint32
	defaultReturn uint32
}
func StreamOfUint32(value []uint32) *Uint32Stream {
	return &Uint32Stream{value:value,defaultReturn:0}
}
func(s *Uint32Stream) OrElase(defaultReturn uint32)  *Uint32Stream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Uint32Stream) Concate(given []uint32)  *Uint32Stream {
	value := make([]uint32, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Uint32Stream) Drop(n int)  *Uint32Stream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Uint32Stream) Filter(fn func(int, uint32)bool)  *Uint32Stream {
	value := make([]uint32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Uint32Stream) First() uint32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Uint32Stream) Last() uint32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Uint32Stream) Map(fn func(int, uint32)) *Uint32Stream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Uint32Stream) Reduce(fn func(uint32, uint32, int) uint32,initial uint32) uint32   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Uint32Stream) Reverse()  *Uint32Stream {
	value := make([]uint32, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Uint32Stream) Unique()  *Uint32Stream{
	value := make([]uint32, 0, len(s.value))
	seen:=make(map[uint32]struct{})
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
func(s *Uint32Stream) Append(given uint32) *Uint32Stream {
	s.value=append(s.value,given)
	return s
}
func(s *Uint32Stream) Len() int {
	return len(s.value)
}
func(s *Uint32Stream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Uint32Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Uint32Stream)  Sort()  *Uint32Stream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] < s.value[j]
	})
	return s 
}
func(s *Uint32Stream) All(fn func(int, uint32)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}
func(s *Uint32Stream) Any(fn func(int, uint32)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}
func(s *Uint32Stream) Paginate(size int)  [][]uint32 {
	var pages  [][]uint32
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
func(s *Uint32Stream) Pop() uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Uint32Stream) Prepend(given uint32) *Uint32Stream {
	s.value = append([]uint32{given},s.value...)
	return s
}
func(s *Uint32Stream) Max() uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max uint32 = s.value[0]
	for _,each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func(s *Uint32Stream) Min() uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min uint32 = s.value[0]
	for _,each := range s.value {
		if each  < min {
			min = each
		}
	}
	return min
}
func(s *Uint32Stream) Random() uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Uint32Stream) Shuffle() *Uint32Stream {
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
func(s *Uint32Stream) Collect() []uint32{
	return s.value
}
type Uint32PStream struct{
	value	[]*uint32
	defaultReturn *uint32
}
func PStreamOfUint32(value []*uint32) *Uint32PStream {
	return &Uint32PStream{value:value,defaultReturn:nil}
}
func(s *Uint32PStream) OrElse(defaultReturn *uint32)  *Uint32PStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *Uint32PStream) Concate(given []*uint32)  *Uint32PStream {
	value := make([]*uint32, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *Uint32PStream) Drop(n int)  *Uint32PStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *Uint32PStream) Filter(fn func(int, *uint32)bool)  *Uint32PStream {
	value := make([]*uint32, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *Uint32PStream) First() *uint32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *Uint32PStream) Last() *uint32 {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *Uint32PStream) Map(fn func(int, *uint32)) *Uint32PStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *Uint32PStream) Reduce(fn func(*uint32, *uint32, int) *uint32,initial *uint32) *uint32   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *Uint32PStream) Reverse()  *Uint32PStream {
	value := make([]*uint32, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *Uint32PStream) Unique()  *Uint32PStream{
	value := make([]*uint32, 0, len(s.value))
	seen:=make(map[*uint32]struct{})
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
func(s *Uint32PStream) Append(given *uint32) *Uint32PStream {
	s.value=append(s.value,given)
	return s
}
func(s *Uint32PStream) Len() int {
	return len(s.value)
}
func(s *Uint32PStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *Uint32PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *Uint32PStream)  Sort(less func(*uint32,*uint32) bool )  *Uint32PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}
func(s *Uint32PStream) All(fn func(int, *uint32)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}



func(s *Uint32PStream) Any(fn func(int, *uint32)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}


func(s *Uint32PStream) Paginate(size int)  [][]*uint32 {
	var pages  [][]*uint32
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
func(s *Uint32PStream) Pop() *uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *Uint32PStream) Prepend(given *uint32) *Uint32PStream {
	s.value = append([]*uint32{given},s.value...)
	return s
}
func(s *Uint32PStream) Max() *uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *uint32 = s.value[0]
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
func(s *Uint32PStream) Min() *uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *uint32 = s.value[0]
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
func(s *Uint32PStream) Random() *uint32{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *Uint32PStream) Shuffle() *Uint32PStream {
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
func(s *Uint32PStream) Collect() []*uint32{
	return s.value
}
