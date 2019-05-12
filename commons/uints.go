
package commons
import (
	"sort"
	"math/rand"
)
type UintStream struct{
	value	[]uint
	defaultReturn uint
}
func StreamOfUint(value []uint) *UintStream {
	return &UintStream{value:value,defaultReturn:0}
}
func(s *UintStream) OrElase(defaultReturn uint)  *UintStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *UintStream) Concate(given []uint)  *UintStream {
	value := make([]uint, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *UintStream) Drop(n int)  *UintStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *UintStream) Filter(fn func(int, uint)bool)  *UintStream {
	value := make([]uint, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *UintStream) First() uint {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *UintStream) Last() uint {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *UintStream) Map(fn func(int, uint)) *UintStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *UintStream) Reduce(fn func(uint, uint, int) uint,initial uint) uint   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *UintStream) Reverse()  *UintStream {
	value := make([]uint, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *UintStream) Unique()  *UintStream{
	value := make([]uint, 0, len(s.value))
	seen:=make(map[uint]struct{})
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
func(s *UintStream) Append(given uint) *UintStream {
	s.value=append(s.value,given)
	return s
}
func(s *UintStream) Len() int {
	return len(s.value)
}
func(s *UintStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *UintStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *UintStream)  Sort()  *UintStream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] < s.value[j]
	})
	return s 
}
func(s *UintStream) All(fn func(int, uint)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}
func(s *UintStream) Any(fn func(int, uint)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}
func(s *UintStream) Paginate(size int)  [][]uint {
	var pages  [][]uint
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
func(s *UintStream) Pop() uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *UintStream) Prepend(given uint) *UintStream {
	s.value = append([]uint{given},s.value...)
	return s
}
func(s *UintStream) Max() uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max uint = s.value[0]
	for _,each := range s.value {
		if max < each {
			max = each
		}
	}
	return max
}
func(s *UintStream) Min() uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min uint = s.value[0]
	for _,each := range s.value {
		if each  < min {
			min = each
		}
	}
	return min
}
func(s *UintStream) Random() uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *UintStream) Shuffle() *UintStream {
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
func(s *UintStream) Collect() []uint{
	return s.value
}
type UintPStream struct{
	value	[]*uint
	defaultReturn *uint
}
func PStreamOfUint(value []*uint) *UintPStream {
	return &UintPStream{value:value,defaultReturn:nil}
}
func(s *UintPStream) OrElse(defaultReturn *uint)  *UintPStream {
	s.defaultReturn = defaultReturn
	return s
}
func(s *UintPStream) Concate(given []*uint)  *UintPStream {
	value := make([]*uint, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}
func(s *UintPStream) Drop(n int)  *UintPStream {
	l := len(s.value) - n
	if l < 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}
func(s *UintPStream) Filter(fn func(int, *uint)bool)  *UintPStream {
	value := make([]*uint, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
func(s *UintPStream) First() *uint {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}
func(s *UintPStream) Last() *uint {
	if len(s.value) <= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}
func(s *UintPStream) Map(fn func(int, *uint)) *UintPStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}
func(s *UintPStream) Reduce(fn func(*uint, *uint, int) *uint,initial *uint) *uint   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}
func(s *UintPStream) Reverse()  *UintPStream {
	value := make([]*uint, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}
func(s *UintPStream) Unique()  *UintPStream{
	value := make([]*uint, 0, len(s.value))
	seen:=make(map[*uint]struct{})
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
func(s *UintPStream) Append(given *uint) *UintPStream {
	s.value=append(s.value,given)
	return s
}
func(s *UintPStream) Len() int {
	return len(s.value)
}
func(s *UintPStream) IsEmpty() bool {
	return len(s.value) == 0
}
func(s *UintPStream) IsNotEmpty() bool {
	return len(s.value) != 0
}
func(s *UintPStream)  Sort(less func(*uint,*uint) bool )  *UintPStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}
func(s *UintPStream) All(fn func(int, *uint)bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}



func(s *UintPStream) Any(fn func(int, *uint)bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}



func(s *UintPStream) Paginate(size int)  [][]*uint {
	var pages  [][]*uint
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
func(s *UintPStream) Pop() *uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}
func(s *UintPStream) Prepend(given *uint) *UintPStream {
	s.value = append([]*uint{given},s.value...)
	return s
}
func(s *UintPStream) Max() *uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var max *uint = s.value[0]
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
func(s *UintPStream) Min() *uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	var min *uint = s.value[0]
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
func(s *UintPStream) Random() *uint{
	if len(s.value) <= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}
func(s *UintPStream) Shuffle() *UintPStream {
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
func(s *UintPStream) Collect() []*uint{
	return s.value
}
